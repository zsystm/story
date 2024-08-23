package bindings

import (
	_ "embed"
)

const (
	IPTokenSlashingDeployedBytecode = "0x6080604052600436106100dd5760003560e01c806379ba50971161007f578063c4d66de811610059578063c4d66de814610231578063e30c397814610251578063e4dfccd814610266578063f2fde38b1461027957600080fd5b806379ba5097146101c95780638da5cb5b146101de578063ad3cb1cc146101f357600080fd5b806340eda14a116100bb57806340eda14a146101795780634f1ef2861461018c57806352d1902d1461019f578063715018a6146101b457600080fd5b806304ff53ed146100e25780630c863f77146101335780632801f1ec14610155575b600080fd5b3480156100ee57600080fd5b506101167f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561013f57600080fd5b5061015361014e36600461107f565b610299565b005b34801561016157600080fd5b5061016b60005481565b60405190815260200161012a565b610153610187366004611098565b6102dc565b61015361019a366004611195565b61035e565b3480156101ab57600080fd5b5061016b610379565b3480156101c057600080fd5b50610153610396565b3480156101d557600080fd5b506101536103aa565b3480156101ea57600080fd5b506101166103f7565b3480156101ff57600080fd5b50610224604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161012a9190611276565b34801561023d57600080fd5b5061015361024c366004611289565b61042c565b34801561025d57600080fd5b506101166105b8565b610153610274366004611098565b6105e1565b34801561028557600080fd5b50610153610294366004611289565b610724565b6102a16107a9565b60008190556040518181527feac81de2f20162b0540ca5d3f43896af15b471a55729ff0c000e611d8b2723639060200160405180910390a150565b61031b82828080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506107db92505050565b61035a82828080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061096b92505050565b5050565b610366610a37565b61036f82610adc565b61035a8282610ae4565b6000610383610ba6565b506000805160206114e783398151915290565b61039e6107a9565b6103a86000610bef565b565b33806103b46105b8565b6001600160a01b0316146103eb5760405163118cdaa760e01b81526001600160a01b03821660048201526024015b60405180910390fd5b6103f481610bef565b50565b6000807f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005b546001600160a01b031692915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff166000811580156104725750825b905060008267ffffffffffffffff16600114801561048f5750303b155b90508115801561049d575080155b156104bb5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156104e557845460ff60401b1916600160401b1785555b6001600160a01b0386166105595760405162461bcd60e51b815260206004820152603560248201527f4950546f6b656e536c617368696e673a206163636573734d616e616765722063604482015274616e6e6f74206265207a65726f206164647265737360581b60648201526084016103e2565b610561610c27565b61056a86610c2f565b83156105b057845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b6000807f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0061041c565b818133604182146106045760405162461bcd60e51b81526004016103e2906112a4565b82826000818110610617576106176112ea565b9050013560f81c60f81b6001600160f81b031916600460f81b1461064d5760405162461bcd60e51b81526004016103e290611300565b806001600160a01b03166106618484610c40565b6001600160a01b0316146106cf5760405162461bcd60e51b815260206004820152602f60248201527f4950546f6b656e536c617368696e673a20496e76616c6964207075626b65792060448201526e64657269766564206164647265737360881b60648201526084016103e2565b600061071086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610c6f92505050565b905061071b816107db565b6105b08161096b565b61072c6107a9565b7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0080546001600160a01b0319166001600160a01b03831690811782556107706103f7565b6001600160a01b03167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a35050565b336107b26103f7565b6001600160a01b0316146103a85760405163118cdaa760e01b81523360048201526024016103e2565b80516021146107fc5760405162461bcd60e51b81526004016103e2906112a4565b8060008151811061080f5761080f6112ea565b6020910101516001600160f81b031916600160f91b148061085557508060008151811061083e5761083e6112ea565b6020910101516001600160f81b031916600360f81b145b6108715760405162461bcd60e51b81526004016103e290611300565b604051638d3e1e4160e01b81526000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638d3e1e41906108c0908590600401611276565b600060405180830381865afa1580156108dd573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610905919081019061135a565b505050505090508061035a5760405162461bcd60e51b815260206004820152602960248201527f4950546f6b656e536c617368696e673a2056616c696461746f7220646f6573206044820152681b9bdd08195e1a5cdd60ba1b60648201526084016103e2565b60005434146109c65760405162461bcd60e51b815260206004820152602160248201527f4950546f6b656e536c617368696e673a20496e73756666696369656e742066656044820152606560f81b60648201526084016103e2565b6040516000903480156108fc029183818181858288f193505050501580156109f2573d6000803e3d6000fd5b50336001600160a01b03167f4a90ea32527ecacc0f4b32b31f99e4c633a2b4fe81ea7444989e2e68bc9ece3b82604051610a2c9190611276565b60405180910390a250565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610abe57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610ab26000805160206114e7833981519152546001600160a01b031690565b6001600160a01b031614155b156103a85760405163703e46dd60e11b815260040160405180910390fd5b6103f46107a9565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610b3e575060408051601f3d908101601f19168201909252610b3b91810190611426565b60015b610b6657604051634c9c8ce360e01b81526001600160a01b03831660048201526024016103e2565b6000805160206114e78339815191528114610b9757604051632a87526960e21b8152600481018290526024016103e2565b610ba18383610dbb565b505050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103a85760405163703e46dd60e11b815260040160405180910390fd5b7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0080546001600160a01b031916815561035a82610e11565b6103a8610e82565b610c37610e82565b6103f481610ecb565b6000610c4f826001818661143f565b604051610c5d929190611469565b60405190819003902090505b92915050565b60608151604114610cd15760405162461bcd60e51b815260206004820152602660248201527f496e76616c696420756e636f6d70726573736564207075626c6963206b6579206044820152650d8cadccee8d60d31b60648201526084016103e2565b602182015160418301516000610ceb600260ff8416611479565b60ff1615610cfd57600360f81b610d03565b600160f91b5b60408051602180825260608201909252919250600091906020820181803683370190505090508181600081518110610d3d57610d3d6112ea565b60200101906001600160f81b031916908160001a90535060005b6020811015610db157848160208110610d7257610d726112ea565b1a60f81b82610d828360016114a9565b81518110610d9257610d926112ea565b60200101906001600160f81b031916908160001a905350600101610d57565b5095945050505050565b610dc482610efd565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a2805115610e0957610ba18282610f62565b61035a610fd8565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166103a857604051631afcd79f60e31b815260040160405180910390fd5b610ed3610e82565b6001600160a01b0381166103eb57604051631e4fbdf760e01b8152600060048201526024016103e2565b806001600160a01b03163b600003610f3357604051634c9c8ce360e01b81526001600160a01b03821660048201526024016103e2565b6000805160206114e783398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b031684604051610f7f91906114ca565b600060405180830381855af49150503d8060008114610fba576040519150601f19603f3d011682016040523d82523d6000602084013e610fbf565b606091505b5091509150610fcf858383610ff7565b95945050505050565b34156103a85760405163b398979f60e01b815260040160405180910390fd5b60608261100c5761100782611056565b61104f565b815115801561102357506001600160a01b0384163b155b1561104c57604051639996b31560e01b81526001600160a01b03851660048201526024016103e2565b50805b9392505050565b8051156110665780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b60006020828403121561109157600080fd5b5035919050565b600080602083850312156110ab57600080fd5b823567ffffffffffffffff808211156110c357600080fd5b818501915085601f8301126110d757600080fd5b8135818111156110e657600080fd5b8660208285010111156110f857600080fd5b60209290920196919550909350505050565b80356001600160a01b038116811461112157600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561116557611165611126565b604052919050565b600067ffffffffffffffff82111561118757611187611126565b50601f01601f191660200190565b600080604083850312156111a857600080fd5b6111b18361110a565b9150602083013567ffffffffffffffff8111156111cd57600080fd5b8301601f810185136111de57600080fd5b80356111f16111ec8261116d565b61113c565b81815286602083850101111561120657600080fd5b816020840160208301376000602083830101528093505050509250929050565b60005b83811015611241578181015183820152602001611229565b50506000910152565b60008151808452611262816020860160208601611226565b601f01601f19169290920160200192915050565b60208152600061104f602083018461124a565b60006020828403121561129b57600080fd5b61104f8261110a565b60208082526026908201527f4950546f6b656e536c617368696e673a20496e76616c6964207075626b6579206040820152650d8cadccee8d60d31b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60208082526026908201527f4950546f6b656e536c617368696e673a20496e76616c6964207075626b6579206040820152650e0e4caccd2f60d31b606082015260800190565b805163ffffffff8116811461112157600080fd5b60008060008060008060c0878903121561137357600080fd5b8651801515811461138357600080fd5b602088015190965067ffffffffffffffff8111156113a057600080fd5b8701601f810189136113b157600080fd5b80516113bf6111ec8261116d565b8181528a60208385010111156113d457600080fd5b6113e5826020830160208601611226565b809750505050604087015193506113fe60608801611346565b925061140c60808801611346565b915061141a60a08801611346565b90509295509295509295565b60006020828403121561143857600080fd5b5051919050565b6000808585111561144f57600080fd5b8386111561145c57600080fd5b5050820193919092039150565b8183823760009101908152919050565b600060ff83168061149a57634e487b7160e01b600052601260045260246000fd5b8060ff84160691505092915050565b80820180821115610c6957634e487b7160e01b600052601160045260246000fd5b600082516114dc818460208701611226565b919091019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca264697066735822122057127e720484b2a6e7827a9f911d9b8d98ca7ad65cecc0616bfdbda309d9ac4a64736f6c63430008180033"
)

//go:embed iptokenslashing_storage_layout.json
var iptokenslashingStorageLayoutJSON []byte

var IPTokenSlashingStorageLayout = mustGetStorageLayout(iptokenslashingStorageLayoutJSON)
