// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script} from "forge-std/Script.sol";
import {IntentHub} from "../src/IntentHub.sol";

contract Deploy is Script {
    function run() external {
        vm.startBroadcast();
        new IntentHub();
        vm.stopBroadcast();
    }
}