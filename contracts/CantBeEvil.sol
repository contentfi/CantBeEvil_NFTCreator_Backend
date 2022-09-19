// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";

interface ICantBeEvil {
  function getLicenseURI() external view returns (string memory);

  function getLicenseName() external view returns (string memory);
}

library License {
  enum LicenseVersion {
    CBE_CC0,
    CBE_ECR,
    CBE_NECR,
    CBE_NECR_HS,
    CBE_PR,
    CBE_PR_HS
  }
}

contract CantBeEvil is ERC165, ICantBeEvil {
  using Strings for uint256;
  string internal constant _BASE_LICENSE_URI = "ar://_D9kN1WrNWbCq55BSAGRbTB4bS3v8QAPTYmBThSbX3A/";
  License.LicenseVersion public licenseVersion; // return string

  constructor(License.LicenseVersion _licenseVersion) {
    licenseVersion = _licenseVersion;
  }

  function getLicenseURI() public view returns (string memory) {
    return string.concat(_BASE_LICENSE_URI, uint256(licenseVersion).toString());
  }

  function getLicenseName() public view returns (string memory) {
    return _getLicenseVersionKeyByValue(licenseVersion);
  }

  function supportsInterface(bytes4 interfaceId)
    public
    view
    virtual
    override(ERC165)
    returns (bool)
  {
    return interfaceId == type(ICantBeEvil).interfaceId || super.supportsInterface(interfaceId);
  }

  function _getLicenseVersionKeyByValue(License.LicenseVersion _licenseVersion)
    internal
    pure
    returns (string memory)
  {
    require(uint8(_licenseVersion) <= 6);
    if (License.LicenseVersion.CBE_CC0 == _licenseVersion) return "CBE_CC0";
    if (License.LicenseVersion.CBE_ECR == _licenseVersion) return "CBE_ECR";
    if (License.LicenseVersion.CBE_NECR == _licenseVersion) return "CBE_NECR";
    if (License.LicenseVersion.CBE_NECR_HS == _licenseVersion) return "CBE_NECR_HS";
    if (License.LicenseVersion.CBE_PR == _licenseVersion) return "CBE_PR";
    else return "CBE_PR_HS";
  }
}