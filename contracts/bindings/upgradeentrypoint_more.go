package bindings

import (
	_ "embed"
)

const (
	UpgradeEntrypointDeployedBytecode = "0x608060405234801561001057600080fd5b50600436106100625760003560e01c8063715018a61461006757806379ba5097146100715780638da5cb5b14610079578063e30c3978146100a2578063ef176e0e146100b3578063f2fde38b146100c6575b600080fd5b61006f6100d9565b005b61006f6100ed565b6000546001600160a01b03165b6040516001600160a01b03909116815260200160405180910390f35b6001546001600160a01b0316610086565b61006f6100c13660046102cf565b610136565b61006f6100d436600461035b565b610184565b6100e16101f5565b6100eb6000610222565b565b60015433906001600160a01b0316811461012a5760405163118cdaa760e01b81526001600160a01b03821660048201526024015b60405180910390fd5b61013381610222565b50565b61013e6101f5565b7f112749e79b2098b58eab36c21f123b2883c3ecbbb4f41623a744fa6d9b3e37c685858585856040516101759594939291906103b4565b60405180910390a15050505050565b61018c6101f5565b600180546001600160a01b0383166001600160a01b031990911681179091556101bd6000546001600160a01b031690565b6001600160a01b03167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b6000546001600160a01b031633146100eb5760405163118cdaa760e01b8152336004820152602401610121565b600180546001600160a01b031916905561013381600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60008083601f84011261029857600080fd5b50813567ffffffffffffffff8111156102b057600080fd5b6020830191508360208285010111156102c857600080fd5b9250929050565b6000806000806000606086880312156102e757600080fd5b853567ffffffffffffffff808211156102ff57600080fd5b61030b89838a01610286565b909750955060208801359150600782900b821461032757600080fd5b9093506040870135908082111561033d57600080fd5b5061034a88828901610286565b969995985093965092949392505050565b60006020828403121561036d57600080fd5b81356001600160a01b038116811461038457600080fd5b9392505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6060815260006103c860608301878961038b565b8560070b602084015282810360408401526103e481858761038b565b9897505050505050505056fea2646970667358221220b5368e57893f3d7629f1e008c242a82e4675b883e77bab8dd19a2293fdfbae5464736f6c63430008180033"
)

//go:embed upgradeentrypoint_storage_layout.json
var upgradeentrypointStorageLayoutJSON []byte

var UpgradeEntrypointStorageLayout = mustGetStorageLayout(upgradeentrypointStorageLayoutJSON)
