pragma solidity ^0.5.17;

import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v2.5.0/contracts/math/SafeMath.sol";
import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v2.5.0/contracts/token/ERC20/IERC20.sol";
import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v2.5.0/contracts/token/ERC20/SafeERC20.sol";

contract DipManager {
    using SafeMath for uint256;
    using SafeERC20 for IERC20;

    event TokenLocked(address indexed from, string to, uint256 amount);

    mapping(address => mapping(string => uint256)) public lockInfo;

    address public erc20Addr;

    constructor(address addr) public {
        erc20Addr = addr;
    }

    function LockToken(string memory dipAddr, uint256 amount) public {
        require(bytes(dipAddr).length == 42, "dipAddr must 42 bytes length");
        string memory s = substring(dipAddr, 0, 3);
        require(keccak256(bytes("dip")) == keccak256(bytes(s)), "dipAddr must start with dip");

        IERC20 dipERC20 = IERC20(erc20Addr);
        dipERC20.safeTransferFrom(msg.sender, address(this), amount);
        lockInfo[msg.sender][dipAddr] = lockInfo[msg.sender][dipAddr].add(amount);
        emit TokenLocked(msg.sender, dipAddr, amount);
    }

    function substring(string memory str, uint startIndex, uint endIndex) public pure returns (string memory) {
        bytes memory strBytes = bytes(str);
        bytes memory result = new bytes(endIndex-startIndex);
        for(uint i = startIndex; i < endIndex; i++) {
            result[i-startIndex] = strBytes[i];
        }
        return string(result);
    }
}

