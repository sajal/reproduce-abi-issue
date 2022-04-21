# Issue with vyper-generated ABI

We noticed an issue with ABI generated with vyper when using struct-arrays. I don't yet know if this is an issue with vyper, or with other projects attempting to consume the ABI. Specifically the issue is with `type": "(uint256,uint256)[5]"` part. Would like to know if this type is legal per the spec, if so then I should open bug report with upstream projects.


Here is the status

|platform   |solidity       |vyper              |
|-----------|---------------|-------------------|
|web3.js    |TODO           |TODO               |
|ethers     |TODO           |TODO               |
|abigen(go) |Valid bindings |Invalid bindings   |
|web3.py    |TODO           |TODO               |


## web3.js

TODO

## ethers

TODO

## abigen(go)

Solidity ABI generates correct method

```go
// ExampleNumbers is an auto generated low-level Go binding around an user-defined struct.
type ExampleNumbers struct {
	A *big.Int
	B *big.Int
}

// DoMath is a free data retrieval call binding the contract method 0x0b48636c.
//
// Solidity: function doMath((uint256,uint256)[5] items) pure returns(uint256)
func DoMath(items [5]ExampleNumbers) (*big.Int, error) {
	return _Solidity.Contract.DoMath(&_Solidity.CallOpts, items)
}

```

Whereas the ABI from vyper generates a binding that accepts a list of numbers, not struct/tuple

```go
// DoMath is a free data retrieval call binding the contract method 0x0b48636c.
//
// Solidity: function doMath((uint256,uint256)[5] items) pure returns(uint256)
func DoMath(items [5]*big.Int) (*big.Int, error) {
	return _Vyper.Contract.DoMath(&_Vyper.CallOpts, items)
}
```

## web3.py

Deployed contracts (on polygon testnet)

|compiler   |contract address   |
|-----------|-------------------|
|solidity   ||
|vyper      ||



The following 2 contracts should be the same.

- [Implemented in solidity](contracts/example.sol)
- [Implemented in solidity](contracts/example.vy)

The abis generated

- [vyper.abi.json](abis/vyper.abi.json) : `vyper -f abi  contracts/example.vy` version 0.3.2
- [solidity.abi.json](abis/solidity.abi.json) : `solc --abi ./contracts/example.sol` version 0.8.10+commit.fc410830


