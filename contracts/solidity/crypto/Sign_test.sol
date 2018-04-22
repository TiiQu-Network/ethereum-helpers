pragma solidity ^0.4.0;

import "./Signature.sol";

contract Sign {

    address vettingAddress;

    function Sign(address VettingAddress){
        vettingAddress = VettingAddress;
    }

    function verify(bytes32 r, bytes32 s, uint8 v) view public returns (bool) {
        return Signature.verify(keccak256(msg.sender), r, s, v, vettingAddress);
    }

    function getVettingAddress() view public returns (address)
    {
        return vettingAddress;
    }
}
