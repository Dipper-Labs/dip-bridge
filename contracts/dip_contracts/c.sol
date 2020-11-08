pragma solidity ^0.5.17;


contract DipManager {
    mapping(bytes32 => bool) public txFilter;

    function mintToken(bytes32 txid, address payable to, uint256 amount) public {
        require(txFilter[txid] == false, "txid already processed");
        
        to.transfer(amount);
        txFilter[txid] = true;
    }
}


