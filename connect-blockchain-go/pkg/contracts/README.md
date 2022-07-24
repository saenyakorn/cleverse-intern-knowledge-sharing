# contracts

### Flatten Methods

- `npx hardhat flatten {INPUT_PATH} > {OUTPUT_PATH}`
- `sol-merger {INPUT_PATH} {OUTPUT_PATH}`

reference: https://github.com/RyuuGan/sol-merger

### Generate ABI and Bytecode with bytecode optimizer

```
solc --bin --abi --optimize -o {OUTPUT_DIR} {INPUT_FILES...}
```

### Generate Go Contract file (with deploy methods)

```
$ abigen --abi {ABI_FILE} --bin {BINARY_FILE} --pkg contracts --out {GOLANG_OUTPUT_FILE} --type {CONTRACT_NAME}
```
