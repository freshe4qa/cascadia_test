package keeper

import (
	"encoding/json"
	"math/big"

	"github.com/cascadiafoundation/cascadia/contracts"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/evmos/ethermint/server/config"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

// CallEVM performs a smart contract method call using given args
func (k Keeper) CallEVM(
	ctx sdk.Context,
	abi abi.ABI,
	from, contract common.Address,
	commit bool,
	method string,
	args ...interface{},
) (*evmtypes.MsgEthereumTxResponse, error) {
	data, err := abi.Pack(method, args...)
	if err != nil {
		return nil, err
	}

	resp, err := k.CallEVMWithData(ctx, from, &contract, data, commit)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "contract call failed: method '%s', contract '%s'", method, contract)
	}
	return resp, nil
}

// CallEVMWithData performs a smart contract method call using contract data
func (k Keeper) CallEVMWithData(
	ctx sdk.Context,
	from common.Address,
	contract *common.Address,
	data []byte,
	commit bool,
) (*evmtypes.MsgEthereumTxResponse, error) {
	nonce, err := k.authKeeper.GetSequence(ctx, from.Bytes())
	if err != nil {
		return nil, err
	}

	gasCap := config.DefaultGasCap
	if commit {
		args, err := json.Marshal(evmtypes.TransactionArgs{
			From: &from,
			To:   contract,
			Data: (*hexutil.Bytes)(&data),
		})

		if err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "failed to marshal tx args: %s", err.Error())
		}

		gasRes, err := k.evmKeeper.EstimateGas(sdk.WrapSDKContext(ctx), &evmtypes.EthCallRequest{
			Args:   args,
			GasCap: config.DefaultGasCap,
		})

		if err != nil {
			return nil, err
		}
		gasCap = gasRes.Gas
	}

	msg := ethtypes.NewMessage(
		from,
		contract,
		nonce,
		big.NewInt(0), // amount
		gasCap,        // gasLimit
		big.NewInt(0), // gasFeeCap
		big.NewInt(0), // gasTipCap
		big.NewInt(0), // gasPrice
		data,
		ethtypes.AccessList{}, // AccessList
		!commit,               // isFake
	)

	res, err := k.evmKeeper.ApplyMessage(ctx, msg, evmtypes.NewNoOpTracer(), commit)
	if err != nil {
		return nil, err
	}

	if res.Failed() {
		return nil, sdkerrors.Wrap(evmtypes.ErrVMExecution, res.VmError)
	}

	return res, nil
}

// DeployERC20Contract creates and deploys an ERC20 contract on the EVM with the
// erc20 module account as owner.
func (k Keeper) DeployVotingEscrowContract(
	ctx sdk.Context,
) (common.Address, error) {
	ctorArgs, err := contracts.VotingEscrowContract.ABI.Pack(
		"", "veCC", "veCC", "1.0",
	)

	if err != nil {
		return common.Address{}, err
	}

	data := make([]byte, len(contracts.VotingEscrowContract.Bin)+len(ctorArgs))

	copy(data[:len(contracts.VotingEscrowContract.Bin)], contracts.VotingEscrowContract.Bin)
	copy(data[len(contracts.VotingEscrowContract.Bin):], ctorArgs)

	nonce, err := k.authKeeper.GetSequence(ctx, ModuleAddress.Bytes())

	if err != nil {
		return common.Address{}, err
	}

	contractAddr := crypto.CreateAddress(ModuleAddress, nonce)

	_, err = k.CallEVMWithData(ctx, ModuleAddress, nil, data, true)

	if err != nil {
		return common.Address{}, err
	}

	return contractAddr, nil
}

// BalanceOf queries an account's balance for a given ERC20 contract
// func (k Keeper) BalanceOf(
// 	ctx sdk.Context,
// 	abi abi.ABI,
// 	contract, account common.Address,
// ) *big.Int {
// 	res, err := k.CallEVM(ctx, abi, ModuleAddress, contract, false, "balanceOf", account)
// 	if err != nil {
// 		return nil
// 	}

// 	unpacked, err := abi.Unpack("balanceOf", res.Ret)
// 	if err != nil || len(unpacked) == 0 {
// 		return nil
// 	}

// 	balance, ok := unpacked[0].(*big.Int)
// 	if !ok {
// 		return nil
// 	}

// 	return balance
// }
