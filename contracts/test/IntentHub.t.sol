// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.13;

import { Test } from "forge-std/Test.sol";
import { ERC20Mock } from "@openzeppelin/contracts/mocks/token/ERC20Mock.sol";

import { IntentHub } from "../src/IntentHub.sol";

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
}
