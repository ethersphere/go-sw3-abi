// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sw3abi

const ERC20SimpleSwapABIv0_6_7 = `[
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

const SimpleSwapFactoryABIv0_6_7 = `[
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

const ERC20ABIv0_6_7 = `[
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

const SimpleSwapFactoryBinv0_6_7 = "0x608060405234801561001057600080fd5b50604051611df2380380611df283398101604081905261002f91610112565b600180546001600160a01b0319166001600160a01b03831617905560405160009061005990610105565b604051809103906000f080158015610075573d6000803e3d6000fd5b506040516343431f6360e11b81529091506001600160a01b038216906386863ec6906100ab906001906000908190600401610140565b600060405180830381600087803b1580156100c557600080fd5b505af11580156100d9573d6000803e3d6000fd5b5050600280546001600160a01b0319166001600160a01b03949094169390931790925550610164915050565b6118e58061050d83390190565b600060208284031215610123578081fd5b81516001600160a01b0381168114610139578182fd5b9392505050565b6001600160a01b039384168152919092166020820152604081019190915260600190565b61039a806101736000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806315efd8a714610051578063a6021ace1461007a578063c70242ad14610082578063ee97f7f3146100a2575b600080fd5b61006461005f3660046102d6565b6100aa565b6040516100719190610308565b60405180910390f35b6100646101bb565b6100956100903660046102b5565b6101ca565b6040516100719190610359565b6100646101df565b60025460405160009182916100ee916001600160a01b0316906100d3903390879060200161031c565b604051602081830303815290604052805190602001206101ee565b6001546040516343431f6360e11b81529192506001600160a01b03808416926386863ec692610126928a929116908990600401610335565b600060405180830381600087803b15801561014057600080fd5b505af1158015610154573d6000803e3d6000fd5b5050506001600160a01b03821660009081526020819052604090819020805460ff19166001179055517fc0ffc525a1c7689549d7f79b49eca900e61ac49b43d977f680bcc3b36224c00491506101ab908390610308565b60405180910390a1949350505050565b6001546001600160a01b031681565b60006020819052908152604090205460ff1681565b6002546001600160a01b031681565b6000604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528360601b60148201526e5af43d82803e903d91602b57fd5bf360881b6028820152826037826000f59150506001600160a01b038116610293576040805162461bcd60e51b815260206004820152601760248201527f455243313136373a2063726561746532206661696c6564000000000000000000604482015290519081900360640190fd5b92915050565b80356001600160a01b03811681146102b057600080fd5b919050565b6000602082840312156102c6578081fd5b6102cf82610299565b9392505050565b6000806000606084860312156102ea578182fd5b6102f384610299565b95602085013595506040909401359392505050565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b90151581526020019056fea2646970667358221220879500e712d6cdee6e7b2cc582fef56f09aa010787a01406fe50796b05d33cba64736f6c63430007060033608060405234801561001057600080fd5b506118c5806100206000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c8063946f46a2116100c3578063b7ec1a331161007c578063b7ec1a331461027f578063c49f91d314610287578063c76a4d311461028f578063d4c9a8e8146102a2578063e0bcf13a146102b5578063fc0c546a146102bd5761014d565b8063946f46a214610211578063b6343b0d14610224578063b648b41714610247578063b69ef8a81461025c578063b777035014610264578063b7998907146102775761014d565b80631d143848116101155780631d143848146101a85780632e1a7d4d146101bd578063338f3fed146101d0578063488b017c146101e357806381f03fcb146101eb57806386863ec6146101fe5761014d565b80630d5f26591461015257806312101021146101675780631357e1dc1461018557806315c3343f1461018d5780631633fb1d14610195575b600080fd5b61016561016036600461146b565b6102c5565b005b61016f6102d8565b60405161017c9190611563565b60405180910390f35b61016f6102de565b61016f6102e4565b6101656101a33660046113ae565b610308565b6101b061036c565b60405161017c919061152b565b6101656101cb3660046114e0565b61037b565b6101656101de366004611442565b610473565b61016f610556565b61016f6101f9366004611359565b61057a565b61016561020c366004611373565b61058c565b61016561021f366004611359565b610610565b610237610232366004611359565b6106d0565b60405161017c94939291906115c7565b61024f6106f7565b60405161017c9190611558565b61016f610707565b610165610272366004611442565b61078d565b61016f610867565b61016f61088b565b61016f6108a1565b61016f61029d366004611359565b6108c5565b6101656102b036600461146b565b6108f8565b61016f6109b4565b6101b06109ba565b6102d33384846000856109c9565b505050565b60005481565b60035481565b7f48ebe6deff4a5ee645c01506a026031e2a945d6f41f1f4e5098ad65347492c1281565b61031e6103183033878987610dd2565b84610e30565b6001600160a01b0316866001600160a01b0316146103575760405162461bcd60e51b815260040161034e906117d5565b60405180910390fd5b61036486868685856109c9565b505050505050565b6006546001600160a01b031681565b6006546001600160a01b031633146103a55760405162461bcd60e51b815260040161034e906115e2565b6103ad61088b565b8111156103cc5760405162461bcd60e51b815260040161034e90611695565b60015460065460405163a9059cbb60e01b81526001600160a01b039283169263a9059cbb9261040292911690859060040161153f565b602060405180830381600087803b15801561041c57600080fd5b505af1158015610430573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061045491906114c0565b6104705760405162461bcd60e51b815260040161034e9061180c565b50565b6006546001600160a01b0316331461049d5760405162461bcd60e51b815260040161034e90611665565b6104a5610707565b6005546104b29083610e7f565b11156104d05760405162461bcd60e51b815260040161034e90611730565b6001600160a01b038216600090815260046020526040902080546104f49083610e7f565b81556005546105039083610e7f565b6005556000600382015580546040516001600160a01b038516917f2506c43272ded05d095b91dbba876e66e46888157d3e078db5691496e96c5fad916105499190611563565b60405180910390a2505050565b7f7d824962dd0f01520922ea1766c987b1db570cd5db90bdba5ccf5e320607950281565b60026020526000908152604090205481565b6001600160a01b0383166105b25760405162461bcd60e51b815260040161034e9061163d565b6006546001600160a01b0316156105db5760405162461bcd60e51b815260040161034e906116cc565b600680546001600160a01b039485166001600160a01b0319918216179091556001805493909416921691909117909155600055565b6001600160a01b03811660009081526004602052604090206003810154421080159061063f5750600381015415155b61065b5760405162461bcd60e51b815260040161034e90611606565b6001810154815461066b91610ee0565b815560006003820155600181015460055461068591610ee0565b60055580546040516001600160a01b038416917f2506c43272ded05d095b91dbba876e66e46888157d3e078db5691496e96c5fad916106c49190611563565b60405180910390a25050565b60046020526000908152604090208054600182015460028301546003909301549192909184565b600654600160a01b900460ff1681565b6001546040516370a0823160e01b81526000916001600160a01b0316906370a082319061073890309060040161152b565b60206040518083038186803b15801561075057600080fd5b505afa158015610764573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061078891906114f8565b905090565b6006546001600160a01b031633146107b75760405162461bcd60e51b815260040161034e90611665565b6001600160a01b038216600090815260046020526040902080548211156107f05760405162461bcd60e51b815260040161034e906116f9565b6000816002015460001461080857816002015461080c565b6000545b4281016003840155600183018490556040519091506001600160a01b038516907fc8305077b495025ec4c1d977b176a762c350bb18cad4666ce1ee85c32b78698a90610859908690611563565b60405180910390a250505050565b7fe95f353750f192082df064ca5142d3a2d6f0bef0f3ffad66d80d8af86b7a749a81565b600061078860055461089b610707565b90610ee0565b7fc2f8787176b8ac6bf7215b4adcc1e069bf4ab82d9ab1df05a57a91d425935b6e81565b6001600160a01b0381166000908152600460205260408120546108f0906108ea61088b565b90610e7f565b90505b919050565b6006546001600160a01b031633146109225760405162461bcd60e51b815260040161034e906115e2565b610936610930308585610f3d565b82610e30565b6001600160a01b0316836001600160a01b0316146109665760405162461bcd60e51b815260040161034e906117d5565b6001600160a01b03831660008181526004602052604090819020600201849055517f7b816003a769eb718bd9c66bdbd2dd5827da3f92bc6645276876bd7957b08cf090610549908590611563565b60055481565b6001546001600160a01b031681565b6006546001600160a01b03163314610a16576109e9610930308786610f95565b6006546001600160a01b03908116911614610a165760405162461bcd60e51b815260040161034e9061179e565b6001600160a01b038516600090815260026020526040812054610a3a908590610ee0565b90506000610a5082610a4b896108c5565b610fce565b6001600160a01b03881660009081526004602052604081205491925090610a78908390610fce565b905084821015610a9a5760405162461bcd60e51b815260040161034e90611767565b8015610aed576001600160a01b038816600090815260046020526040902054610ac39082610ee0565b6001600160a01b038916600090815260046020526040902055600554610ae99082610ee0565b6005555b6001600160a01b038816600090815260026020526040902054610b109083610e7f565b6001600160a01b038916600090815260026020526040902055600354610b369083610e7f565b600355828214610b7d576006805460ff60a01b1916600160a01b1790556040517f3f4449c047e11092ec54dc0751b6b4817a9162745de856c893a26e611d18ffc490600090a15b8415610ccf5760015460405163a9059cbb60e01b81526001600160a01b039091169063a9059cbb90610bb5903390899060040161153f565b602060405180830381600087803b158015610bcf57600080fd5b505af1158015610be3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c0791906114c0565b610c235760405162461bcd60e51b815260040161034e9061180c565b6001546001600160a01b031663a9059cbb88610c3f8589610ee0565b6040518363ffffffff1660e01b8152600401610c5c92919061153f565b602060405180830381600087803b158015610c7657600080fd5b505af1158015610c8a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cae91906114c0565b610cca5760405162461bcd60e51b815260040161034e9061180c565b610d6f565b60015460405163a9059cbb60e01b81526001600160a01b039091169063a9059cbb90610d01908a90869060040161153f565b602060405180830381600087803b158015610d1b57600080fd5b505af1158015610d2f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d5391906114c0565b610d6f5760405162461bcd60e51b815260040161034e9061180c565b336001600160a01b0316876001600160a01b0316896001600160a01b03167f950494fc3642fae5221b6c32e0e45765c95ebb382a04a71b160db0843e74c99f858a8a604051610dc093929190611835565b60405180910390a45050505050505050565b60007f7d824962dd0f01520922ea1766c987b1db570cd5db90bdba5ccf5e32060795028686868686604051602001610e0f96959493929190611591565b60405160208183030381529060405280519060200120905095945050505050565b600080610e43610e3e610fe4565b61103e565b84604051602001610e55929190611510565b604051602081830303815290604052805190602001209050610e7781846110ae565b949350505050565b600082820183811015610ed9576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b600082821115610f37576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b60007fe95f353750f192082df064ca5142d3a2d6f0bef0f3ffad66d80d8af86b7a749a848484604051602001610f76949392919061156c565b6040516020818303038152906040528051906020012090509392505050565b60007f48ebe6deff4a5ee645c01506a026031e2a945d6f41f1f4e5098ad65347492c12848484604051602001610f76949392919061156c565b6000818310610fdd5781610ed9565b5090919050565b610fec6112ac565b506040805160a081018252600a6060820190815269436865717565626f6f6b60b01b608083015281528151808301835260038152620312e360ec1b602082810191909152820152469181019190915290565b60007fc2f8787176b8ac6bf7215b4adcc1e069bf4ab82d9ab1df05a57a91d425935b6e826000015180519060200120836020015180519060200120846040015160405160200161109194939291906115c7565b604051602081830303815290604052805190602001209050919050565b60008151604114611106576040805162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015290519081900360640190fd5b60208201516040830151606084015160001a6111248682858561112e565b9695505050505050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561118f5760405162461bcd60e51b815260040180806020018281038252602281526020018061184c6022913960400191505060405180910390fd5b8360ff16601b14806111a457508360ff16601c145b6111df5760405162461bcd60e51b815260040180806020018281038252602281526020018061186e6022913960400191505060405180910390fd5b600060018686868660405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa15801561123b573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166112a3576040805162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015290519081900360640190fd5b95945050505050565b60405180606001604052806060815260200160608152602001600081525090565b80356001600160a01b03811681146108f357600080fd5b600082601f8301126112f4578081fd5b813567ffffffffffffffff8082111561130957fe5b604051601f8301601f19168101602001828111828210171561132757fe5b60405282815284830160200186101561133e578384fd5b82602086016020830137918201602001929092529392505050565b60006020828403121561136a578081fd5b610ed9826112cd565b600080600060608486031215611387578182fd5b611390846112cd565b925061139e602085016112cd565b9150604084013590509250925092565b60008060008060008060c087890312156113c6578182fd5b6113cf876112cd565b95506113dd602088016112cd565b945060408701359350606087013567ffffffffffffffff80821115611400578384fd5b61140c8a838b016112e4565b94506080890135935060a0890135915080821115611428578283fd5b5061143589828a016112e4565b9150509295509295509295565b60008060408385031215611454578182fd5b61145d836112cd565b946020939093013593505050565b60008060006060848603121561147f578283fd5b611488846112cd565b925060208401359150604084013567ffffffffffffffff8111156114aa578182fd5b6114b6868287016112e4565b9150509250925092565b6000602082840312156114d1578081fd5b81518015158114610ed9578182fd5b6000602082840312156114f1578081fd5b5035919050565b600060208284031215611509578081fd5b5051919050565b61190160f01b81526002810192909252602282015260420190565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b901515815260200190565b90815260200190565b9384526001600160a01b03928316602085015291166040830152606082015260800190565b9586526001600160a01b03948516602087015292841660408601526060850191909152909116608083015260a082015260c00190565b93845260208401929092526040830152606082015260800190565b6020808252600a90820152693737ba1034b9b9bab2b960b11b604082015260600190565b60208082526019908201527f6465706f736974206e6f74207965742074696d6564206f757400000000000000604082015260600190565b6020808252600e908201526d34b73b30b634b21034b9b9bab2b960911b604082015260600190565b60208082526016908201527529b4b6b83632a9bbb0b81d103737ba1034b9b9bab2b960511b604082015260600190565b6020808252601c908201527f6c697175696442616c616e6365206e6f742073756666696369656e7400000000604082015260600190565b602080825260139082015272185b1c9958591e481a5b9a5d1a585b1a5e9959606a1b604082015260600190565b6020808252601b908201527f68617264206465706f736974206e6f742073756666696369656e740000000000604082015260600190565b6020808252601c908201527f68617264206465706f73697420657863656564732062616c616e636500000000604082015260600190565b6020808252601d908201527f53696d706c65537761703a2063616e6e6f74207061792063616c6c6572000000604082015260600190565b60208082526018908201527f696e76616c696420697373756572207369676e61747572650000000000000000604082015260600190565b6020808252601d908201527f696e76616c69642062656e6566696369617279207369676e6174757265000000604082015260600190565b6020808252600f908201526e1d1c985b9cd9995c8819985a5b1959608a1b604082015260600190565b928352602083019190915260408201526060019056fe45434453413a20696e76616c6964207369676e6174757265202773272076616c756545434453413a20696e76616c6964207369676e6174757265202776272076616c7565a264697066735822122003b4d7b34ba9618525aadaaa284c74578496e5bae3a8f0b479d84dd43e25d1f864736f6c63430007060033"
const SimpleSwapFactoryDeployedBinv0_6_7 = "0x608060405234801561001057600080fd5b506004361061004c5760003560e01c806315efd8a714610051578063a6021ace1461007a578063c70242ad14610082578063ee97f7f3146100a2575b600080fd5b61006461005f3660046102d6565b6100aa565b6040516100719190610308565b60405180910390f35b6100646101bb565b6100956100903660046102b5565b6101ca565b6040516100719190610359565b6100646101df565b60025460405160009182916100ee916001600160a01b0316906100d3903390879060200161031c565b604051602081830303815290604052805190602001206101ee565b6001546040516343431f6360e11b81529192506001600160a01b03808416926386863ec692610126928a929116908990600401610335565b600060405180830381600087803b15801561014057600080fd5b505af1158015610154573d6000803e3d6000fd5b5050506001600160a01b03821660009081526020819052604090819020805460ff19166001179055517fc0ffc525a1c7689549d7f79b49eca900e61ac49b43d977f680bcc3b36224c00491506101ab908390610308565b60405180910390a1949350505050565b6001546001600160a01b031681565b60006020819052908152604090205460ff1681565b6002546001600160a01b031681565b6000604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528360601b60148201526e5af43d82803e903d91602b57fd5bf360881b6028820152826037826000f59150506001600160a01b038116610293576040805162461bcd60e51b815260206004820152601760248201527f455243313136373a2063726561746532206661696c6564000000000000000000604482015290519081900360640190fd5b92915050565b80356001600160a01b03811681146102b057600080fd5b919050565b6000602082840312156102c6578081fd5b6102cf82610299565b9392505050565b6000806000606084860312156102ea578182fd5b6102f384610299565b95602085013595506040909401359392505050565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b90151581526020019056fea2646970667358221220879500e712d6cdee6e7b2cc582fef56f09aa010787a01406fe50796b05d33cba64736f6c63430007060033"
