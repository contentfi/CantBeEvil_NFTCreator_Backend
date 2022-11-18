
// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "./CantBeEvil.sol";

contract Token721 is ERC721URIStorage, ReentrancyGuard, CantBeEvil {
    event TokenMinted(
        uint256 _tokenId,
        string _tokenURI
    );

    event MetadataUriChanged( string _metadataUri);

    using Strings for uint256;

    uint256 private tokenCount = 0;
    address private creatorAddress;
    string private metadataURI;

    constructor(
        string memory _name,
        string memory _symbol,
        string memory _metadataURI,
        LicenseVersion _license
    ) ERC721(_name, _symbol) CantBeEvil(_license) {
        creatorAddress = tx.origin;
        metadataURI = _metadataURI;
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        virtual
        override(CantBeEvil, ERC721)
        returns (bool)
    {
        return CantBeEvil.supportsInterface(interfaceId) || ERC721.supportsInterface(interfaceId);
    }

    function getCreatorAddress() public view returns (address) {
        return creatorAddress;
    }

    function mint(address to_, string memory tokenURI_)
        external
        isCreator
        nonReentrant
    {
        tokenCount++;
        uint256 _tokenId = tokenCount;
        _safeMint(to_, _tokenId);
        _setTokenURI(_tokenId, tokenURI_);
        emit TokenMinted(
            _tokenId,
            tokenURI_
        );
    }

    function setCollectionMetadataURI(string memory _metadataURI)
        external
        isCreator
    {
        metadataURI = _metadataURI;
        emit MetadataUriChanged(_metadataURI);
    }

    modifier isCreator() {
        require(creatorAddress == msg.sender, "caller is not creator");
        _;
    }
}
