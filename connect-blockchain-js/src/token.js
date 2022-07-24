import { ethers } from 'ethers'

/**
 * @param {ethers.Contract} contract
 * @returns {Promise<string>}
 */
export async function getTokenSymbol(contract) {
  const symbol = await contract.symbol()
  return symbol
}

/**
 * @param {ethers.Contract[]} contracts
 * @returns {Promise<string[]>}
 */
export async function getTokenSymbolsV1(contracts) {
  const symbols = []
  for (const contract of contracts) {
    const symbol = await getTokenSymbol(contract)
    symbols.push(symbol)
  }
  return symbols
}

/**
 * @param {ethers.Contract[]} contracts
 * @returns {string[]}
 */
export async function getTokenSymbolsV2(contracts) {
  const symbols = await Promise.all(
    contracts.map(async (c) => {
      return await getTokenSymbol(c)
    }),
  )
  return symbols
}
