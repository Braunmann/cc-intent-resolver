// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import { SafeERC20, IERC20 } from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

enum IntentStatus {
    Created,    
    Fulfilled,  
    Settled,    
    Cancelled   
}

error ZeroAmount();
error DeadlineInPast();
error NotIntentMaker();
error IntentNotFound();
error IntentNotCancellable();

event IntentCreated(
    bytes32 indexed intentId, // filter by indexed field
    address indexed maker, // filter by indexed field
    address inputToken, 
    uint256 inputAmount, 
    address outputToken, 
    uint256 minOutputAmount, 
    uint32 targetChainId, 
    address recipient, 
    uint64 deadline, 
    uint64 nonce
);

event IntentCancelled(bytes32 indexed intentId);

struct Intent {
    bytes32 id;
    address maker;
    address inputToken;
    uint256 inputAmount;
    address outputToken;
    uint256 minOutputAmount;
    uint32 targetChainId;
    address recipient;
    uint64 deadline;
    uint64 nonce;
    IntentStatus status;
}

contract IntentHub {
    using SafeERC20 for IERC20;

    mapping(bytes32 => Intent) public intents;
    mapping(address => uint64) public nonces;
    
    function createIntent( 
        address inputToken, 
        uint256 inputAmount, 
        address outputToken, 
        uint256 minOutputAmount, 
        uint32 targetChainId, 
        address recipient, 
        uint64 deadline
    ) external returns (bytes32) {
        if(inputAmount == 0) revert ZeroAmount();
        if(deadline <= block.timestamp) revert DeadlineInPast();

        uint64 nonce = nonces[msg.sender]++;
        
        bytes32 intentId = keccak256(abi.encode(
            msg.sender,
            inputToken,
            inputAmount,
            outputToken,
            minOutputAmount,
            targetChainId,
            recipient,
            deadline,
            nonce
        ));

        intents[intentId] = Intent({
            id: intentId,
            maker: msg.sender,
            inputToken: inputToken,
            inputAmount: inputAmount,
            outputToken: outputToken,
            minOutputAmount: minOutputAmount,
            targetChainId: targetChainId,
            recipient: recipient,
            deadline: deadline,
            nonce: nonce,
            status: IntentStatus.Created
        });

        // It cost gas, it can be directly emited
        Intent memory intent = intents[intentId];

        emit IntentCreated(
            intent.id, 
            intent.maker, 
            intent.inputToken, 
            intent.inputAmount, 
            intent.outputToken, 
            intent.minOutputAmount, 
            intent.targetChainId, 
            intent.recipient, 
            intent.deadline, 
            intent.nonce
        );

        IERC20(inputToken).safeTransferFrom(msg.sender, address(this), inputAmount);

        return intentId;
    }

    function cancelIntent(bytes32 intentId) external {
        Intent storage intent = intents[intentId];
        if(intent.maker == address(0)) revert IntentNotFound();
        if(intent.maker != msg.sender) revert NotIntentMaker();
        if(intent.status != IntentStatus.Created) revert IntentNotCancellable();
        intent.status = IntentStatus.Cancelled;

        emit IntentCancelled(intentId);

        IERC20(intent.inputToken).safeTransfer(intent.maker, intent.inputAmount);
    }

    function getIntent(bytes32 intentId) external view returns (Intent memory) {
        return intents[intentId];
    }
}

 