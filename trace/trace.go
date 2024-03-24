// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package trace

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// TeatraceTeaInfo is an auto generated low-level Go binding around an user-defined struct.
type TeatraceTeaInfo struct {
	Planet     string
	Location   string
	Grower     string
	Species    string
	Quarantine string
	Packages   string
	Pick       string
	Picktime   string
	Picker     string
	Inserttime *big.Int
	Updatetime *big.Int
}

// TeatracelogisInfo is an auto generated low-level Go binding around an user-defined struct.
type TeatracelogisInfo struct {
	Path       string
	Logisway   string
	Driver     string
	Logistime  string
	Company    string
	Inserttime *big.Int
	Updatetime *big.Int
}

// TeatraceprocessInfo is an auto generated low-level Go binding around an user-defined struct.
type TeatraceprocessInfo struct {
	Process     string
	Processtime string
	Processer   string
	Company     string
	Inserttime  *big.Int
	Updatetime  *big.Int
}

// TeatraceproductionInfo is an auto generated low-level Go binding around an user-defined struct.
type TeatraceproductionInfo struct {
	Production     string
	Productiontime string
	Producter      string
	Company        string
	Inserttime     *big.Int
	Updatetime     *big.Int
}

// TeatracestorageInfo is an auto generated low-level Go binding around an user-defined struct.
type TeatracestorageInfo struct {
	Store      string
	Storetime  string
	Company    string
	Inserttime *big.Int
	Updatetime *big.Int
}

// TraceMetaData contains all meta data concerning the Trace contract.
var TraceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logisway\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logistime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"driver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"company\",\"type\":\"string\"}],\"name\":\"addOrupdateLogissData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"processa\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"processtime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"processer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"company\",\"type\":\"string\"}],\"name\":\"addOrupdateProcessData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"production\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"productiontime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"producter\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"company\",\"type\":\"string\"}],\"name\":\"addOrupdateProdData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"store\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"storetime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"company\",\"type\":\"string\"}],\"name\":\"addOrupdateStoreData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"planet\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"grower\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"quarantine\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"packages\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pick\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"picktime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"picker\",\"type\":\"string\"}],\"name\":\"addOrupdateTeaData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_producer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_storer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_logostic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_process\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tea\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tracedata\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"planet\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"grower\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"quarantine\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"packages\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pick\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"picktime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"picker\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"inserttime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatetime\",\"type\":\"uint256\"}],\"internalType\":\"structTeatrace.TeaInfo\",\"name\":\"tea\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"process\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"processtime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"processer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"company\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"inserttime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatetime\",\"type\":\"uint256\"}],\"internalType\":\"structTeatrace.processInfo\",\"name\":\"process\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"production\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"productiontime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"producter\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"company\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"inserttime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatetime\",\"type\":\"uint256\"}],\"internalType\":\"structTeatrace.productionInfo\",\"name\":\"production\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"store\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"storetime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"company\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"inserttime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatetime\",\"type\":\"uint256\"}],\"internalType\":\"structTeatrace.storageInfo\",\"name\":\"store\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logisway\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"driver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logistime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"company\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"inserttime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatetime\",\"type\":\"uint256\"}],\"internalType\":\"structTeatrace.logisInfo\",\"name\":\"logis\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200216d3803806200216d8339810160408190526200003491620000b4565b600180546001600160a01b03199081166001600160a01b0397881617909155600380548216958716959095179094556004805485169386169390931790925560028054841691851691909117905560068054909216921691909117905562000124565b80516001600160a01b0381168114620000af57600080fd5b919050565b600080600080600060a08688031215620000cd57600080fd5b620000d88662000097565b9450620000e86020870162000097565b9350620000f86040870162000097565b9250620001086060870162000097565b9150620001186080870162000097565b90509295509295909350565b61203980620001346000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80634f315563146100675780635c7e94951461007c5780638725f735146100a9578063a393d9d7146100bc578063cc10c8f8146100cf578063faeff189146100e2575b600080fd5b61007a610075366004611703565b6100f5565b005b61008f61008a3660046117b0565b610219565b6040516100a09594939291906119c6565b60405180910390f35b61007a6100b7366004611b4a565b61112a565b61007a6100ca366004611c3f565b611283565b61007a6100dd366004611d11565b6113bf565b61007a6100f0366004611c3f565b611581565b6003546001600160a01b031633146101495760405162461bcd60e51b81526020600482015260126024820152714f6e6c792073746f7265722063616e20646f60701b60448201526064015b60405180910390fd5b6000808560405161015a9190611e9c565b908152604051908190036020019020601a01546000036101775750425b6040518060a00160405280858152602001848152602001838152602001828152602001428152506000866040516101ae9190611e9c565b908152604051908190036020019020815160179091019081906101d19082611f43565b50602082015160018201906101e69082611f43565b50604082015160028201906101fb9082611f43565b50606082015160038201556080909101516004909101555050505050565b80516020818301810180516000825292820191909301209152604080516101608101909152815482908290829061024f90611eb8565b80601f016020809104026020016040519081016040528092919081815260200182805461027b90611eb8565b80156102c85780601f1061029d576101008083540402835291602001916102c8565b820191906000526020600020905b8154815290600101906020018083116102ab57829003601f168201915b505050505081526020016001820180546102e190611eb8565b80601f016020809104026020016040519081016040528092919081815260200182805461030d90611eb8565b801561035a5780601f1061032f5761010080835404028352916020019161035a565b820191906000526020600020905b81548152906001019060200180831161033d57829003601f168201915b5050505050815260200160028201805461037390611eb8565b80601f016020809104026020016040519081016040528092919081815260200182805461039f90611eb8565b80156103ec5780601f106103c1576101008083540402835291602001916103ec565b820191906000526020600020905b8154815290600101906020018083116103cf57829003601f168201915b5050505050815260200160038201805461040590611eb8565b80601f016020809104026020016040519081016040528092919081815260200182805461043190611eb8565b801561047e5780601f106104535761010080835404028352916020019161047e565b820191906000526020600020905b81548152906001019060200180831161046157829003601f168201915b5050505050815260200160048201805461049790611eb8565b80601f01602080910402602001604051908101604052809291908181526020018280546104c390611eb8565b80156105105780601f106104e557610100808354040283529160200191610510565b820191906000526020600020905b8154815290600101906020018083116104f357829003601f168201915b5050505050815260200160058201805461052990611eb8565b80601f016020809104026020016040519081016040528092919081815260200182805461055590611eb8565b80156105a25780601f10610577576101008083540402835291602001916105a2565b820191906000526020600020905b81548152906001019060200180831161058557829003601f168201915b505050505081526020016006820180546105bb90611eb8565b80601f01602080910402602001604051908101604052809291908181526020018280546105e790611eb8565b80156106345780601f1061060957610100808354040283529160200191610634565b820191906000526020600020905b81548152906001019060200180831161061757829003601f168201915b5050505050815260200160078201805461064d90611eb8565b80601f016020809104026020016040519081016040528092919081815260200182805461067990611eb8565b80156106c65780601f1061069b576101008083540402835291602001916106c6565b820191906000526020600020905b8154815290600101906020018083116106a957829003601f168201915b505050505081526020016008820180546106df90611eb8565b80601f016020809104026020016040519081016040528092919081815260200182805461070b90611eb8565b80156107585780601f1061072d57610100808354040283529160200191610758565b820191906000526020600020905b81548152906001019060200180831161073b57829003601f168201915b5050505050815260200160098201548152602001600a820154815250509080600b016040518060c001604052908160008201805461079590611eb8565b80601f01602080910402602001604051908101604052809291908181526020018280546107c190611eb8565b801561080e5780601f106107e35761010080835404028352916020019161080e565b820191906000526020600020905b8154815290600101906020018083116107f157829003601f168201915b5050505050815260200160018201805461082790611eb8565b80601f016020809104026020016040519081016040528092919081815260200182805461085390611eb8565b80156108a05780601f10610875576101008083540402835291602001916108a0565b820191906000526020600020905b81548152906001019060200180831161088357829003601f168201915b505050505081526020016002820180546108b990611eb8565b80601f01602080910402602001604051908101604052809291908181526020018280546108e590611eb8565b80156109325780601f1061090757610100808354040283529160200191610932565b820191906000526020600020905b81548152906001019060200180831161091557829003601f168201915b5050505050815260200160038201805461094b90611eb8565b80601f016020809104026020016040519081016040528092919081815260200182805461097790611eb8565b80156109c45780601f10610999576101008083540402835291602001916109c4565b820191906000526020600020905b8154815290600101906020018083116109a757829003601f168201915b505050505081526020016004820154815260200160058201548152505090806011016040518060c0016040529081600082018054610a0190611eb8565b80601f0160208091040260200160405190810160405280929190818152602001828054610a2d90611eb8565b8015610a7a5780601f10610a4f57610100808354040283529160200191610a7a565b820191906000526020600020905b815481529060010190602001808311610a5d57829003601f168201915b50505050508152602001600182018054610a9390611eb8565b80601f0160208091040260200160405190810160405280929190818152602001828054610abf90611eb8565b8015610b0c5780601f10610ae157610100808354040283529160200191610b0c565b820191906000526020600020905b815481529060010190602001808311610aef57829003601f168201915b50505050508152602001600282018054610b2590611eb8565b80601f0160208091040260200160405190810160405280929190818152602001828054610b5190611eb8565b8015610b9e5780601f10610b7357610100808354040283529160200191610b9e565b820191906000526020600020905b815481529060010190602001808311610b8157829003601f168201915b50505050508152602001600382018054610bb790611eb8565b80601f0160208091040260200160405190810160405280929190818152602001828054610be390611eb8565b8015610c305780601f10610c0557610100808354040283529160200191610c30565b820191906000526020600020905b815481529060010190602001808311610c1357829003601f168201915b505050505081526020016004820154815260200160058201548152505090806017016040518060a0016040529081600082018054610c6d90611eb8565b80601f0160208091040260200160405190810160405280929190818152602001828054610c9990611eb8565b8015610ce65780601f10610cbb57610100808354040283529160200191610ce6565b820191906000526020600020905b815481529060010190602001808311610cc957829003601f168201915b50505050508152602001600182018054610cff90611eb8565b80601f0160208091040260200160405190810160405280929190818152602001828054610d2b90611eb8565b8015610d785780601f10610d4d57610100808354040283529160200191610d78565b820191906000526020600020905b815481529060010190602001808311610d5b57829003601f168201915b50505050508152602001600282018054610d9190611eb8565b80601f0160208091040260200160405190810160405280929190818152602001828054610dbd90611eb8565b8015610e0a5780601f10610ddf57610100808354040283529160200191610e0a565b820191906000526020600020905b815481529060010190602001808311610ded57829003601f168201915b50505050508152602001600382015481526020016004820154815250509080601c016040518060e0016040529081600082018054610e4790611eb8565b80601f0160208091040260200160405190810160405280929190818152602001828054610e7390611eb8565b8015610ec05780601f10610e9557610100808354040283529160200191610ec0565b820191906000526020600020905b815481529060010190602001808311610ea357829003601f168201915b50505050508152602001600182018054610ed990611eb8565b80601f0160208091040260200160405190810160405280929190818152602001828054610f0590611eb8565b8015610f525780601f10610f2757610100808354040283529160200191610f52565b820191906000526020600020905b815481529060010190602001808311610f3557829003601f168201915b50505050508152602001600282018054610f6b90611eb8565b80601f0160208091040260200160405190810160405280929190818152602001828054610f9790611eb8565b8015610fe45780601f10610fb957610100808354040283529160200191610fe4565b820191906000526020600020905b815481529060010190602001808311610fc757829003601f168201915b50505050508152602001600382018054610ffd90611eb8565b80601f016020809104026020016040519081016040528092919081815260200182805461102990611eb8565b80156110765780601f1061104b57610100808354040283529160200191611076565b820191906000526020600020905b81548152906001019060200180831161105957829003601f168201915b5050505050815260200160048201805461108f90611eb8565b80601f01602080910402602001604051908101604052809291908181526020018280546110bb90611eb8565b80156111085780601f106110dd57610100808354040283529160200191611108565b820191906000526020600020905b8154815290600101906020018083116110eb57829003601f168201915b5050505050815260200160058201548152602001600682015481525050905085565b6004546001600160a01b0316331461117b5760405162461bcd60e51b81526020600482015260146024820152734f6e6c79206c6f676f737469632063616e20646f60601b6044820152606401610140565b6000808760405161118c9190611e9c565b908152604051908190036020019020602101546000036111a95750425b6040518060e00160405280878152602001868152602001848152602001858152602001838152602001828152602001428152506000886040516111ec9190611e9c565b9081526040519081900360200190208151601c90910190819061120f9082611f43565b50602082015160018201906112249082611f43565b50604082015160028201906112399082611f43565b506060820151600382019061124e9082611f43565b50608082015160048201906112639082611f43565b5060a0820151600582015560c09091015160069091015550505050505050565b6002546001600160a01b031633146112d35760405162461bcd60e51b81526020600482015260136024820152724f6e6c792070726f636573732063616e20646f60681b6044820152606401610140565b600080866040516112e49190611e9c565b908152604051908190036020019020600f01546000036113015750425b6040518060c001604052808681526020018581526020018481526020018381526020018281526020014281525060008760405161133e9190611e9c565b9081526040519081900360200190208151600b9091019081906113619082611f43565b50602082015160018201906113769082611f43565b506040820151600282019061138b9082611f43565b50606082015160038201906113a09082611f43565b506080820151600482015560a090910151600590910155505050505050565b6006546001600160a01b0316331461140b5760405162461bcd60e51b815260206004820152600f60248201526e4f6e6c79207465612063616e20646f60881b6044820152606401610140565b6000808b60405161141c9190611e9c565b908152604051908190036020019020600901546000036114395750425b6040518061016001604052808b81526020018a81526020018981526020018881526020018781526020018681526020018581526020018481526020018381526020018281526020014281525060008c6040516114959190611e9c565b908152604051908190036020019020815181906114b29082611f43565b50602082015160018201906114c79082611f43565b50604082015160028201906114dc9082611f43565b50606082015160038201906114f19082611f43565b50608082015160048201906115069082611f43565b5060a0820151600582019061151b9082611f43565b5060c082015160068201906115309082611f43565b5060e082015160078201906115459082611f43565b50610100820151600882019061155b9082611f43565b50610120820151600982015561014090910151600a909101555050505050505050505050565b6001546001600160a01b031633146115d25760405162461bcd60e51b81526020600482015260146024820152734f6e6c792070726f64756365722063616e20646f60601b6044820152606401610140565b600080866040516115e39190611e9c565b908152604051908190036020019020601501546000036116005750425b6040518060c001604052808681526020018581526020018481526020018381526020018281526020014281525060008760405161163d9190611e9c565b908152604051908190036020019020815160119091019081906113619082611f43565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261168757600080fd5b813567ffffffffffffffff808211156116a2576116a2611660565b604051601f8301601f19908116603f011681019082821181831017156116ca576116ca611660565b816040528381528660208588010111156116e357600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000806080858703121561171957600080fd5b843567ffffffffffffffff8082111561173157600080fd5b61173d88838901611676565b9550602087013591508082111561175357600080fd5b61175f88838901611676565b9450604087013591508082111561177557600080fd5b61178188838901611676565b9350606087013591508082111561179757600080fd5b506117a487828801611676565b91505092959194509250565b6000602082840312156117c257600080fd5b813567ffffffffffffffff8111156117d957600080fd5b6117e584828501611676565b949350505050565b60005b838110156118085781810151838201526020016117f0565b50506000910152565b600081518084526118298160208601602086016117ed565b601f01601f19169290920160200192915050565b6000815160c0845261185260c0850182611811565b90506020830151848203602086015261186b8282611811565b915050604083015184820360408601526118858282611811565b9150506060830151848203606086015261189f8282611811565b9150506080830151608085015260a083015160a08501528091505092915050565b6000815160a084526118d560a0850182611811565b9050602083015184820360208601526118ee8282611811565b915050604083015184820360408601526119088282611811565b91505060608301516060850152608083015160808501528091505092915050565b6000815160e0845261193e60e0850182611811565b9050602083015184820360208601526119578282611811565b915050604083015184820360408601526119718282611811565b9150506060830151848203606086015261198b8282611811565b915050608083015184820360808601526119a58282611811565b91505060a083015160a085015260c083015160c08501528091505092915050565b60a08152600086516101608060a08501526119e5610200850183611811565b91506020890151609f19808685030160c0870152611a038483611811565b935060408b01519150808685030160e0870152611a208483611811565b935060608b01519150610100818786030181880152611a3f8584611811565b945060808c01519250610120828887030181890152611a5e8685611811565b955060a08d015193506101408389880301818a0152611a7d8786611811565b965060c08e015194508389880301868a0152611a998786611811565b965060e08e0151955083898803016101808a0152611ab78787611811565b9650828e0151955083898803016101a08a0152611ad48787611811565b9650818e01516101c08a0152808e01516101e08a01525050505050508281036020840152611b02818861183d565b90508281036040840152611b16818761183d565b90508281036060840152611b2a81866118c0565b90508281036080840152611b3e8185611929565b98975050505050505050565b60008060008060008060c08789031215611b6357600080fd5b863567ffffffffffffffff80821115611b7b57600080fd5b611b878a838b01611676565b97506020890135915080821115611b9d57600080fd5b611ba98a838b01611676565b96506040890135915080821115611bbf57600080fd5b611bcb8a838b01611676565b95506060890135915080821115611be157600080fd5b611bed8a838b01611676565b94506080890135915080821115611c0357600080fd5b611c0f8a838b01611676565b935060a0890135915080821115611c2557600080fd5b50611c3289828a01611676565b9150509295509295509295565b600080600080600060a08688031215611c5757600080fd5b853567ffffffffffffffff80821115611c6f57600080fd5b611c7b89838a01611676565b96506020880135915080821115611c9157600080fd5b611c9d89838a01611676565b95506040880135915080821115611cb357600080fd5b611cbf89838a01611676565b94506060880135915080821115611cd557600080fd5b611ce189838a01611676565b93506080880135915080821115611cf757600080fd5b50611d0488828901611676565b9150509295509295909350565b6000806000806000806000806000806101408b8d031215611d3157600080fd5b8a3567ffffffffffffffff80821115611d4957600080fd5b611d558e838f01611676565b9b5060208d0135915080821115611d6b57600080fd5b611d778e838f01611676565b9a5060408d0135915080821115611d8d57600080fd5b611d998e838f01611676565b995060608d0135915080821115611daf57600080fd5b611dbb8e838f01611676565b985060808d0135915080821115611dd157600080fd5b611ddd8e838f01611676565b975060a08d0135915080821115611df357600080fd5b611dff8e838f01611676565b965060c08d0135915080821115611e1557600080fd5b611e218e838f01611676565b955060e08d0135915080821115611e3757600080fd5b611e438e838f01611676565b94506101008d0135915080821115611e5a57600080fd5b611e668e838f01611676565b93506101208d0135915080821115611e7d57600080fd5b50611e8a8d828e01611676565b9150509295989b9194979a5092959850565b60008251611eae8184602087016117ed565b9190910192915050565b600181811c90821680611ecc57607f821691505b602082108103611eec57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115611f3e576000816000526020600020601f850160051c81016020861015611f1b5750805b601f850160051c820191505b81811015611f3a57828155600101611f27565b5050505b505050565b815167ffffffffffffffff811115611f5d57611f5d611660565b611f7181611f6b8454611eb8565b84611ef2565b602080601f831160018114611fa65760008415611f8e5750858301515b600019600386901b1c1916600185901b178555611f3a565b600085815260208120601f198616915b82811015611fd557888601518255948401946001909101908401611fb6565b5085821015611ff35787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea26469706673582212204c8269e2f8ea44cdb983f347f96827c34901097a4155333c1b15c809da15662764736f6c63430008180033",
}

// TraceABI is the input ABI used to generate the binding from.
// Deprecated: Use TraceMetaData.ABI instead.
var TraceABI = TraceMetaData.ABI

// TraceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TraceMetaData.Bin instead.
var TraceBin = TraceMetaData.Bin

// DeployTrace deploys a new Ethereum contract, binding an instance of Trace to it.
func DeployTrace(auth *bind.TransactOpts, backend bind.ContractBackend, _producer common.Address, _storer common.Address, _logostic common.Address, _process common.Address, _tea common.Address) (common.Address, *types.Transaction, *Trace, error) {
	parsed, err := TraceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TraceBin), backend, _producer, _storer, _logostic, _process, _tea)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Trace{TraceCaller: TraceCaller{contract: contract}, TraceTransactor: TraceTransactor{contract: contract}, TraceFilterer: TraceFilterer{contract: contract}}, nil
}

// Trace is an auto generated Go binding around an Ethereum contract.
type Trace struct {
	TraceCaller     // Read-only binding to the contract
	TraceTransactor // Write-only binding to the contract
	TraceFilterer   // Log filterer for contract events
}

// TraceCaller is an auto generated read-only Go binding around an Ethereum contract.
type TraceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TraceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraceSession struct {
	Contract     *Trace            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraceCallerSession struct {
	Contract *TraceCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TraceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraceTransactorSession struct {
	Contract     *TraceTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraceRaw is an auto generated low-level Go binding around an Ethereum contract.
type TraceRaw struct {
	Contract *Trace // Generic contract binding to access the raw methods on
}

// TraceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraceCallerRaw struct {
	Contract *TraceCaller // Generic read-only contract binding to access the raw methods on
}

// TraceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraceTransactorRaw struct {
	Contract *TraceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrace creates a new instance of Trace, bound to a specific deployed contract.
func NewTrace(address common.Address, backend bind.ContractBackend) (*Trace, error) {
	contract, err := bindTrace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trace{TraceCaller: TraceCaller{contract: contract}, TraceTransactor: TraceTransactor{contract: contract}, TraceFilterer: TraceFilterer{contract: contract}}, nil
}

// NewTraceCaller creates a new read-only instance of Trace, bound to a specific deployed contract.
func NewTraceCaller(address common.Address, caller bind.ContractCaller) (*TraceCaller, error) {
	contract, err := bindTrace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraceCaller{contract: contract}, nil
}

// NewTraceTransactor creates a new write-only instance of Trace, bound to a specific deployed contract.
func NewTraceTransactor(address common.Address, transactor bind.ContractTransactor) (*TraceTransactor, error) {
	contract, err := bindTrace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraceTransactor{contract: contract}, nil
}

// NewTraceFilterer creates a new log filterer instance of Trace, bound to a specific deployed contract.
func NewTraceFilterer(address common.Address, filterer bind.ContractFilterer) (*TraceFilterer, error) {
	contract, err := bindTrace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraceFilterer{contract: contract}, nil
}

// bindTrace binds a generic wrapper to an already deployed contract.
func bindTrace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TraceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trace *TraceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trace.Contract.TraceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trace *TraceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trace.Contract.TraceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trace *TraceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trace.Contract.TraceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trace *TraceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trace *TraceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trace *TraceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trace.Contract.contract.Transact(opts, method, params...)
}

// Tracedata is a free data retrieval call binding the contract method 0x5c7e9495.
//
// Solidity: function tracedata(string ) view returns((string,string,string,string,string,string,string,string,string,uint256,uint256) tea, (string,string,string,string,uint256,uint256) process, (string,string,string,string,uint256,uint256) production, (string,string,string,uint256,uint256) store, (string,string,string,string,string,uint256,uint256) logis)
func (_Trace *TraceCaller) Tracedata(opts *bind.CallOpts, arg0 string) (struct {
	Tea        TeatraceTeaInfo
	Process    TeatraceprocessInfo
	Production TeatraceproductionInfo
	Store      TeatracestorageInfo
	Logis      TeatracelogisInfo
}, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "tracedata", arg0)

	outstruct := new(struct {
		Tea        TeatraceTeaInfo
		Process    TeatraceprocessInfo
		Production TeatraceproductionInfo
		Store      TeatracestorageInfo
		Logis      TeatracelogisInfo
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tea = *abi.ConvertType(out[0], new(TeatraceTeaInfo)).(*TeatraceTeaInfo)
	outstruct.Process = *abi.ConvertType(out[1], new(TeatraceprocessInfo)).(*TeatraceprocessInfo)
	outstruct.Production = *abi.ConvertType(out[2], new(TeatraceproductionInfo)).(*TeatraceproductionInfo)
	outstruct.Store = *abi.ConvertType(out[3], new(TeatracestorageInfo)).(*TeatracestorageInfo)
	outstruct.Logis = *abi.ConvertType(out[4], new(TeatracelogisInfo)).(*TeatracelogisInfo)

	return *outstruct, err

}

// Tracedata is a free data retrieval call binding the contract method 0x5c7e9495.
//
// Solidity: function tracedata(string ) view returns((string,string,string,string,string,string,string,string,string,uint256,uint256) tea, (string,string,string,string,uint256,uint256) process, (string,string,string,string,uint256,uint256) production, (string,string,string,uint256,uint256) store, (string,string,string,string,string,uint256,uint256) logis)
func (_Trace *TraceSession) Tracedata(arg0 string) (struct {
	Tea        TeatraceTeaInfo
	Process    TeatraceprocessInfo
	Production TeatraceproductionInfo
	Store      TeatracestorageInfo
	Logis      TeatracelogisInfo
}, error) {
	return _Trace.Contract.Tracedata(&_Trace.CallOpts, arg0)
}

// Tracedata is a free data retrieval call binding the contract method 0x5c7e9495.
//
// Solidity: function tracedata(string ) view returns((string,string,string,string,string,string,string,string,string,uint256,uint256) tea, (string,string,string,string,uint256,uint256) process, (string,string,string,string,uint256,uint256) production, (string,string,string,uint256,uint256) store, (string,string,string,string,string,uint256,uint256) logis)
func (_Trace *TraceCallerSession) Tracedata(arg0 string) (struct {
	Tea        TeatraceTeaInfo
	Process    TeatraceprocessInfo
	Production TeatraceproductionInfo
	Store      TeatracestorageInfo
	Logis      TeatracelogisInfo
}, error) {
	return _Trace.Contract.Tracedata(&_Trace.CallOpts, arg0)
}

// AddOrupdateLogissData is a paid mutator transaction binding the contract method 0x8725f735.
//
// Solidity: function addOrupdateLogissData(string id, string path, string logisway, string logistime, string driver, string company) returns()
func (_Trace *TraceTransactor) AddOrupdateLogissData(opts *bind.TransactOpts, id string, path string, logisway string, logistime string, driver string, company string) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "addOrupdateLogissData", id, path, logisway, logistime, driver, company)
}

// AddOrupdateLogissData is a paid mutator transaction binding the contract method 0x8725f735.
//
// Solidity: function addOrupdateLogissData(string id, string path, string logisway, string logistime, string driver, string company) returns()
func (_Trace *TraceSession) AddOrupdateLogissData(id string, path string, logisway string, logistime string, driver string, company string) (*types.Transaction, error) {
	return _Trace.Contract.AddOrupdateLogissData(&_Trace.TransactOpts, id, path, logisway, logistime, driver, company)
}

// AddOrupdateLogissData is a paid mutator transaction binding the contract method 0x8725f735.
//
// Solidity: function addOrupdateLogissData(string id, string path, string logisway, string logistime, string driver, string company) returns()
func (_Trace *TraceTransactorSession) AddOrupdateLogissData(id string, path string, logisway string, logistime string, driver string, company string) (*types.Transaction, error) {
	return _Trace.Contract.AddOrupdateLogissData(&_Trace.TransactOpts, id, path, logisway, logistime, driver, company)
}

// AddOrupdateProcessData is a paid mutator transaction binding the contract method 0xa393d9d7.
//
// Solidity: function addOrupdateProcessData(string id, string processa, string processtime, string processer, string company) returns()
func (_Trace *TraceTransactor) AddOrupdateProcessData(opts *bind.TransactOpts, id string, processa string, processtime string, processer string, company string) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "addOrupdateProcessData", id, processa, processtime, processer, company)
}

// AddOrupdateProcessData is a paid mutator transaction binding the contract method 0xa393d9d7.
//
// Solidity: function addOrupdateProcessData(string id, string processa, string processtime, string processer, string company) returns()
func (_Trace *TraceSession) AddOrupdateProcessData(id string, processa string, processtime string, processer string, company string) (*types.Transaction, error) {
	return _Trace.Contract.AddOrupdateProcessData(&_Trace.TransactOpts, id, processa, processtime, processer, company)
}

// AddOrupdateProcessData is a paid mutator transaction binding the contract method 0xa393d9d7.
//
// Solidity: function addOrupdateProcessData(string id, string processa, string processtime, string processer, string company) returns()
func (_Trace *TraceTransactorSession) AddOrupdateProcessData(id string, processa string, processtime string, processer string, company string) (*types.Transaction, error) {
	return _Trace.Contract.AddOrupdateProcessData(&_Trace.TransactOpts, id, processa, processtime, processer, company)
}

// AddOrupdateProdData is a paid mutator transaction binding the contract method 0xfaeff189.
//
// Solidity: function addOrupdateProdData(string id, string production, string productiontime, string producter, string company) returns()
func (_Trace *TraceTransactor) AddOrupdateProdData(opts *bind.TransactOpts, id string, production string, productiontime string, producter string, company string) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "addOrupdateProdData", id, production, productiontime, producter, company)
}

// AddOrupdateProdData is a paid mutator transaction binding the contract method 0xfaeff189.
//
// Solidity: function addOrupdateProdData(string id, string production, string productiontime, string producter, string company) returns()
func (_Trace *TraceSession) AddOrupdateProdData(id string, production string, productiontime string, producter string, company string) (*types.Transaction, error) {
	return _Trace.Contract.AddOrupdateProdData(&_Trace.TransactOpts, id, production, productiontime, producter, company)
}

// AddOrupdateProdData is a paid mutator transaction binding the contract method 0xfaeff189.
//
// Solidity: function addOrupdateProdData(string id, string production, string productiontime, string producter, string company) returns()
func (_Trace *TraceTransactorSession) AddOrupdateProdData(id string, production string, productiontime string, producter string, company string) (*types.Transaction, error) {
	return _Trace.Contract.AddOrupdateProdData(&_Trace.TransactOpts, id, production, productiontime, producter, company)
}

// AddOrupdateStoreData is a paid mutator transaction binding the contract method 0x4f315563.
//
// Solidity: function addOrupdateStoreData(string id, string store, string storetime, string company) returns()
func (_Trace *TraceTransactor) AddOrupdateStoreData(opts *bind.TransactOpts, id string, store string, storetime string, company string) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "addOrupdateStoreData", id, store, storetime, company)
}

// AddOrupdateStoreData is a paid mutator transaction binding the contract method 0x4f315563.
//
// Solidity: function addOrupdateStoreData(string id, string store, string storetime, string company) returns()
func (_Trace *TraceSession) AddOrupdateStoreData(id string, store string, storetime string, company string) (*types.Transaction, error) {
	return _Trace.Contract.AddOrupdateStoreData(&_Trace.TransactOpts, id, store, storetime, company)
}

// AddOrupdateStoreData is a paid mutator transaction binding the contract method 0x4f315563.
//
// Solidity: function addOrupdateStoreData(string id, string store, string storetime, string company) returns()
func (_Trace *TraceTransactorSession) AddOrupdateStoreData(id string, store string, storetime string, company string) (*types.Transaction, error) {
	return _Trace.Contract.AddOrupdateStoreData(&_Trace.TransactOpts, id, store, storetime, company)
}

// AddOrupdateTeaData is a paid mutator transaction binding the contract method 0xcc10c8f8.
//
// Solidity: function addOrupdateTeaData(string id, string planet, string location, string grower, string species, string quarantine, string packages, string pick, string picktime, string picker) returns()
func (_Trace *TraceTransactor) AddOrupdateTeaData(opts *bind.TransactOpts, id string, planet string, location string, grower string, species string, quarantine string, packages string, pick string, picktime string, picker string) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "addOrupdateTeaData", id, planet, location, grower, species, quarantine, packages, pick, picktime, picker)
}

// AddOrupdateTeaData is a paid mutator transaction binding the contract method 0xcc10c8f8.
//
// Solidity: function addOrupdateTeaData(string id, string planet, string location, string grower, string species, string quarantine, string packages, string pick, string picktime, string picker) returns()
func (_Trace *TraceSession) AddOrupdateTeaData(id string, planet string, location string, grower string, species string, quarantine string, packages string, pick string, picktime string, picker string) (*types.Transaction, error) {
	return _Trace.Contract.AddOrupdateTeaData(&_Trace.TransactOpts, id, planet, location, grower, species, quarantine, packages, pick, picktime, picker)
}

// AddOrupdateTeaData is a paid mutator transaction binding the contract method 0xcc10c8f8.
//
// Solidity: function addOrupdateTeaData(string id, string planet, string location, string grower, string species, string quarantine, string packages, string pick, string picktime, string picker) returns()
func (_Trace *TraceTransactorSession) AddOrupdateTeaData(id string, planet string, location string, grower string, species string, quarantine string, packages string, pick string, picktime string, picker string) (*types.Transaction, error) {
	return _Trace.Contract.AddOrupdateTeaData(&_Trace.TransactOpts, id, planet, location, grower, species, quarantine, packages, pick, picktime, picker)
}
