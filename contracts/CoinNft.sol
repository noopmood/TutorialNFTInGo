// Contract based on https://docs.openzeppelin.com/contracts/4.x/erc1155
// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;
import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";

contract CoinNft is ERC1155 {
    // GOLD_ID = 0000000000000000000000000000000000000000000000000000000000000000;
    // SILVER_ID = 0000000000000000000000000000000000000000000000000000000000000001;
    
    constructor() ERC1155("https://raw.githubusercontent.com/noopmood/TutorialNFTInGo/main/metadata/{id}.json") {}
    
    function mintCaller(uint256 tokenId, uint256 amount) public {
        _mint(msg.sender, tokenId, amount, "");
    }
    function mintAddress(uint256 tokenId, uint256 amount, address addr) public {
        _mint(addr, tokenId, amount, "");
    }
    function mintCallerBatch(uint256[] memory tokenIds, uint256[] memory amounts) public {
        _mintBatch(msg.sender, tokenIds, amounts, "");
    }
    function mintAddressBatch(uint256[] memory tokenIds, uint256[] memory amounts, address addr) public {
        _mintBatch(addr, tokenIds, amounts, "");
    }
    function burnCaller(uint256 tokenId, uint256 amount) public {
        _burn(msg.sender, tokenId, amount);
    }
    function burnAddress(uint256 tokenId, uint256 amount, address addr) public {
        _burn(addr, tokenId, amount);
    }
}

