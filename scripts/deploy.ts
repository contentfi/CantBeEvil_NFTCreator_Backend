import { ethers } from "hardhat";

async function main() {
  const p = await ethers.getContractFactory("Token721");
  const pr = await p.deploy();
  await pr.deployed();
  console.log("proxy deployed to:", pr.address);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
