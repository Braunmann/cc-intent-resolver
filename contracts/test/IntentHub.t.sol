// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.13;

import {Test} from "forge-std/Test.sol";
import {ERC20Mock} from "@openzeppelin/contracts/mocks/token/ERC20Mock.sol";

import {IntentHub, Intent, IntentStatus, ZeroAmount, NotIntentMaker} from "../src/IntentHub.sol";

contract IntentHubTest is Test {
    IntentHub public intentHub;

    address maker = makeAddr("maker");
    address recipient = makeAddr("recipient");
    ERC20Mock token;

    function setUp() public {
        intentHub = new IntentHub();
        token = new ERC20Mock();
        token.mint(maker, 1000e18);
    }

    function test_createIntent_created() public {
        vm.startPrank(maker);

        token.approve(address(intentHub), 100e18);

        bytes32 intentId = intentHub.createIntent(
            address(token),
            100e18,
            address(token),
            50e18,
            1,
            recipient,
            uint64(block.timestamp + 1 hours)
        );

        vm.stopPrank();

        Intent memory intent = intentHub.getIntent(intentId);
        assertEq(intent.maker, maker);
        // Assert can not compare enums -> cast on uint8
        assertEq(uint8(intent.status), uint8(IntentStatus.Created));
    }

    function test_createIntent_transfersTokensToEscrow() public {
        vm.startPrank(maker);

        token.approve(address(intentHub), 100e18);

        intentHub.createIntent(
            address(token),
            100e18,
            address(token),
            50e18,
            1,
            recipient,
            uint64(block.timestamp + 1 hours)
        );

        vm.stopPrank();

        assertEq(token.balanceOf(address(intentHub)), 100e18);
    }

    function test_cancelIntent_returnsTokensToMaker() public {
        vm.startPrank(maker);

        token.approve(address(intentHub), 100e18);

        bytes32 intentId = intentHub.createIntent(
            address(token),
            100e18,
            address(token),
            50e18,
            1,
            recipient,
            uint64(block.timestamp + 1 hours)
        );

        intentHub.cancelIntent(intentId);
        vm.stopPrank();

        Intent memory intent = intentHub.getIntent(intentId);
        assertEq(uint8(intent.status), uint8(IntentStatus.Cancelled));
        assertEq(token.balanceOf(maker), 1000e18);
    }

    function test_cancelIntent_revertsWhenNotMaker() public {
        vm.startPrank(maker);

        token.approve(address(intentHub), 100e18);

        bytes32 intentId = intentHub.createIntent(
            address(token),
            100e18,
            address(token),
            50e18,
            1,
            recipient,
            uint64(block.timestamp + 1 hours)
        );

        vm.stopPrank();

        vm.startPrank(recipient);
        vm.expectRevert(abi.encodeWithSelector(NotIntentMaker.selector));       
        intentHub.cancelIntent(intentId);
        vm.stopPrank();
    }

    function test_createIntent_revertsWhenZeroAmount() public {
        vm.startPrank(maker);

        token.approve(address(intentHub), 0);

        vm.expectRevert(abi.encodeWithSelector(ZeroAmount.selector));

        intentHub.createIntent(
            address(token),
            0,
            address(token),
            50e18,
            1,
            recipient,
            uint64(block.timestamp + 1 hours)
        );
    }

    function test_fulfillIntent_setsStatusToFulfilled() public {
        vm.startPrank(maker);

        token.approve(address(intentHub), 100e18);

        address solver = address(0x123);

        bytes32 intentId = intentHub.createIntent(
            address(token),
            100e18,
            address(token),
            50e18,
            1,
            recipient,
            uint64(block.timestamp + 1 hours)
        );

        intentHub.fulfillIntent(intentId, solver);
        vm.stopPrank();

        Intent memory intent = intentHub.getIntent(intentId);
        assertEq(uint8(intent.status), uint8(IntentStatus.Fulfilled));
    }

    function test_settleIntent_transfersEscrowToSolver() public {
        vm.startPrank(maker);
        token.approve(address(intentHub), 100e18);
        
        bytes32 intentId = intentHub.createIntent(
            address(token),
            100e18,
            address(token),
            50e18,
            1,
            recipient,
            uint64(block.timestamp + 1 hours)
        );
        vm.stopPrank();

        address solver = address(0x123);
        vm.startPrank(address(solver));    

        intentHub.fulfillIntent(intentId, solver);
        intentHub.settleIntent(intentId);

        vm.stopPrank();
        
        assertEq(token.balanceOf(solver), 100e18);
    }

}
