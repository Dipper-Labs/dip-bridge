pragma solidity ^0.5.17;

import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v2.5.0/contracts/math/SafeMath.sol";

contract DipManager {
    using SafeMath for uint256;

    event TokenMinted(bytes32 ethTxid, address to, uint256 amount);

    mapping(bytes32 => bool) public txFilter;
    address public admin;

    constructor(address addr) public {
        admin = addr;
    }

    function MintToken(bytes32 txid, address payable to, uint256 amount) public {
        require(txFilter[txid] == false, "txid already processed");
        require(msg.sender == admin, "no authorized account");

        to.transfer(amount);
        txFilter[txid] = true;

        emit TokenMinted(txid, to, amount);
    }

    function Grant(address addr) public {
        require(msg.sender == admin, "no authorized account");
        admin = addr;
    }
}

