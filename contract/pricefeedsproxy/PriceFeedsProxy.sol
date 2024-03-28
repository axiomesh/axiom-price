// SPDX-License-Identifier: MIT
pragma solidity 0.8.0;

contract PriceFeedsProxy {
    address public implementation;
    address private owner;

    modifier onlyOwner() {
        require(msg.sender == owner, "Only the owner can call this function.");
        _;
    }

    constructor(address _implementation) {
        implementation = _implementation;
        owner = msg.sender;
    }

    function upgradeImplementation(address _newImplementation) external onlyOwner {
        implementation = _newImplementation;
    }

    function transferOwnership(address _newOwner) external onlyOwner {
        require(_newOwner != address(0), "New owner cannot be the zero address.");
        owner = _newOwner;
    }

    fallback() external payable {
        address _impl = implementation;
        require(_impl != address(0), "Implementation contract not set.");

        assembly {
            let ptr := mload(0x40)
            calldatacopy(ptr, 0, calldatasize())
            let result := delegatecall(gas(), _impl, ptr, calldatasize(), 0, 0)
            let size := returndatasize()
            returndatacopy(ptr, 0, size)

            switch result
            case 0 { revert(ptr, size) }
            default { return(ptr, size) }
        }
    }
}
