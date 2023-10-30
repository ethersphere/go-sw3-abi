// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sw3abi

const ERC20SimpleSwapABIv0_5_4 = `[
  {
    "anonymous": false,
    "inputs": [],
    "name": "ChequeBounced",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "beneficiary",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "recipient",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "caller",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "totalPayout",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "cumulativePayout",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "callerPayout",
        "type": "uint256"
      }
    ],
    "name": "ChequeCashed",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "beneficiary",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "amount",
        "type": "uint256"
      }
    ],
    "name": "HardDepositAmountChanged",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "beneficiary",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "decreaseAmount",
        "type": "uint256"
      }
    ],
    "name": "HardDepositDecreasePrepared",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "beneficiary",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "timeout",
        "type": "uint256"
      }
    ],
    "name": "HardDepositTimeoutChanged",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "amount",
        "type": "uint256"
      }
    ],
    "name": "Withdraw",
    "type": "event"
  },
  {
    "inputs": [],
    "name": "CASHOUT_TYPEHASH",
    "outputs": [
      {
        "internalType": "bytes32",
        "name": "",
        "type": "bytes32"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "CHEQUE_TYPEHASH",
    "outputs": [
      {
        "internalType": "bytes32",
        "name": "",
        "type": "bytes32"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "CUSTOMDECREASETIMEOUT_TYPEHASH",
    "outputs": [
      {
        "internalType": "bytes32",
        "name": "",
        "type": "bytes32"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "EIP712DOMAIN_TYPEHASH",
    "outputs": [
      {
        "internalType": "bytes32",
        "name": "",
        "type": "bytes32"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "balance",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "bounced",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "beneficiary",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "recipient",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "cumulativePayout",
        "type": "uint256"
      },
      {
        "internalType": "bytes",
        "name": "beneficiarySig",
        "type": "bytes"
      },
      {
        "internalType": "uint256",
        "name": "callerPayout",
        "type": "uint256"
      },
      {
        "internalType": "bytes",
        "name": "issuerSig",
        "type": "bytes"
      }
    ],
    "name": "cashCheque",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "recipient",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "cumulativePayout",
        "type": "uint256"
      },
      {
        "internalType": "bytes",
        "name": "issuerSig",
        "type": "bytes"
      }
    ],
    "name": "cashChequeBeneficiary",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "beneficiary",
        "type": "address"
      }
    ],
    "name": "decreaseHardDeposit",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "defaultHardDepositTimeout",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "name": "hardDeposits",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "amount",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "decreaseAmount",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "timeout",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "canBeDecreasedAt",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "beneficiary",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "amount",
        "type": "uint256"
      }
    ],
    "name": "increaseHardDeposit",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_issuer",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "_token",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "_defaultHardDepositTimeout",
        "type": "uint256"
      }
    ],
    "name": "init",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "issuer",
    "outputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "liquidBalance",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "beneficiary",
        "type": "address"
      }
    ],
    "name": "liquidBalanceFor",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "name": "paidOut",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "beneficiary",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "decreaseAmount",
        "type": "uint256"
      }
    ],
    "name": "prepareDecreaseHardDeposit",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "beneficiary",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "hardDepositTimeout",
        "type": "uint256"
      },
      {
        "internalType": "bytes",
        "name": "beneficiarySig",
        "type": "bytes"
      }
    ],
    "name": "setCustomHardDepositTimeout",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "token",
    "outputs": [
      {
        "internalType": "contract IERC20",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "totalHardDeposit",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "totalPaidOut",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "amount",
        "type": "uint256"
      }
    ],
    "name": "withdraw",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  }
]`

const SimpleSwapFactoryABIv0_5_4 = `[
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_ERC20Address",
        "type": "address"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "constructor"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "address",
        "name": "contractAddress",
        "type": "address"
      }
    ],
    "name": "SimpleSwapDeployed",
    "type": "event"
  },
  {
    "inputs": [],
    "name": "ERC20Address",
    "outputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "issuer",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "defaultHardDepositTimeoutDuration",
        "type": "uint256"
      },
      {
        "internalType": "bytes32",
        "name": "salt",
        "type": "bytes32"
      }
    ],
    "name": "deploySimpleSwap",
    "outputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "name": "deployedContracts",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "master",
    "outputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
]`

const ERC20ABIv0_5_4 = `[
  {
    "inputs": [
      {
        "internalType": "string",
        "name": "name_",
        "type": "string"
      },
      {
        "internalType": "string",
        "name": "symbol_",
        "type": "string"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "constructor"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "owner",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "spender",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "value",
        "type": "uint256"
      }
    ],
    "name": "Approval",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "from",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "to",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "value",
        "type": "uint256"
      }
    ],
    "name": "Transfer",
    "type": "event"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "owner",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "spender",
        "type": "address"
      }
    ],
    "name": "allowance",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "spender",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "amount",
        "type": "uint256"
      }
    ],
    "name": "approve",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "account",
        "type": "address"
      }
    ],
    "name": "balanceOf",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "decimals",
    "outputs": [
      {
        "internalType": "uint8",
        "name": "",
        "type": "uint8"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "spender",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "subtractedValue",
        "type": "uint256"
      }
    ],
    "name": "decreaseAllowance",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "spender",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "addedValue",
        "type": "uint256"
      }
    ],
    "name": "increaseAllowance",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "name",
    "outputs": [
      {
        "internalType": "string",
        "name": "",
        "type": "string"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "symbol",
    "outputs": [
      {
        "internalType": "string",
        "name": "",
        "type": "string"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "totalSupply",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "recipient",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "amount",
        "type": "uint256"
      }
    ],
    "name": "transfer",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "sender",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "recipient",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "amount",
        "type": "uint256"
      }
    ],
    "name": "transferFrom",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  }
]`

const SimpleSwapFactoryBinv0_5_4 = "0x608060405234801561001057600080fd5b50604051611d90380380611d9083398101604081905261002f91610116565b600180546001600160a01b0319166001600160a01b03831617905560405160009061005990610109565b604051809103906000f080158015610075573d6000803e3d6000fd5b506040516343431f6360e11b81526001600482015260006024820181905260448201529091506001600160a01b038216906386863ec690606401600060405180830381600087803b1580156100c957600080fd5b505af11580156100dd573d6000803e3d6000fd5b5050600280546001600160a01b0319166001600160a01b03949094169390931790925550610146915050565b6118ff8061049183390190565b60006020828403121561012857600080fd5b81516001600160a01b038116811461013f57600080fd5b9392505050565b61033c806101556000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806315efd8a714610051578063a6021ace14610081578063c70242ad14610094578063ee97f7f3146100c7575b600080fd5b61006461005f3660046102b1565b6100da565b6040516001600160a01b0390911681526020015b60405180910390f35b600154610064906001600160a01b031681565b6100b76100a23660046102e4565b60006020819052908152604090205460ff1681565b6040519015158152602001610078565b600254610064906001600160a01b031681565b60025460408051336020820152908101839052600091829161011e916001600160a01b031690606001604051602081830303815290604052805190602001206101eb565b6001546040516343431f6360e11b81526001600160a01b0388811660048301529182166024820152604481018790529192508216906386863ec690606401600060405180830381600087803b15801561017657600080fd5b505af115801561018a573d6000803e3d6000fd5b505050506001600160a01b03811660008181526020818152604091829020805460ff1916600117905590519182527fc0ffc525a1c7689549d7f79b49eca900e61ac49b43d977f680bcc3b36224c004910160405180910390a1949350505050565b6000604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528360601b60148201526e5af43d82803e903d91602b57fd5bf360881b6028820152826037826000f59150506001600160a01b03811661028f5760405162461bcd60e51b815260206004820152601760248201527f455243313136373a2063726561746532206661696c6564000000000000000000604482015260640160405180910390fd5b92915050565b80356001600160a01b03811681146102ac57600080fd5b919050565b6000806000606084860312156102c657600080fd5b6102cf84610295565b95602085013595506040909401359392505050565b6000602082840312156102f657600080fd5b6102ff82610295565b939250505056fea26469706673582212203f76e42322bce6a35f6ef0357cad6785061a54bbbe936b5cc1a425f13408ffbc64736f6c63430008130033608060405234801561001057600080fd5b506118df806100206000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c8063946f46a2116100c3578063b7ec1a331161007c578063b7ec1a331461033f578063c49f91d314610347578063c76a4d311461036e578063d4c9a8e814610381578063e0bcf13a14610394578063fc0c546a1461039d57600080fd5b8063946f46a214610271578063b6343b0d14610284578063b648b417146102d9578063b69ef8a8146102fd578063b777035014610305578063b79989071461031857600080fd5b80631d143848116101155780631d143848146101c65780632e1a7d4d146101f1578063338f3fed14610204578063488b017c1461021757806381f03fcb1461023e57806386863ec61461025e57600080fd5b80630d5f26591461015257806312101021146101675780631357e1dc1461018357806315c3343f1461018c5780631633fb1d146101b3575b600080fd5b61016561016036600461166b565b6103b0565b005b61017060005481565b6040519081526020015b60405180910390f35b61017060035481565b6101707f48ebe6deff4a5ee645c01506a026031e2a945d6f41f1f4e5098ad65347492c1281565b6101656101c13660046116c2565b6103c3565b6006546101d9906001600160a01b031681565b6040516001600160a01b03909116815260200161017a565b6101656101ff366004611759565b6104b5565b610165610212366004611772565b6105eb565b6101707f7d824962dd0f01520922ea1766c987b1db570cd5db90bdba5ccf5e320607950281565b61017061024c36600461179c565b60026020526000908152604090205481565b61016561026c3660046117b7565b610727565b61016561027f36600461179c565b6107f2565b6102b961029236600461179c565b60046020526000908152604090208054600182015460028301546003909301549192909184565b60408051948552602085019390935291830152606082015260800161017a565b6006546102ed90600160a01b900460ff1681565b604051901515815260200161017a565b6101706108e0565b610165610313366004611772565b610952565b6101707fe95f353750f192082df064ca5142d3a2d6f0bef0f3ffad66d80d8af86b7a749a81565b610170610a86565b6101707fc2f8787176b8ac6bf7215b4adcc1e069bf4ab82d9ab1df05a57a91d425935b6e81565b61017061037c36600461179c565b610a9d565b61016561038f36600461166b565b610ace565b61017060055481565b6001546101d9906001600160a01b031681565b6103be338484600085610bd9565b505050565b604080517f7d824962dd0f01520922ea1766c987b1db570cd5db90bdba5ccf5e32060795026020808301919091523082840152336060830152608082018790526001600160a01b03881660a083015260c08083018690528351808403909101815260e0909201909252805191012061043b9084611023565b6001600160a01b0316866001600160a01b0316146104a05760405162461bcd60e51b815260206004820152601d60248201527f696e76616c69642062656e6566696369617279207369676e617475726500000060448201526064015b60405180910390fd5b6104ad8686868585610bd9565b505050505050565b6006546001600160a01b031633146104fc5760405162461bcd60e51b815260206004820152600a6024820152693737ba1034b9b9bab2b960b11b6044820152606401610497565b610504610a86565b8111156105535760405162461bcd60e51b815260206004820152601c60248201527f6c697175696442616c616e6365206e6f742073756666696369656e74000000006044820152606401610497565b60015460065460405163a9059cbb60e01b81526001600160a01b0391821660048201526024810184905291169063a9059cbb906044016020604051808303816000875af11580156105a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105cc91906117f3565b6105e85760405162461bcd60e51b815260040161049790611815565b50565b6006546001600160a01b0316331461063e5760405162461bcd60e51b815260206004820152601660248201527529b4b6b83632a9bbb0b81d103737ba1034b9b9bab2b960511b6044820152606401610497565b6106466108e0565b816005546106549190611854565b11156106a25760405162461bcd60e51b815260206004820152601c60248201527f68617264206465706f73697420657863656564732062616c616e6365000000006044820152606401610497565b6001600160a01b038216600090815260046020526040902080546106c7908390611854565b81556005546106d7908390611854565b6005556000600382015580546040519081526001600160a01b038416907f2506c43272ded05d095b91dbba876e66e46888157d3e078db5691496e96c5fad906020015b60405180910390a2505050565b6001600160a01b03831661076e5760405162461bcd60e51b815260206004820152600e60248201526d34b73b30b634b21034b9b9bab2b960911b6044820152606401610497565b6006546001600160a01b0316156107bd5760405162461bcd60e51b8152602060048201526013602482015272185b1c9958591e481a5b9a5d1a585b1a5e9959606a1b6044820152606401610497565b600680546001600160a01b039485166001600160a01b0319918216179091556001805493909416921691909117909155600055565b6001600160a01b0381166000908152600460205260409020600381015442108015906108215750600381015415155b61086d5760405162461bcd60e51b815260206004820152601960248201527f6465706f736974206e6f74207965742074696d6564206f7574000000000000006044820152606401610497565b6001810154815461087e9190611867565b81556000600382015560018101546005546108999190611867565b60055580546040519081526001600160a01b038316907f2506c43272ded05d095b91dbba876e66e46888157d3e078db5691496e96c5fad9060200160405180910390a25050565b6001546040516370a0823160e01b81523060048201526000916001600160a01b0316906370a0823190602401602060405180830381865afa158015610929573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061094d919061187a565b905090565b6006546001600160a01b031633146109a55760405162461bcd60e51b815260206004820152601660248201527529b4b6b83632a9bbb0b81d103737ba1034b9b9bab2b960511b6044820152606401610497565b6001600160a01b03821660009081526004602052604090208054821115610a0e5760405162461bcd60e51b815260206004820152601b60248201527f68617264206465706f736974206e6f742073756666696369656e7400000000006044820152606401610497565b60008160020154600014610a26578160020154610a2a565b6000545b9050610a368142611854565b6003830155600182018390556040518381526001600160a01b038516907fc8305077b495025ec4c1d977b176a762c350bb18cad4666ce1ee85c32b78698a9060200160405180910390a250505050565b6000600554610a936108e0565b61094d9190611867565b6001600160a01b038116600090815260046020526040812054610abe610a86565b610ac89190611854565b92915050565b6006546001600160a01b03163314610b155760405162461bcd60e51b815260206004820152600a6024820152693737ba1034b9b9bab2b960b11b6044820152606401610497565b610b29610b233085856110f3565b82611023565b6001600160a01b0316836001600160a01b031614610b895760405162461bcd60e51b815260206004820152601d60248201527f696e76616c69642062656e6566696369617279207369676e61747572650000006044820152606401610497565b6001600160a01b03831660008181526004602052604090819020600201849055517f7b816003a769eb718bd9c66bdbd2dd5827da3f92bc6645276876bd7957b08cf09061071a9085815260200190565b6006546001600160a01b03163314610c5657610bf9610b23308786611164565b6006546001600160a01b03908116911614610c565760405162461bcd60e51b815260206004820152601860248201527f696e76616c696420697373756572207369676e617475726500000000000000006044820152606401610497565b6001600160a01b038516600090815260026020526040812054610c799085611867565b90506000610c8f82610c8a89610a9d565b6111ba565b6001600160a01b03881660009081526004602052604081205491925090610cb79083906111ba565b905084821015610d095760405162461bcd60e51b815260206004820152601d60248201527f53696d706c65537761703a2063616e6e6f74207061792063616c6c65720000006044820152606401610497565b8015610d5e576001600160a01b038816600090815260046020526040902054610d33908290611867565b6001600160a01b038916600090815260046020526040902055600554610d5a908290611867565b6005555b6001600160a01b038816600090815260026020526040902054610d82908390611854565b6001600160a01b038916600090815260026020526040902055600354610da9908390611854565b600355828214610df0576006805460ff60a01b1916600160a01b1790556040517f3f4449c047e11092ec54dc0751b6b4817a9162745de856c893a26e611d18ffc490600090a15b8415610f335760015460405163a9059cbb60e01b8152336004820152602481018790526001600160a01b039091169063a9059cbb906044016020604051808303816000875af1158015610e47573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e6b91906117f3565b610e875760405162461bcd60e51b815260040161049790611815565b6001546001600160a01b031663a9059cbb88610ea38886611867565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af1158015610eee573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f1291906117f3565b610f2e5760405162461bcd60e51b815260040161049790611815565b610fc6565b60015460405163a9059cbb60e01b81526001600160a01b038981166004830152602482018590529091169063a9059cbb906044016020604051808303816000875af1158015610f86573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610faa91906117f3565b610fc65760405162461bcd60e51b815260040161049790611815565b604080518381526020810188905290810186905233906001600160a01b0389811691908b16907f950494fc3642fae5221b6c32e0e45765c95ebb382a04a71b160db0843e74c99f9060600160405180910390a45050505050505050565b6000806110a96110316111d2565b805180516020918201208183015180519083012060408085015181517fc2f8787176b8ac6bf7215b4adcc1e069bf4ab82d9ab1df05a57a91d425935b6e95810195909552908401929092526060830152608082015260009060a001604051602081830303815290604052805190602001209050919050565b60405161190160f01b60208201526022810191909152604281018590526062016040516020818303038152906040528051906020012090506110eb8184611248565b949350505050565b604080517fe95f353750f192082df064ca5142d3a2d6f0bef0f3ffad66d80d8af86b7a749a60208201526001600160a01b038086169282019290925290831660608201526080810182905260009060a0015b6040516020818303038152906040528051906020012090509392505050565b604080517f48ebe6deff4a5ee645c01506a026031e2a945d6f41f1f4e5098ad65347492c1260208201526001600160a01b038086169282019290925290831660608201526080810182905260009060a001611145565b60008183106111c957816111cb565b825b9392505050565b6111f660405180606001604052806060815260200160608152602001600081525090565b506040805160a081018252600a6060820190815269436865717565626f6f6b60b01b608083015281528151808301835260038152620312e360ec1b602082810191909152820152469181019190915290565b6000806000611257858561126c565b91509150611264816112da565b509392505050565b60008082516041036112a25760208301516040840151606085015160001a61129687828585611490565b945094505050506112d3565b82516040036112cb57602083015160408401516112c086838361157d565b9350935050506112d3565b506000905060025b9250929050565b60008160048111156112ee576112ee611893565b036112f65750565b600181600481111561130a5761130a611893565b036113575760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610497565b600281600481111561136b5761136b611893565b036113b85760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610497565b60038160048111156113cc576113cc611893565b036114245760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610497565b600481600481111561143857611438611893565b036105e85760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610497565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156114c75750600090506003611574565b8460ff16601b141580156114df57508460ff16601c14155b156114f05750600090506004611574565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611544573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661156d57600060019250925050611574565b9150600090505b94509492505050565b6000806001600160ff1b03831660ff84901c601b0161159e87828885611490565b935093505050935093915050565b80356001600160a01b03811681146115c357600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126115ef57600080fd5b813567ffffffffffffffff8082111561160a5761160a6115c8565b604051601f8301601f19908116603f01168101908282118183101715611632576116326115c8565b8160405283815286602085880101111561164b57600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060006060848603121561168057600080fd5b611689846115ac565b925060208401359150604084013567ffffffffffffffff8111156116ac57600080fd5b6116b8868287016115de565b9150509250925092565b60008060008060008060c087890312156116db57600080fd5b6116e4876115ac565b95506116f2602088016115ac565b945060408701359350606087013567ffffffffffffffff8082111561171657600080fd5b6117228a838b016115de565b94506080890135935060a089013591508082111561173f57600080fd5b5061174c89828a016115de565b9150509295509295509295565b60006020828403121561176b57600080fd5b5035919050565b6000806040838503121561178557600080fd5b61178e836115ac565b946020939093013593505050565b6000602082840312156117ae57600080fd5b6111cb826115ac565b6000806000606084860312156117cc57600080fd5b6117d5846115ac565b92506117e3602085016115ac565b9150604084013590509250925092565b60006020828403121561180557600080fd5b815180151581146111cb57600080fd5b6020808252600f908201526e1d1c985b9cd9995c8819985a5b1959608a1b604082015260600190565b634e487b7160e01b600052601160045260246000fd5b80820180821115610ac857610ac861183e565b81810381811115610ac857610ac861183e565b60006020828403121561188c57600080fd5b5051919050565b634e487b7160e01b600052602160045260246000fdfea2646970667358221220e7a6fb62d773ba85f997655dd2fe29f2d8180eac7cdbc266d60046566974e58264736f6c63430008130033"
const SimpleSwapFactoryDeployedBinv0_5_4 = "0x608060405234801561001057600080fd5b506004361061004c5760003560e01c806315efd8a714610051578063a6021ace14610081578063c70242ad14610094578063ee97f7f3146100c7575b600080fd5b61006461005f3660046102b1565b6100da565b6040516001600160a01b0390911681526020015b60405180910390f35b600154610064906001600160a01b031681565b6100b76100a23660046102e4565b60006020819052908152604090205460ff1681565b6040519015158152602001610078565b600254610064906001600160a01b031681565b60025460408051336020820152908101839052600091829161011e916001600160a01b031690606001604051602081830303815290604052805190602001206101eb565b6001546040516343431f6360e11b81526001600160a01b0388811660048301529182166024820152604481018790529192508216906386863ec690606401600060405180830381600087803b15801561017657600080fd5b505af115801561018a573d6000803e3d6000fd5b505050506001600160a01b03811660008181526020818152604091829020805460ff1916600117905590519182527fc0ffc525a1c7689549d7f79b49eca900e61ac49b43d977f680bcc3b36224c004910160405180910390a1949350505050565b6000604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528360601b60148201526e5af43d82803e903d91602b57fd5bf360881b6028820152826037826000f59150506001600160a01b03811661028f5760405162461bcd60e51b815260206004820152601760248201527f455243313136373a2063726561746532206661696c6564000000000000000000604482015260640160405180910390fd5b92915050565b80356001600160a01b03811681146102ac57600080fd5b919050565b6000806000606084860312156102c657600080fd5b6102cf84610295565b95602085013595506040909401359392505050565b6000602082840312156102f657600080fd5b6102ff82610295565b939250505056fea26469706673582212203f76e42322bce6a35f6ef0357cad6785061a54bbbe936b5cc1a425f13408ffbc64736f6c63430008130033"
