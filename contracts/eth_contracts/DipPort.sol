pragma solidity ^0.5.17;

import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v2.5.0/contracts/math/SafeMath.sol";
import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v2.5.0/contracts/token/ERC20/IERC20.sol";
import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v2.5.0/contracts/token/ERC20/SafeERC20.sol";

contract DipManager {
    using SafeMath for uint256;
    using SafeERC20 for IERC20;
    
    event TokenLocked(address indexed from, string to, uint256 amount);
    
    mapping(address => mapping(string => uint256)) private lockInfo;
    
    address public erc20Addr;
    
    constructor(address addr) public {
        erc20Addr = addr;
    }
    
    function lockToken(string memory dipAddr, uint256 amount) public {
        IERC20 ethToken = IERC20(erc20Addr);
        ethToken.safeTransferFrom(msg.sender, address(this), amount);
        lockInfo[msg.sender][dipAddr] = lockInfo[msg.sender][dipAddr] + amount;
        emit TokenLocked(msg.sender, dipAddr, amount);
    }
    
    function unlockToken(string memory dipAddr, uint256 amount) public {
        
    }
}


