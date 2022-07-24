import { ethers } from 'ethers'
import erc20Abi from '../../abi/ERC20.json'
import { getTokenSymbolsV1, getTokenSymbolsV2 } from './token.js'

// On BSC Chain
const provider = new ethers.providers.JsonRpcProvider('https://rpc.ankr.com/bsc')

// Constants
const WBNDAddr = '0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c'
const BUSDAddr = '0xe9e7cea3dedca5984780bafc599bd69add087d56'

// Call ERC20 contract
const WBNBContract = new ethers.Contract(WBNDAddr, erc20Abi, provider)
const BUSDContract = new ethers.Contract(BUSDAddr, erc20Abi, provider)

const mockContracts = Array.from({ length: 10 }, (_, i) => [WBNBContract, BUSDContract]).reduce(
  (acc, [a, b]) => [...acc, a, b],
  [],
)
var startTime = new Date()
console.log(await getTokenSymbolsV1(mockContracts))
var endTime = new Date()
console.log(`getTokenSymbolsV1: ${(endTime.getTime() - startTime.getTime()) / 1000}s`)

startTime = new Date()
console.log(await getTokenSymbolsV2(mockContracts))
endTime = new Date()
console.log(`getTokenSymbolsV2: ${(endTime.getTime() - startTime.getTime()) / 1000}s`)
