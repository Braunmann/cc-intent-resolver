// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

enum IntentStatus {
    Created,    
    Fulfilled,  
    Settled,    
    Cancelled   
}

error ZeroAmount();
error DeadlineInPast();

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

        // TODO: Implement intent creation logic
        return intentId;
    }
}
