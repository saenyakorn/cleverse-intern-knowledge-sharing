import { ethers } from 'ethers'
import erc20Abi from '../../abi/ERC20.json'

// On BSC Chain
const provider = new ethers.providers.JsonRpcProvider('https://rpc.ankr.com/bsc')

// Constants
const WBNDAddr = '0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c'
const BUSDAddr = '0xe9e7cea3dedca5984780bafc599bd69add087d56'

// Call ERC20 contract
const WBNBContract = new ethers.Contract(WBNDAddr, erc20Abi, provider)
const BUSDContract = new ethers.Contract(BUSDAddr, erc20Abi, provider)

/**
 * @param {ethers.Contract} contract
 * @returns {string}
 */
async function getTokenSymbol(contract) {
  const symbol = await contract.symbol()
  return symbol
}

/**
 * @returns {number}
 */
async function queryCurrentBlockNumber() {
  const blockNumber = await provider.getBlockNumber()
  return blockNumber
}

console.log(await queryCurrentBlockNumber())
console.log(await getTokenSymbol(WBNBContract))
console.log(await getTokenSymbol(BUSDContract))
