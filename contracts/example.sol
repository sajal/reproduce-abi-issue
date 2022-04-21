// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.3;
pragma experimental ABIEncoderV2;

contract Example{
    uint256 constant NUM_SIZE = 5;

    struct Numbers{
        uint256 a;
        uint256 b;
    }

    constructor(){}

    function doMath(Numbers[NUM_SIZE] calldata items) external pure returns (uint256){
        uint256 res = 0;
        for (uint i=0; i<NUM_SIZE; i++) {
            res = res + (items[i].a * items[i].b);
        }
        return res;
    }
}
