// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pricefeeds

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

// PriceFeedsTickerInfo is an auto generated low-level Go binding around an user-defined struct.
type PriceFeedsTickerInfo struct {
	Price         *big.Int
	Decimals      *big.Int
	Timestamp     *big.Int
	Source        string
	Sender        common.Address
	Currency      string
	Ticker        string
	TickerAddress common.Address
}

// PricefeedsMetaData contains all meta data concerning the Pricefeeds contract.
var PricefeedsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updateNums\",\"type\":\"uint256\"}],\"name\":\"PriceUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tickerAddresses\",\"type\":\"address[]\"}],\"name\":\"getTickerInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"source\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"currency\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tickerAddress\",\"type\":\"address\"}],\"internalType\":\"structPriceFeeds.TickerInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tickerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getTokenValueOfAXC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxArrayLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTickersLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"prices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"source\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"currency\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tickerAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxArrayLength\",\"type\":\"uint256\"}],\"name\":\"setMaxArrayLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxTickersLength\",\"type\":\"uint256\"}],\"name\":\"setMaxTickersLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"source\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"currency\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tickerAddress\",\"type\":\"address\"}],\"internalType\":\"structPriceFeeds.TickerInfo[]\",\"name\":\"_TickerInfos\",\"type\":\"tuple[]\"}],\"name\":\"updatePrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526103e86003600050909055601460046000509090553480156100265760006000fd5b505b33600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600260005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b6100ca565b611fd2806100d96000396000f3fe60806040523480156100115760006000fd5b50600436106100d45760003560e01c80639b19251a1161008d578063a6f9dae111610067578063a6f9dae114610218578063cb353e4214610234578063cfed246b14610250578063e43252d714610287576100d4565b80639b19251a1461019c5780639dc1b0cd146101cc578063a2b4bfbf146101e8576100d4565b80632a9ff49c146100da57806347b8183e1461010a57806353bda15b146101265780638ab1d681146101445780638da5cb5b146101605780639a3128f91461017e576100d4565b60006000fd5b6100f460048036038101906100ef9190611617565b6102a3565b6040516101019190611b47565b60405180910390f35b610124600480360381019061011f91906116de565b61035c565b005b61012e6103ff565b60405161013b9190611b47565b60405180910390f35b61015e600480360381019061015991906115ec565b610408565b005b6101686104fa565b6040516101759190611a68565b60405180910390f35b610186610520565b6040516101939190611b47565b60405180910390f35b6101b660048036038101906101b191906115ec565b610529565b6040516101c39190611aa7565b60405180910390f35b6101e660048036038101906101e191906116de565b61054e565b005b61020260048036038101906101fd9190611656565b6105f1565b60405161020f9190611a84565b60405180910390f35b610232600480360381019061022d91906115ec565b610a65565b005b61024e6004803603810190610249919061169a565b610b3d565b005b61026a600480360381019061026591906115ec565b610ec8565b60405161027e989796959493929190611b8d565b60405180910390f35b6102a1600480360381019061029c91906115ec565b6110fd565b005b60006000600060005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005090506000816000016000505411151561033a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033190611ae4565b60405180910390fd5b80600001600050548361034d9190611d4b565b91505061035656505b92915050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156103ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e590611ac3565b60405180910390fd5b8060046000508190909055505b5b50565b60036000505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561049a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049190611ac3565b60405180910390fd5b6000600260005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b5b50565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60046000505481565b600260005060205280600052604060002060009150909054906101000a900460ff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156105e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105d790611ac3565b60405180910390fd5b8060036000508190909055505b5b50565b606060046000505482511115151561063e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063590611b05565b60405180910390fd5b6000825167ffffffffffffffff811115610681577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156106ba57816020015b6106a76111ef565b81526020019060019003908161069f5790505b5090506000600090505b8351811015610a55576000600060005060008684815181101515610711577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000509050806040518061010001604052908160008201600050548152602001600182016000505481526020016002820160005054815260200160038201600050805461079f90611e10565b80601f01602080910402602001604051908101604052809291908181526020018280546107cb90611e10565b80156108185780601f106107ed57610100808354040283529160200191610818565b820191906000526020600020905b8154815290600101906020018083116107fb57829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160058201600050805461088a90611e10565b80601f01602080910402602001604051908101604052809291908181526020018280546108b690611e10565b80156109035780601f106108d857610100808354040283529160200191610903565b820191906000526020600020905b8154815290600101906020018083116108e657829003601f168201915b5050505050815260200160068201600050805461091f90611e10565b80601f016020809104026020016040519081016040528092919081815260200182805461094b90611e10565b80156109985780601f1061096d57610100808354040283529160200191610998565b820191906000526020600020905b81548152906001019060200180831161097b57829003601f168201915b505050505081526020016007820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200150508383815181101515610a35577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010181905250505b8080610a4d90611e45565b9150506106c4565b5080915050610a6056505b919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610af7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aee90611ac3565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5b50565b600260005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610bce576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bc590611b26565b60405180910390fd5b600360005054815111151515610c19576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1090611b05565b60405180910390fd5b6000600090505b8151811015610e715760008282815181101515610c66577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101519050604051806101000160405280826000015181526020018260200151815260200182604001518152602001826060015181526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018260a0015181526020018260c0015181526020018260e0015173ffffffffffffffffffffffffffffffffffffffff16815260200150600060005060008360e0015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000506000820151816000016000509090556020820151816001016000509090556040820151816002016000509090556060820151816003016000509080519060200190610d8a929190611263565b5060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a0820151816005016000509080519060200190610df1929190611263565b5060c0820151816006016000509080519060200190610e11929190611263565b5060e08201518160070160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550905050505b8080610e6990611e45565b915050610c20565b503373ffffffffffffffffffffffffffffffffffffffff167fb556fac599c3c70efb9ab1fa725ecace6c81cc48d1455f886607def065f3e0c0428351604051610ebb929190611b63565b60405180910390a25b5b50565b6000600050602052806000526040600020600091509050806000016000505490806001016000505490806002016000505490806003016000508054610f0c90611e10565b80601f0160208091040260200160405190810160405280929190818152602001828054610f3890611e10565b8015610f855780601f10610f5a57610100808354040283529160200191610f85565b820191906000526020600020905b815481529060010190602001808311610f6857829003601f168201915b5050505050908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806005016000508054610fc390611e10565b80601f0160208091040260200160405190810160405280929190818152602001828054610fef90611e10565b801561103c5780601f106110115761010080835404028352916020019161103c565b820191906000526020600020905b81548152906001019060200180831161101f57829003601f168201915b50505050509080600601600050805461105490611e10565b80601f016020809104026020016040519081016040528092919081815260200182805461108090611e10565b80156110cd5780601f106110a2576101008083540402835291602001916110cd565b820191906000526020600020905b8154815290600101906020018083116110b057829003601f168201915b5050505050908060070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905088565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561118f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161118690611ac3565b60405180910390fd5b6001600260005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b5b50565b60405180610100016040528060008152602001600081526020016000815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020015090565b82805461126f90611e10565b90600052602060002090601f01602090048101928261129157600085556112dd565b82601f106112aa57805160ff19168380011785556112dd565b828001600101855582156112dd579182015b828111156112dc57825182600050909055916020019190600101906112bc565b5b5090506112ea91906112ee565b5090565b6112f3565b8082111561130d57600081815060009055506001016112f3565b509056611f9b565b600061132861132384611c53565b611c21565b905080838252602082019050828560208602820111156113485760006000fd5b60005b85811015611379578161135e8882611423565b8452602084019350602083019250505b60018101905061134b565b5050505b9392505050565b600061139761139284611c80565b611c21565b9050808382526020820190508260005b858110156113d857813585016113bd88826114c3565b8452602084019350602083019250505b6001810190506113a7565b5050505b9392505050565b60006113f66113f184611cad565b611c21565b90508281526020810184848401111561140f5760006000fd5b61141a848285611dcb565b505b9392505050565b60008135905061143281611f65565b5b92915050565b600082601f830112151561144d5760006000fd5b813561145d848260208601611315565b9150505b92915050565b600082601f830112151561147b5760006000fd5b813561148b848260208601611384565b9150505b92915050565b600082601f83011215156114a95760006000fd5b81356114b98482602086016113e3565b9150505b92915050565b600061010082840312156114d75760006000fd5b6114e2610100611c21565b905060006114f2848285016115d6565b6000830152506020611506848285016115d6565b602083015250604061151a848285016115d6565b604083015250606082013567ffffffffffffffff81111561153b5760006000fd5b61154784828501611495565b606083015250608061155b84828501611423565b60808301525060a082013567ffffffffffffffff81111561157c5760006000fd5b61158884828501611495565b60a08301525060c082013567ffffffffffffffff8111156115a95760006000fd5b6115b584828501611495565b60c08301525060e06115c984828501611423565b60e0830152505b92915050565b6000813590506115e581611f80565b5b92915050565b6000602082840312156115ff5760006000fd5b600061160d84828501611423565b9150505b92915050565b600060006040838503121561162c5760006000fd5b600061163a85828601611423565b925050602061164b858286016115d6565b9150505b9250929050565b6000602082840312156116695760006000fd5b600082013567ffffffffffffffff8111156116845760006000fd5b61169084828501611439565b9150505b92915050565b6000602082840312156116ad5760006000fd5b600082013567ffffffffffffffff8111156116c85760006000fd5b6116d484828501611467565b9150505b92915050565b6000602082840312156116f15760006000fd5b60006116ff848285016115d6565b9150505b92915050565b60006117158383611989565b90505b92915050565b61172781611d7f565b82525b5050565b61173781611d7f565b82525b5050565b600061174982611cef565b6117538185611d15565b93508360208202850161176585611cde565b8060005b858110156117a257848403895281516117828582611709565b945061178d83611d07565b925060208a019950505b600181019050611769565b5082975087955050505050505b92915050565b6117be81611d92565b82525b5050565b60006117d082611cfb565b6117da8185611d27565b93506117ea818560208601611ddb565b6117f381611f53565b84019150505b92915050565b600061180a82611cfb565b6118148185611d39565b9350611824818560208601611ddb565b61182d81611f53565b84019150505b92915050565b6000611846602683611d39565b91507f4f6e6c7920746865206f776e65722063616e2063616c6c20746869732066756e60008301527f6374696f6e2e000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b60006118ad601f83611d39565b91507f5072696365206d7573742062652067726561746572207468616e207a65726f0060008301526020820190505b919050565b60006118ee602283611d39565b91507f496e7075742061727261792065786365656473206d6178696d756d206c656e6760008301527f746800000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000611955601f83611d39565b91507f43616c6c6572206973206e6f7420696e207468652077686974656c6973742e0060008301526020820190505b919050565b6000610100830160008301516119a26000860182611a48565b5060208301516119b56020860182611a48565b5060408301516119c86040860182611a48565b50606083015184820360608601526119e082826117c5565b91505060808301516119f5608086018261171e565b5060a083015184820360a0860152611a0d82826117c5565b91505060c083015184820360c0860152611a2782826117c5565b91505060e0830151611a3c60e086018261171e565b50809150505b92915050565b611a5181611dc0565b82525b5050565b611a6181611dc0565b82525b5050565b6000602082019050611a7d600083018461172e565b5b92915050565b60006020820190508181036000830152611a9e818461173e565b90505b92915050565b6000602082019050611abc60008301846117b5565b5b92915050565b60006020820190508181036000830152611adc81611839565b90505b919050565b60006020820190508181036000830152611afd816118a0565b90505b919050565b60006020820190508181036000830152611b1e816118e1565b90505b919050565b60006020820190508181036000830152611b3f81611948565b90505b919050565b6000602082019050611b5c6000830184611a58565b5b92915050565b6000604082019050611b786000830185611a58565b611b856020830184611a58565b5b9392505050565b600061010082019050611ba3600083018b611a58565b611bb0602083018a611a58565b611bbd6040830189611a58565b8181036060830152611bcf81886117ff565b9050611bde608083018761172e565b81810360a0830152611bf081866117ff565b905081810360c0830152611c0481856117ff565b9050611c1360e083018461172e565b5b9998505050505050505050565b6000604051905081810181811067ffffffffffffffff82111715611c4857611c47611f22565b5b80604052505b919050565b600067ffffffffffffffff821115611c6e57611c6d611f22565b5b6020820290506020810190505b919050565b600067ffffffffffffffff821115611c9b57611c9a611f22565b5b6020820290506020810190505b919050565b600067ffffffffffffffff821115611cc857611cc7611f22565b5b601f19601f83011690506020810190505b919050565b60008190506020820190505b919050565b6000815190505b919050565b6000815190505b919050565b60006020820190505b919050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b6000611d5682611dc0565b9150611d6183611dc0565b9250821515611d7357611d72611ec0565b5b82820490505b92915050565b6000611d8a82611d9f565b90505b919050565b600081151590505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b828183376000838301525b505050565b60005b83811015611dfa5780820151818401525b602081019050611dde565b83811115611e09576000848401525b505b505050565b600060028204905060018216801515611e2a57607f821691505b60208210811415611e3e57611e3d611ef1565b5b505b919050565b6000611e5082611dc0565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611e8357611e82611e8f565b5b6001820190505b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b565b6000601f19601f83011690505b919050565b611f6e81611d7f565b81141515611f7c5760006000fd5b5b50565b611f8981611dc0565b81141515611f975760006000fd5b5b50565bfea2646970667358221220f5ec5d411a88862fd16c1da8711892727d2fda6655b48c4d7229aa5bffd146a564736f6c63430008000033",
}

// PricefeedsABI is the input ABI used to generate the binding from.
// Deprecated: Use PricefeedsMetaData.ABI instead.
var PricefeedsABI = PricefeedsMetaData.ABI

// PricefeedsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PricefeedsMetaData.Bin instead.
var PricefeedsBin = PricefeedsMetaData.Bin

// DeployPricefeeds deploys a new Ethereum contract, binding an instance of Pricefeeds to it.
func DeployPricefeeds(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pricefeeds, error) {
	parsed, err := PricefeedsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PricefeedsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pricefeeds{PricefeedsCaller: PricefeedsCaller{contract: contract}, PricefeedsTransactor: PricefeedsTransactor{contract: contract}, PricefeedsFilterer: PricefeedsFilterer{contract: contract}}, nil
}

// Pricefeeds is an auto generated Go binding around an Ethereum contract.
type Pricefeeds struct {
	PricefeedsCaller     // Read-only binding to the contract
	PricefeedsTransactor // Write-only binding to the contract
	PricefeedsFilterer   // Log filterer for contract events
}

// PricefeedsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PricefeedsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PricefeedsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PricefeedsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PricefeedsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PricefeedsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PricefeedsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PricefeedsSession struct {
	Contract     *Pricefeeds       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PricefeedsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PricefeedsCallerSession struct {
	Contract *PricefeedsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PricefeedsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PricefeedsTransactorSession struct {
	Contract     *PricefeedsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PricefeedsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PricefeedsRaw struct {
	Contract *Pricefeeds // Generic contract binding to access the raw methods on
}

// PricefeedsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PricefeedsCallerRaw struct {
	Contract *PricefeedsCaller // Generic read-only contract binding to access the raw methods on
}

// PricefeedsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PricefeedsTransactorRaw struct {
	Contract *PricefeedsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPricefeeds creates a new instance of Pricefeeds, bound to a specific deployed contract.
func NewPricefeeds(address common.Address, backend bind.ContractBackend) (*Pricefeeds, error) {
	contract, err := bindPricefeeds(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pricefeeds{PricefeedsCaller: PricefeedsCaller{contract: contract}, PricefeedsTransactor: PricefeedsTransactor{contract: contract}, PricefeedsFilterer: PricefeedsFilterer{contract: contract}}, nil
}

// NewPricefeedsCaller creates a new read-only instance of Pricefeeds, bound to a specific deployed contract.
func NewPricefeedsCaller(address common.Address, caller bind.ContractCaller) (*PricefeedsCaller, error) {
	contract, err := bindPricefeeds(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PricefeedsCaller{contract: contract}, nil
}

// NewPricefeedsTransactor creates a new write-only instance of Pricefeeds, bound to a specific deployed contract.
func NewPricefeedsTransactor(address common.Address, transactor bind.ContractTransactor) (*PricefeedsTransactor, error) {
	contract, err := bindPricefeeds(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PricefeedsTransactor{contract: contract}, nil
}

// NewPricefeedsFilterer creates a new log filterer instance of Pricefeeds, bound to a specific deployed contract.
func NewPricefeedsFilterer(address common.Address, filterer bind.ContractFilterer) (*PricefeedsFilterer, error) {
	contract, err := bindPricefeeds(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PricefeedsFilterer{contract: contract}, nil
}

// bindPricefeeds binds a generic wrapper to an already deployed contract.
func bindPricefeeds(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PricefeedsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pricefeeds *PricefeedsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pricefeeds.Contract.PricefeedsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pricefeeds *PricefeedsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pricefeeds.Contract.PricefeedsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pricefeeds *PricefeedsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pricefeeds.Contract.PricefeedsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pricefeeds *PricefeedsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pricefeeds.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pricefeeds *PricefeedsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pricefeeds.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pricefeeds *PricefeedsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pricefeeds.Contract.contract.Transact(opts, method, params...)
}

// GetTickerInfos is a free data retrieval call binding the contract method 0xa2b4bfbf.
//
// Solidity: function getTickerInfos(address[] _tickerAddresses) view returns((uint256,uint256,uint256,string,address,string,string,address)[])
func (_Pricefeeds *PricefeedsCaller) GetTickerInfos(opts *bind.CallOpts, _tickerAddresses []common.Address) ([]PriceFeedsTickerInfo, error) {
	var out []interface{}
	err := _Pricefeeds.contract.Call(opts, &out, "getTickerInfos", _tickerAddresses)

	if err != nil {
		return *new([]PriceFeedsTickerInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]PriceFeedsTickerInfo)).(*[]PriceFeedsTickerInfo)

	return out0, err

}

// GetTickerInfos is a free data retrieval call binding the contract method 0xa2b4bfbf.
//
// Solidity: function getTickerInfos(address[] _tickerAddresses) view returns((uint256,uint256,uint256,string,address,string,string,address)[])
func (_Pricefeeds *PricefeedsSession) GetTickerInfos(_tickerAddresses []common.Address) ([]PriceFeedsTickerInfo, error) {
	return _Pricefeeds.Contract.GetTickerInfos(&_Pricefeeds.CallOpts, _tickerAddresses)
}

// GetTickerInfos is a free data retrieval call binding the contract method 0xa2b4bfbf.
//
// Solidity: function getTickerInfos(address[] _tickerAddresses) view returns((uint256,uint256,uint256,string,address,string,string,address)[])
func (_Pricefeeds *PricefeedsCallerSession) GetTickerInfos(_tickerAddresses []common.Address) ([]PriceFeedsTickerInfo, error) {
	return _Pricefeeds.Contract.GetTickerInfos(&_Pricefeeds.CallOpts, _tickerAddresses)
}

// GetTokenValueOfAXC is a free data retrieval call binding the contract method 0x2a9ff49c.
//
// Solidity: function getTokenValueOfAXC(address _tickerAddress, uint256 _amount) view returns(uint256)
func (_Pricefeeds *PricefeedsCaller) GetTokenValueOfAXC(opts *bind.CallOpts, _tickerAddress common.Address, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pricefeeds.contract.Call(opts, &out, "getTokenValueOfAXC", _tickerAddress, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenValueOfAXC is a free data retrieval call binding the contract method 0x2a9ff49c.
//
// Solidity: function getTokenValueOfAXC(address _tickerAddress, uint256 _amount) view returns(uint256)
func (_Pricefeeds *PricefeedsSession) GetTokenValueOfAXC(_tickerAddress common.Address, _amount *big.Int) (*big.Int, error) {
	return _Pricefeeds.Contract.GetTokenValueOfAXC(&_Pricefeeds.CallOpts, _tickerAddress, _amount)
}

// GetTokenValueOfAXC is a free data retrieval call binding the contract method 0x2a9ff49c.
//
// Solidity: function getTokenValueOfAXC(address _tickerAddress, uint256 _amount) view returns(uint256)
func (_Pricefeeds *PricefeedsCallerSession) GetTokenValueOfAXC(_tickerAddress common.Address, _amount *big.Int) (*big.Int, error) {
	return _Pricefeeds.Contract.GetTokenValueOfAXC(&_Pricefeeds.CallOpts, _tickerAddress, _amount)
}

// MaxArrayLength is a free data retrieval call binding the contract method 0x53bda15b.
//
// Solidity: function maxArrayLength() view returns(uint256)
func (_Pricefeeds *PricefeedsCaller) MaxArrayLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pricefeeds.contract.Call(opts, &out, "maxArrayLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxArrayLength is a free data retrieval call binding the contract method 0x53bda15b.
//
// Solidity: function maxArrayLength() view returns(uint256)
func (_Pricefeeds *PricefeedsSession) MaxArrayLength() (*big.Int, error) {
	return _Pricefeeds.Contract.MaxArrayLength(&_Pricefeeds.CallOpts)
}

// MaxArrayLength is a free data retrieval call binding the contract method 0x53bda15b.
//
// Solidity: function maxArrayLength() view returns(uint256)
func (_Pricefeeds *PricefeedsCallerSession) MaxArrayLength() (*big.Int, error) {
	return _Pricefeeds.Contract.MaxArrayLength(&_Pricefeeds.CallOpts)
}

// MaxTickersLength is a free data retrieval call binding the contract method 0x9a3128f9.
//
// Solidity: function maxTickersLength() view returns(uint256)
func (_Pricefeeds *PricefeedsCaller) MaxTickersLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pricefeeds.contract.Call(opts, &out, "maxTickersLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTickersLength is a free data retrieval call binding the contract method 0x9a3128f9.
//
// Solidity: function maxTickersLength() view returns(uint256)
func (_Pricefeeds *PricefeedsSession) MaxTickersLength() (*big.Int, error) {
	return _Pricefeeds.Contract.MaxTickersLength(&_Pricefeeds.CallOpts)
}

// MaxTickersLength is a free data retrieval call binding the contract method 0x9a3128f9.
//
// Solidity: function maxTickersLength() view returns(uint256)
func (_Pricefeeds *PricefeedsCallerSession) MaxTickersLength() (*big.Int, error) {
	return _Pricefeeds.Contract.MaxTickersLength(&_Pricefeeds.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pricefeeds *PricefeedsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pricefeeds.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pricefeeds *PricefeedsSession) Owner() (common.Address, error) {
	return _Pricefeeds.Contract.Owner(&_Pricefeeds.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pricefeeds *PricefeedsCallerSession) Owner() (common.Address, error) {
	return _Pricefeeds.Contract.Owner(&_Pricefeeds.CallOpts)
}

// Prices is a free data retrieval call binding the contract method 0xcfed246b.
//
// Solidity: function prices(address ) view returns(uint256 price, uint256 decimals, uint256 timestamp, string source, address sender, string currency, string ticker, address tickerAddress)
func (_Pricefeeds *PricefeedsCaller) Prices(opts *bind.CallOpts, arg0 common.Address) (struct {
	Price         *big.Int
	Decimals      *big.Int
	Timestamp     *big.Int
	Source        string
	Sender        common.Address
	Currency      string
	Ticker        string
	TickerAddress common.Address
}, error) {
	var out []interface{}
	err := _Pricefeeds.contract.Call(opts, &out, "prices", arg0)

	outstruct := new(struct {
		Price         *big.Int
		Decimals      *big.Int
		Timestamp     *big.Int
		Source        string
		Sender        common.Address
		Currency      string
		Ticker        string
		TickerAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Decimals = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Source = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Sender = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Currency = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Ticker = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.TickerAddress = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Prices is a free data retrieval call binding the contract method 0xcfed246b.
//
// Solidity: function prices(address ) view returns(uint256 price, uint256 decimals, uint256 timestamp, string source, address sender, string currency, string ticker, address tickerAddress)
func (_Pricefeeds *PricefeedsSession) Prices(arg0 common.Address) (struct {
	Price         *big.Int
	Decimals      *big.Int
	Timestamp     *big.Int
	Source        string
	Sender        common.Address
	Currency      string
	Ticker        string
	TickerAddress common.Address
}, error) {
	return _Pricefeeds.Contract.Prices(&_Pricefeeds.CallOpts, arg0)
}

// Prices is a free data retrieval call binding the contract method 0xcfed246b.
//
// Solidity: function prices(address ) view returns(uint256 price, uint256 decimals, uint256 timestamp, string source, address sender, string currency, string ticker, address tickerAddress)
func (_Pricefeeds *PricefeedsCallerSession) Prices(arg0 common.Address) (struct {
	Price         *big.Int
	Decimals      *big.Int
	Timestamp     *big.Int
	Source        string
	Sender        common.Address
	Currency      string
	Ticker        string
	TickerAddress common.Address
}, error) {
	return _Pricefeeds.Contract.Prices(&_Pricefeeds.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Pricefeeds *PricefeedsCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Pricefeeds.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Pricefeeds *PricefeedsSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Pricefeeds.Contract.Whitelist(&_Pricefeeds.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Pricefeeds *PricefeedsCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Pricefeeds.Contract.Whitelist(&_Pricefeeds.CallOpts, arg0)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address _address) returns()
func (_Pricefeeds *PricefeedsTransactor) AddToWhitelist(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Pricefeeds.contract.Transact(opts, "addToWhitelist", _address)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address _address) returns()
func (_Pricefeeds *PricefeedsSession) AddToWhitelist(_address common.Address) (*types.Transaction, error) {
	return _Pricefeeds.Contract.AddToWhitelist(&_Pricefeeds.TransactOpts, _address)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address _address) returns()
func (_Pricefeeds *PricefeedsTransactorSession) AddToWhitelist(_address common.Address) (*types.Transaction, error) {
	return _Pricefeeds.Contract.AddToWhitelist(&_Pricefeeds.TransactOpts, _address)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_Pricefeeds *PricefeedsTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Pricefeeds.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_Pricefeeds *PricefeedsSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Pricefeeds.Contract.ChangeOwner(&_Pricefeeds.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_Pricefeeds *PricefeedsTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Pricefeeds.Contract.ChangeOwner(&_Pricefeeds.TransactOpts, _newOwner)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address _address) returns()
func (_Pricefeeds *PricefeedsTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Pricefeeds.contract.Transact(opts, "removeFromWhitelist", _address)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address _address) returns()
func (_Pricefeeds *PricefeedsSession) RemoveFromWhitelist(_address common.Address) (*types.Transaction, error) {
	return _Pricefeeds.Contract.RemoveFromWhitelist(&_Pricefeeds.TransactOpts, _address)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address _address) returns()
func (_Pricefeeds *PricefeedsTransactorSession) RemoveFromWhitelist(_address common.Address) (*types.Transaction, error) {
	return _Pricefeeds.Contract.RemoveFromWhitelist(&_Pricefeeds.TransactOpts, _address)
}

// SetMaxArrayLength is a paid mutator transaction binding the contract method 0x9dc1b0cd.
//
// Solidity: function setMaxArrayLength(uint256 _maxArrayLength) returns()
func (_Pricefeeds *PricefeedsTransactor) SetMaxArrayLength(opts *bind.TransactOpts, _maxArrayLength *big.Int) (*types.Transaction, error) {
	return _Pricefeeds.contract.Transact(opts, "setMaxArrayLength", _maxArrayLength)
}

// SetMaxArrayLength is a paid mutator transaction binding the contract method 0x9dc1b0cd.
//
// Solidity: function setMaxArrayLength(uint256 _maxArrayLength) returns()
func (_Pricefeeds *PricefeedsSession) SetMaxArrayLength(_maxArrayLength *big.Int) (*types.Transaction, error) {
	return _Pricefeeds.Contract.SetMaxArrayLength(&_Pricefeeds.TransactOpts, _maxArrayLength)
}

// SetMaxArrayLength is a paid mutator transaction binding the contract method 0x9dc1b0cd.
//
// Solidity: function setMaxArrayLength(uint256 _maxArrayLength) returns()
func (_Pricefeeds *PricefeedsTransactorSession) SetMaxArrayLength(_maxArrayLength *big.Int) (*types.Transaction, error) {
	return _Pricefeeds.Contract.SetMaxArrayLength(&_Pricefeeds.TransactOpts, _maxArrayLength)
}

// SetMaxTickersLength is a paid mutator transaction binding the contract method 0x47b8183e.
//
// Solidity: function setMaxTickersLength(uint256 _maxTickersLength) returns()
func (_Pricefeeds *PricefeedsTransactor) SetMaxTickersLength(opts *bind.TransactOpts, _maxTickersLength *big.Int) (*types.Transaction, error) {
	return _Pricefeeds.contract.Transact(opts, "setMaxTickersLength", _maxTickersLength)
}

// SetMaxTickersLength is a paid mutator transaction binding the contract method 0x47b8183e.
//
// Solidity: function setMaxTickersLength(uint256 _maxTickersLength) returns()
func (_Pricefeeds *PricefeedsSession) SetMaxTickersLength(_maxTickersLength *big.Int) (*types.Transaction, error) {
	return _Pricefeeds.Contract.SetMaxTickersLength(&_Pricefeeds.TransactOpts, _maxTickersLength)
}

// SetMaxTickersLength is a paid mutator transaction binding the contract method 0x47b8183e.
//
// Solidity: function setMaxTickersLength(uint256 _maxTickersLength) returns()
func (_Pricefeeds *PricefeedsTransactorSession) SetMaxTickersLength(_maxTickersLength *big.Int) (*types.Transaction, error) {
	return _Pricefeeds.Contract.SetMaxTickersLength(&_Pricefeeds.TransactOpts, _maxTickersLength)
}

// UpdatePrices is a paid mutator transaction binding the contract method 0xcb353e42.
//
// Solidity: function updatePrices((uint256,uint256,uint256,string,address,string,string,address)[] _TickerInfos) returns()
func (_Pricefeeds *PricefeedsTransactor) UpdatePrices(opts *bind.TransactOpts, _TickerInfos []PriceFeedsTickerInfo) (*types.Transaction, error) {
	return _Pricefeeds.contract.Transact(opts, "updatePrices", _TickerInfos)
}

// UpdatePrices is a paid mutator transaction binding the contract method 0xcb353e42.
//
// Solidity: function updatePrices((uint256,uint256,uint256,string,address,string,string,address)[] _TickerInfos) returns()
func (_Pricefeeds *PricefeedsSession) UpdatePrices(_TickerInfos []PriceFeedsTickerInfo) (*types.Transaction, error) {
	return _Pricefeeds.Contract.UpdatePrices(&_Pricefeeds.TransactOpts, _TickerInfos)
}

// UpdatePrices is a paid mutator transaction binding the contract method 0xcb353e42.
//
// Solidity: function updatePrices((uint256,uint256,uint256,string,address,string,string,address)[] _TickerInfos) returns()
func (_Pricefeeds *PricefeedsTransactorSession) UpdatePrices(_TickerInfos []PriceFeedsTickerInfo) (*types.Transaction, error) {
	return _Pricefeeds.Contract.UpdatePrices(&_Pricefeeds.TransactOpts, _TickerInfos)
}

// PricefeedsPriceUpdatedIterator is returned from FilterPriceUpdated and is used to iterate over the raw logs and unpacked data for PriceUpdated events raised by the Pricefeeds contract.
type PricefeedsPriceUpdatedIterator struct {
	Event *PricefeedsPriceUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PricefeedsPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PricefeedsPriceUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PricefeedsPriceUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PricefeedsPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PricefeedsPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PricefeedsPriceUpdated represents a PriceUpdated event raised by the Pricefeeds contract.
type PricefeedsPriceUpdated struct {
	Sender     common.Address
	Timestamp  *big.Int
	UpdateNums *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPriceUpdated is a free log retrieval operation binding the contract event 0xb556fac599c3c70efb9ab1fa725ecace6c81cc48d1455f886607def065f3e0c0.
//
// Solidity: event PriceUpdated(address indexed sender, uint256 timestamp, uint256 updateNums)
func (_Pricefeeds *PricefeedsFilterer) FilterPriceUpdated(opts *bind.FilterOpts, sender []common.Address) (*PricefeedsPriceUpdatedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pricefeeds.contract.FilterLogs(opts, "PriceUpdated", senderRule)
	if err != nil {
		return nil, err
	}
	return &PricefeedsPriceUpdatedIterator{contract: _Pricefeeds.contract, event: "PriceUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceUpdated is a free log subscription operation binding the contract event 0xb556fac599c3c70efb9ab1fa725ecace6c81cc48d1455f886607def065f3e0c0.
//
// Solidity: event PriceUpdated(address indexed sender, uint256 timestamp, uint256 updateNums)
func (_Pricefeeds *PricefeedsFilterer) WatchPriceUpdated(opts *bind.WatchOpts, sink chan<- *PricefeedsPriceUpdated, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pricefeeds.contract.WatchLogs(opts, "PriceUpdated", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PricefeedsPriceUpdated)
				if err := _Pricefeeds.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePriceUpdated is a log parse operation binding the contract event 0xb556fac599c3c70efb9ab1fa725ecace6c81cc48d1455f886607def065f3e0c0.
//
// Solidity: event PriceUpdated(address indexed sender, uint256 timestamp, uint256 updateNums)
func (_Pricefeeds *PricefeedsFilterer) ParsePriceUpdated(log types.Log) (*PricefeedsPriceUpdated, error) {
	event := new(PricefeedsPriceUpdated)
	if err := _Pricefeeds.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
