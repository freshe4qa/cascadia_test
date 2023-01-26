package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	"github.com/cascadiafoundation/cascadia/x/erc20/types"
)

var (
	//go:embed compiled_contracts/VotingEscrow.json
	VotingEscrowJSON []byte // nolint: golint

	// VotingEscrowContract is the compiled erc20 contract
	VotingEscrowContract evmtypes.CompiledContract

	// VotingEscrowAddress is the erc20 module address
	VotingEscrowAddress common.Address
)

func init() {
	VotingEscrowAddress = types.ModuleAddress

	err := json.Unmarshal(VotingEscrowJSON, &VotingEscrowContract)
	if err != nil {
		panic(err)
	}

	if len(VotingEscrowContract.Bin) == 0 {
		panic("load contract failed")
	}
}
