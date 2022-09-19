import { FanMake, Proxy, Space } from "@contracts";
const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("Proxy contract", function () {
    let pr: Proxy;
    beforeEach(async () => {
        const [owner] = await ethers.getSigners();
        const p = await ethers.getContractFactory("FanMake");
        pr = await p.deploy() as Proxy;
    });
    it("create space and create collection", async function () {
        expect(await pr.createSpace(
            "test space",
            "0x55085B2Fd83323d98d30d6B3342cc39de6D527f8",
            "0x55085B2Fd83323d98d30d6B3342cc39de6D527f8",
            "",
        )).to.emit(pr, 'SpaceCreated').withArgs(
            "test space",
            "0x55085B2Fd83323d98d30d6B3342cc39de6D527f8",
            "0x55085B2Fd83323d98d30d6B3342cc39de6D527f8",
            "",
        );
    });
});