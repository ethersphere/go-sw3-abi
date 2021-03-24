// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sw3abi

const ERC20SimpleSwapABIv0_3_1 = `[
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
    "stateMutability": "nonpayable",
    "type": "constructor"
  },
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
        "internalType": "contract ERC20",
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

const SimpleSwapFactoryABIv0_3_1 = `[
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
  }
]`

const ERC20ABIv0_3_1 = `[
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

const SimpleSwapFactoryBinv0_3_1 = "0x608060405234801561001057600080fd5b50604051612a25380380612a258339818101604052602081101561003357600080fd5b810190808051906020019092919050505080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050612990806100956000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063576d727114610046578063a6021ace146100be578063c70242ad146100f2575b600080fd5b6100926004803603604081101561005c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061014c565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100c6610291565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101346004803603602081101561010857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506102b7565b60405180821515815260200191505060405180910390f35b60008083600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684604051610180906102d7565b808473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604051809103906000f0801580156101e0573d6000803e3d6000fd5b50905060016000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507fc0ffc525a1c7689549d7f79b49eca900e61ac49b43d977f680bcc3b36224c00481604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a18091505092915050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006020528060005260406000206000915054906101000a900460ff1681565b612676806102e58339019056fe608060405234801561001057600080fd5b506040516126763803806126768339818101604052606081101561003357600080fd5b8101908080519060200190929190805190602001909291908051906020019092919050505082600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600081905550505050612583806100f36000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c8063b6343b0d116100b8578063b7ec1a331161007c578063b7ec1a33146106ad578063c49f91d3146106cb578063c76a4d31146106e9578063d4c9a8e814610741578063e0bcf13a14610826578063fc0c546a1461084457610142565b8063b6343b0d14610596578063b648b41714610603578063b69ef8a814610623578063b777035014610641578063b79989071461068f57610142565b80631d1438481161010a5780631d1438481461042c5780632e1a7d4d14610460578063338f3fed1461048e578063488b017c146104dc57806381f03fcb146104fa578063946f46a21461055257610142565b80630d5f265914610147578063121010211461022c5780631357e1dc1461024a57806315c3343f146102685780631633fb1d14610286575b600080fd5b61022a6004803603606081101561015d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001906401000000008111156101a457600080fd5b8201836020820111156101b657600080fd5b803590602001918460018302840111640100000000831117156101d857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610878565b005b61023461088b565b6040518082815260200191505060405180910390f35b610252610891565b6040518082815260200191505060405180910390f35b610270610897565b6040518082815260200191505060405180910390f35b61042a600480360360c081101561029c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561030357600080fd5b82018360208201111561031557600080fd5b8035906020019184600183028401116401000000008311171561033757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190803590602001906401000000008111156103a457600080fd5b8201836020820111156103b657600080fd5b803590602001918460018302840111640100000000831117156103d857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506108bb565b005b610434610969565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61048c6004803603602081101561047657600080fd5b810190808035906020019092919050505061098f565b005b6104da600480360360408110156104a457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610bfa565b005b6104e4610e0f565b6040518082815260200191505060405180910390f35b61053c6004803603602081101561051057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e33565b6040518082815260200191505060405180910390f35b6105946004803603602081101561056857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e4b565b005b6105d8600480360360208110156105ac57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f9e565b6040518085815260200184815260200183815260200182815260200194505050505060405180910390f35b61060b610fce565b60405180821515815260200191505060405180910390f35b61062b610fe1565b6040518082815260200191505060405180910390f35b61068d6004803603604081101561065757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506110ac565b005b610697611294565b6040518082815260200191505060405180910390f35b6106b56112b8565b6040518082815260200191505060405180910390f35b6106d36112db565b6040518082815260200191505060405180910390f35b61072b600480360360208110156106ff57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506112ff565b6040518082815260200191505060405180910390f35b6108246004803603606081101561075757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561079e57600080fd5b8201836020820111156107b057600080fd5b803590602001918460018302840111640100000000831117156107d257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611364565b005b61082e611558565b6040518082815260200191505060405180910390f35b61084c61155e565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610886338484600085611584565b505050565b60005481565b60035481565b7f48ebe6deff4a5ee645c01506a026031e2a945d6f41f1f4e5098ad65347492c1281565b6108d16108cb3033878987611c7a565b84611d38565b73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614610954576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806124b26029913960400191505060405180910390fd5b6109618686868585611584565b505050505050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a52576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f53696d706c65537761703a206e6f74206973737565720000000000000000000081525060200191505060405180910390fd5b610a5a6112b8565b811115610ab2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001806125266028913960400191505060405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610b6757600080fd5b505af1158015610b7b573d6000803e3d6000fd5b505050506040513d6020811015610b9157600080fd5b8101908080519060200190929190505050610bf7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602781526020018061248b6027913960400191505060405180910390fd5b50565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610cbd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f53696d706c65537761703a206e6f74206973737565720000000000000000000081525060200191505060405180910390fd5b610cc5610fe1565b610cda82600554611db690919063ffffffff16565b1115610d31576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260348152602001806123ee6034913960400191505060405180910390fd5b6000600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050610d8b828260000154611db690919063ffffffff16565b8160000181905550610da882600554611db690919063ffffffff16565b600581905550600081600301819055508273ffffffffffffffffffffffffffffffffffffffff167f2506c43272ded05d095b91dbba876e66e46888157d3e078db5691496e96c5fad82600001546040518082815260200191505060405180910390a2505050565b7f7d824962dd0f01520922ea1766c987b1db570cd5db90bdba5ccf5e320607950281565b60026020528060005260406000206000915090505481565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905080600301544210158015610ea757506000816003015414155b610efc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806124666025913960400191505060405180910390fd5b610f1781600101548260000154611e3e90919063ffffffff16565b816000018190555060008160030181905550610f428160010154600554611e3e90919063ffffffff16565b6005819055508173ffffffffffffffffffffffffffffffffffffffff167f2506c43272ded05d095b91dbba876e66e46888157d3e078db5691496e96c5fad82600001546040518082815260200191505060405180910390a25050565b60046020528060005260406000206000915090508060000154908060010154908060020154908060030154905084565b600660149054906101000a900460ff1681565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561106c57600080fd5b505afa158015611080573d6000803e3d6000fd5b505050506040513d602081101561109657600080fd5b8101908080519060200190929190505050905090565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461116f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f53696d706c65537761703a206e6f74206973737565720000000000000000000081525060200191505060405180910390fd5b6000600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050806000015482111561120f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260278152602001806124db6027913960400191505060405180910390fd5b60008082600201541461122657816002015461122a565b6000545b905080420182600301819055508282600101819055508373ffffffffffffffffffffffffffffffffffffffff167fc8305077b495025ec4c1d977b176a762c350bb18cad4666ce1ee85c32b78698a846040518082815260200191505060405180910390a250505050565b7fe95f353750f192082df064ca5142d3a2d6f0bef0f3ffad66d80d8af86b7a749a81565b60006112d66005546112c8610fe1565b611e3e90919063ffffffff16565b905090565b7fc2f8787176b8ac6bf7215b4adcc1e069bf4ab82d9ab1df05a57a91d425935b6e81565b600061135d600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015461134f6112b8565b611db690919063ffffffff16565b9050919050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611427576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f53696d706c65537761703a206e6f74206973737565720000000000000000000081525060200191505060405180910390fd5b61143b611435308585611ec1565b82611d38565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16146114be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806124b26029913960400191505060405180910390fd5b81600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201819055508273ffffffffffffffffffffffffffffffffffffffff167f7b816003a769eb718bd9c66bdbd2dd5827da3f92bc6645276876bd7957b08cf0836040518082815260200191505060405180910390a2505050565b60055481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611693576115ed6115e7308786611f57565b82611d38565b73ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611692576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806125026024913960400191505060405180910390fd5b5b60006116e7600260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205485611e3e90919063ffffffff16565b905060006116fd826116f8896112ff565b611fed565b9050600061174d82600460008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154611fed565b9050848210156117c5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f53696d706c65537761703a2063616e6e6f74207061792063616c6c657200000081525060200191505060405180910390fd5b600081146118845761182281600460008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154611e3e90919063ffffffff16565b600460008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555061187d81600554611e3e90919063ffffffff16565b6005819055505b6118d682600260008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611db690919063ffffffff16565b600260008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061192e82600354611db690919063ffffffff16565b600381905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb886119868886611e3e90919063ffffffff16565b6040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156119d957600080fd5b505af11580156119ed573d6000803e3d6000fd5b505050506040513d6020811015611a0357600080fd5b8101908080519060200190929190505050611a69576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602781526020018061248b6027913960400191505060405180910390fd5b60008514611b9557600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33876040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015611b0457600080fd5b505af1158015611b18573d6000803e3d6000fd5b505050506040513d6020811015611b2e57600080fd5b8101908080519060200190929190505050611b94576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602781526020018061248b6027913960400191505060405180910390fd5b5b3373ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff167f950494fc3642fae5221b6c32e0e45765c95ebb382a04a71b160db0843e74c99f858a8a60405180848152602001838152602001828152602001935050505060405180910390a4818314611c70576001600660146101000a81548160ff0219169083151502179055507f3f4449c047e11092ec54dc0751b6b4817a9162745de856c893a26e611d18ffc460405160405180910390a15b5050505050505050565b60007f7d824962dd0f01520922ea1766c987b1db570cd5db90bdba5ccf5e32060795028686868686604051602001808781526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018373ffffffffffffffffffffffffffffffffffffffff168152602001828152602001965050505050505060405160208183030381529060405280519060200120905095945050505050565b600080611d4b611d46612006565b61209e565b8460405160200180807f190100000000000000000000000000000000000000000000000000000000000081525060020183815260200182815260200192505050604051602081830303815290604052805190602001209050611dad8184612120565b91505092915050565b600080828401905083811015611e34576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600082821115611eb6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525060200191505060405180910390fd5b818303905092915050565b60007fe95f353750f192082df064ca5142d3a2d6f0bef0f3ffad66d80d8af86b7a749a848484604051602001808581526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019450505050506040516020818303038152906040528051906020012090509392505050565b60007f48ebe6deff4a5ee645c01506a026031e2a945d6f41f1f4e5098ad65347492c12848484604051602001808581526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019450505050506040516020818303038152906040528051906020012090509392505050565b6000818310611ffc5781611ffe565b825b905092915050565b61200e6123cc565b600046905060405180606001604052806040518060400160405280600a81526020017f436865717565626f6f6b0000000000000000000000000000000000000000000081525081526020016040518060400160405280600381526020017f312e30000000000000000000000000000000000000000000000000000000000081525081526020018281525091505090565b60007fc2f8787176b8ac6bf7215b4adcc1e069bf4ab82d9ab1df05a57a91d425935b6e826000015180519060200120836020015180519060200120846040015160405160200180858152602001848152602001838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b60006041825114612199576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45434453413a20696e76616c6964207369676e6174757265206c656e6774680081525060200191505060405180910390fd5b60008060006020850151925060408501519150606085015160001a90506121c2868285856121cd565b935050505092915050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c111561224b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806124226022913960400191505060405180910390fd5b601b8460ff1614806122605750601c8460ff16145b6122b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806124446022913960400191505060405180910390fd5b600060018686868660405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015612311573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156123c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f45434453413a20696e76616c6964207369676e6174757265000000000000000081525060200191505060405180910390fd5b80915050949350505050565b6040518060600160405280606081526020016060815260200160008152509056fe53696d706c65537761703a2068617264206465706f7369742063616e6e6f74206265206d6f7265207468616e2062616c616e636545434453413a20696e76616c6964207369676e6174757265202773272076616c756545434453413a20696e76616c6964207369676e6174757265202776272076616c756553696d706c65537761703a206465706f736974206e6f74207965742074696d6564206f757453696d706c65537761703a2053696d706c65537761703a207472616e73666572206661696c656453696d706c65537761703a20696e76616c69642062656e6566696369617279207369676e617475726553696d706c65537761703a2068617264206465706f736974206e6f742073756666696369656e7453696d706c65537761703a20696e76616c696420697373756572207369676e617475726553696d706c65537761703a206c697175696442616c616e6365206e6f742073756666696369656e74a26469706673582212204c2fad3ba5b2869707bfdb3e58a5fda75dd03c980a284d08e08cb68fed97b1ca64736f6c634300060c0033a2646970667358221220a220a5710b2f0d01a60691ce90af8b6a7bcc9d105081c7800abac1baad0e833464736f6c634300060c0033"
const SimpleSwapFactoryDeployedBinv0_3_1 = "0x608060405234801561001057600080fd5b50600436106100415760003560e01c8063576d727114610046578063a6021ace146100be578063c70242ad146100f2575b600080fd5b6100926004803603604081101561005c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061014c565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100c6610291565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101346004803603602081101561010857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506102b7565b60405180821515815260200191505060405180910390f35b60008083600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684604051610180906102d7565b808473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604051809103906000f0801580156101e0573d6000803e3d6000fd5b50905060016000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507fc0ffc525a1c7689549d7f79b49eca900e61ac49b43d977f680bcc3b36224c00481604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a18091505092915050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006020528060005260406000206000915054906101000a900460ff1681565b612676806102e58339019056fe608060405234801561001057600080fd5b506040516126763803806126768339818101604052606081101561003357600080fd5b8101908080519060200190929190805190602001909291908051906020019092919050505082600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600081905550505050612583806100f36000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c8063b6343b0d116100b8578063b7ec1a331161007c578063b7ec1a33146106ad578063c49f91d3146106cb578063c76a4d31146106e9578063d4c9a8e814610741578063e0bcf13a14610826578063fc0c546a1461084457610142565b8063b6343b0d14610596578063b648b41714610603578063b69ef8a814610623578063b777035014610641578063b79989071461068f57610142565b80631d1438481161010a5780631d1438481461042c5780632e1a7d4d14610460578063338f3fed1461048e578063488b017c146104dc57806381f03fcb146104fa578063946f46a21461055257610142565b80630d5f265914610147578063121010211461022c5780631357e1dc1461024a57806315c3343f146102685780631633fb1d14610286575b600080fd5b61022a6004803603606081101561015d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001906401000000008111156101a457600080fd5b8201836020820111156101b657600080fd5b803590602001918460018302840111640100000000831117156101d857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610878565b005b61023461088b565b6040518082815260200191505060405180910390f35b610252610891565b6040518082815260200191505060405180910390f35b610270610897565b6040518082815260200191505060405180910390f35b61042a600480360360c081101561029c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561030357600080fd5b82018360208201111561031557600080fd5b8035906020019184600183028401116401000000008311171561033757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190803590602001906401000000008111156103a457600080fd5b8201836020820111156103b657600080fd5b803590602001918460018302840111640100000000831117156103d857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506108bb565b005b610434610969565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61048c6004803603602081101561047657600080fd5b810190808035906020019092919050505061098f565b005b6104da600480360360408110156104a457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610bfa565b005b6104e4610e0f565b6040518082815260200191505060405180910390f35b61053c6004803603602081101561051057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e33565b6040518082815260200191505060405180910390f35b6105946004803603602081101561056857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e4b565b005b6105d8600480360360208110156105ac57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f9e565b6040518085815260200184815260200183815260200182815260200194505050505060405180910390f35b61060b610fce565b60405180821515815260200191505060405180910390f35b61062b610fe1565b6040518082815260200191505060405180910390f35b61068d6004803603604081101561065757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506110ac565b005b610697611294565b6040518082815260200191505060405180910390f35b6106b56112b8565b6040518082815260200191505060405180910390f35b6106d36112db565b6040518082815260200191505060405180910390f35b61072b600480360360208110156106ff57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506112ff565b6040518082815260200191505060405180910390f35b6108246004803603606081101561075757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561079e57600080fd5b8201836020820111156107b057600080fd5b803590602001918460018302840111640100000000831117156107d257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611364565b005b61082e611558565b6040518082815260200191505060405180910390f35b61084c61155e565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610886338484600085611584565b505050565b60005481565b60035481565b7f48ebe6deff4a5ee645c01506a026031e2a945d6f41f1f4e5098ad65347492c1281565b6108d16108cb3033878987611c7a565b84611d38565b73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614610954576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806124b26029913960400191505060405180910390fd5b6109618686868585611584565b505050505050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a52576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f53696d706c65537761703a206e6f74206973737565720000000000000000000081525060200191505060405180910390fd5b610a5a6112b8565b811115610ab2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001806125266028913960400191505060405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610b6757600080fd5b505af1158015610b7b573d6000803e3d6000fd5b505050506040513d6020811015610b9157600080fd5b8101908080519060200190929190505050610bf7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602781526020018061248b6027913960400191505060405180910390fd5b50565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610cbd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f53696d706c65537761703a206e6f74206973737565720000000000000000000081525060200191505060405180910390fd5b610cc5610fe1565b610cda82600554611db690919063ffffffff16565b1115610d31576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260348152602001806123ee6034913960400191505060405180910390fd5b6000600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050610d8b828260000154611db690919063ffffffff16565b8160000181905550610da882600554611db690919063ffffffff16565b600581905550600081600301819055508273ffffffffffffffffffffffffffffffffffffffff167f2506c43272ded05d095b91dbba876e66e46888157d3e078db5691496e96c5fad82600001546040518082815260200191505060405180910390a2505050565b7f7d824962dd0f01520922ea1766c987b1db570cd5db90bdba5ccf5e320607950281565b60026020528060005260406000206000915090505481565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905080600301544210158015610ea757506000816003015414155b610efc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806124666025913960400191505060405180910390fd5b610f1781600101548260000154611e3e90919063ffffffff16565b816000018190555060008160030181905550610f428160010154600554611e3e90919063ffffffff16565b6005819055508173ffffffffffffffffffffffffffffffffffffffff167f2506c43272ded05d095b91dbba876e66e46888157d3e078db5691496e96c5fad82600001546040518082815260200191505060405180910390a25050565b60046020528060005260406000206000915090508060000154908060010154908060020154908060030154905084565b600660149054906101000a900460ff1681565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561106c57600080fd5b505afa158015611080573d6000803e3d6000fd5b505050506040513d602081101561109657600080fd5b8101908080519060200190929190505050905090565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461116f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f53696d706c65537761703a206e6f74206973737565720000000000000000000081525060200191505060405180910390fd5b6000600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050806000015482111561120f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260278152602001806124db6027913960400191505060405180910390fd5b60008082600201541461122657816002015461122a565b6000545b905080420182600301819055508282600101819055508373ffffffffffffffffffffffffffffffffffffffff167fc8305077b495025ec4c1d977b176a762c350bb18cad4666ce1ee85c32b78698a846040518082815260200191505060405180910390a250505050565b7fe95f353750f192082df064ca5142d3a2d6f0bef0f3ffad66d80d8af86b7a749a81565b60006112d66005546112c8610fe1565b611e3e90919063ffffffff16565b905090565b7fc2f8787176b8ac6bf7215b4adcc1e069bf4ab82d9ab1df05a57a91d425935b6e81565b600061135d600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015461134f6112b8565b611db690919063ffffffff16565b9050919050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611427576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f53696d706c65537761703a206e6f74206973737565720000000000000000000081525060200191505060405180910390fd5b61143b611435308585611ec1565b82611d38565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16146114be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806124b26029913960400191505060405180910390fd5b81600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201819055508273ffffffffffffffffffffffffffffffffffffffff167f7b816003a769eb718bd9c66bdbd2dd5827da3f92bc6645276876bd7957b08cf0836040518082815260200191505060405180910390a2505050565b60055481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611693576115ed6115e7308786611f57565b82611d38565b73ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611692576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806125026024913960400191505060405180910390fd5b5b60006116e7600260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205485611e3e90919063ffffffff16565b905060006116fd826116f8896112ff565b611fed565b9050600061174d82600460008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154611fed565b9050848210156117c5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f53696d706c65537761703a2063616e6e6f74207061792063616c6c657200000081525060200191505060405180910390fd5b600081146118845761182281600460008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154611e3e90919063ffffffff16565b600460008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555061187d81600554611e3e90919063ffffffff16565b6005819055505b6118d682600260008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611db690919063ffffffff16565b600260008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061192e82600354611db690919063ffffffff16565b600381905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb886119868886611e3e90919063ffffffff16565b6040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156119d957600080fd5b505af11580156119ed573d6000803e3d6000fd5b505050506040513d6020811015611a0357600080fd5b8101908080519060200190929190505050611a69576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602781526020018061248b6027913960400191505060405180910390fd5b60008514611b9557600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33876040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015611b0457600080fd5b505af1158015611b18573d6000803e3d6000fd5b505050506040513d6020811015611b2e57600080fd5b8101908080519060200190929190505050611b94576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602781526020018061248b6027913960400191505060405180910390fd5b5b3373ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff167f950494fc3642fae5221b6c32e0e45765c95ebb382a04a71b160db0843e74c99f858a8a60405180848152602001838152602001828152602001935050505060405180910390a4818314611c70576001600660146101000a81548160ff0219169083151502179055507f3f4449c047e11092ec54dc0751b6b4817a9162745de856c893a26e611d18ffc460405160405180910390a15b5050505050505050565b60007f7d824962dd0f01520922ea1766c987b1db570cd5db90bdba5ccf5e32060795028686868686604051602001808781526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018373ffffffffffffffffffffffffffffffffffffffff168152602001828152602001965050505050505060405160208183030381529060405280519060200120905095945050505050565b600080611d4b611d46612006565b61209e565b8460405160200180807f190100000000000000000000000000000000000000000000000000000000000081525060020183815260200182815260200192505050604051602081830303815290604052805190602001209050611dad8184612120565b91505092915050565b600080828401905083811015611e34576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600082821115611eb6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525060200191505060405180910390fd5b818303905092915050565b60007fe95f353750f192082df064ca5142d3a2d6f0bef0f3ffad66d80d8af86b7a749a848484604051602001808581526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019450505050506040516020818303038152906040528051906020012090509392505050565b60007f48ebe6deff4a5ee645c01506a026031e2a945d6f41f1f4e5098ad65347492c12848484604051602001808581526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019450505050506040516020818303038152906040528051906020012090509392505050565b6000818310611ffc5781611ffe565b825b905092915050565b61200e6123cc565b600046905060405180606001604052806040518060400160405280600a81526020017f436865717565626f6f6b0000000000000000000000000000000000000000000081525081526020016040518060400160405280600381526020017f312e30000000000000000000000000000000000000000000000000000000000081525081526020018281525091505090565b60007fc2f8787176b8ac6bf7215b4adcc1e069bf4ab82d9ab1df05a57a91d425935b6e826000015180519060200120836020015180519060200120846040015160405160200180858152602001848152602001838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b60006041825114612199576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45434453413a20696e76616c6964207369676e6174757265206c656e6774680081525060200191505060405180910390fd5b60008060006020850151925060408501519150606085015160001a90506121c2868285856121cd565b935050505092915050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c111561224b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806124226022913960400191505060405180910390fd5b601b8460ff1614806122605750601c8460ff16145b6122b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806124446022913960400191505060405180910390fd5b600060018686868660405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015612311573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156123c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f45434453413a20696e76616c6964207369676e6174757265000000000000000081525060200191505060405180910390fd5b80915050949350505050565b6040518060600160405280606081526020016060815260200160008152509056fe53696d706c65537761703a2068617264206465706f7369742063616e6e6f74206265206d6f7265207468616e2062616c616e636545434453413a20696e76616c6964207369676e6174757265202773272076616c756545434453413a20696e76616c6964207369676e6174757265202776272076616c756553696d706c65537761703a206465706f736974206e6f74207965742074696d6564206f757453696d706c65537761703a2053696d706c65537761703a207472616e73666572206661696c656453696d706c65537761703a20696e76616c69642062656e6566696369617279207369676e617475726553696d706c65537761703a2068617264206465706f736974206e6f742073756666696369656e7453696d706c65537761703a20696e76616c696420697373756572207369676e617475726553696d706c65537761703a206c697175696442616c616e6365206e6f742073756666696369656e74a26469706673582212204c2fad3ba5b2869707bfdb3e58a5fda75dd03c980a284d08e08cb68fed97b1ca64736f6c634300060c0033a2646970667358221220a220a5710b2f0d01a60691ce90af8b6a7bcc9d105081c7800abac1baad0e833464736f6c634300060c0033"
