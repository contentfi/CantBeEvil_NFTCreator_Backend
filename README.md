# SmartContract for https://fanmake.xyz

## Overview
FanMake contract is a ERC721 smart contract which implement rarible royalty specification.
Creator can set the royalties through this contract, and space contract will recieve royalty fee
when a token was bought from rarible NFT marketplace.

Space contract is a wallet contract to hold the royalties earned from derivative NFT sales,
IP holder can withdraw royalties from this contract.

Proxy contract is an endpoint to create FanMake contract and Space contract.

## Compile & Test & Deploy:
```shell
GAS_REPORT=true npx hardhat compile
npx hardhat test
npx hardhat run --network polygon-mumbai scripts/deploy.ts
```
