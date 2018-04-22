pragma solidity ^0.4.20;

library Signature {
    function verify(bytes32 hash, bytes32 r, bytes32 s, uint8 v, address vettingAddress) internal pure returns(bool)
    {
        return ecrecover(hash, v, r, s) == vettingAddress;
    }
}