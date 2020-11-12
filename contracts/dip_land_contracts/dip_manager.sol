pragma solidity ^0.5.17;

import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v2.5.0/contracts/math/SafeMath.sol";

contract DipManager {
    using SafeMath for uint256;

    event TokenMinted(bytes32 ethTxid, address to, uint256 amount);

    mapping(bytes32 => bool) public txFilter;
    address public admin;
    bool public maintaining = false;
    uint256 public mintedAmount = 0;

    constructor(address addr) public {
        admin = addr;
    }

    function MintToken(bytes32 txid, address payable to, uint256 amount) public {
        require(maintaining == false, "maintaining");
        require(txFilter[txid] == false, "txid already processed");
        require(msg.sender == admin, "no authorized account");

        to.transfer(amount);
        txFilter[txid] = true;
        mintedAmount = mintedAmount.add(amount);

        emit TokenMinted(txid, to, amount);
    }

    function Grant(address addr) public {
        require(msg.sender == admin, "no authorized account");
        admin = addr;
    }

    function Maintain() public {
        require(msg.sender == admin, "no authorized account");
        maintaining = true;
    }

    function UnMaintain() public {
        require(msg.sender == admin, "no authorized account");
        maintaining = false;
    }
}

