package keeper

import (
	"math/big"

	"github.com/cascadiafoundation/cascadia/contracts"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/ethereum/go-ethereum/common"
)

// TODO: Break into several smaller functions for clarity

// Tally iterates over the votes and updates the tally of a proposal based on the voting power of the
// voters
func (keeper Keeper) Tally(ctx sdk.Context, proposal types.Proposal) (passes bool, burnDeposits bool, tallyResults types.TallyResult) {
	results := make(map[types.VoteOption]sdk.Dec)
	totalVotingPower := sdk.ZeroDec()
	results[types.OptionYes] = sdk.ZeroDec()
	results[types.OptionAbstain] = sdk.ZeroDec()
	results[types.OptionNo] = sdk.ZeroDec()
	results[types.OptionNoWithVeto] = sdk.ZeroDec()

	// Get the vote-escrowed contract
	contract, found := keeper.rk.GetReward(ctx, "vecontract")
	if !found {
		panic("no ve contract")
	}
	contractEvmAddr := common.HexToAddress(contract.Contract)

	// Get the time of submission of proposal
	submitTime := big.NewInt(proposal.SubmitTime.Unix())

	// Get the total voting power
	totalBalance := sdk.NewDecFromBigInt(keeper.rk.TotalSupply(ctx, contracts.VotingEscrowContract.ABI, contractEvmAddr, submitTime))

	keeper.IterateVotes(ctx, proposal.ProposalId, func(vote types.Vote) bool {
		// if validator, just record it in the map
		voter := sdk.MustAccAddressFromBech32(vote.Voter)
		voterEvmAddr := common.BytesToAddress(voter.Bytes())
		voterBalance := sdk.NewDecFromBigInt(keeper.rk.BalanceOf(ctx, contracts.VotingEscrowContract.ABI, contractEvmAddr, voterEvmAddr, submitTime))

		for _, option := range vote.Options {

			results[option.Option] = results[option.Option].Add(voterBalance)
		}

		totalVotingPower.Add(voterBalance)

		keeper.deleteVote(ctx, vote.ProposalId, voter)
		return false
	})

	tallyParams := keeper.GetTallyParams(ctx)
	tallyResults = types.NewTallyResultFromMap(results)

	// TODO: Upgrade the spec to cover all of these cases & remove pseudocode.
	// If there is no staked coins, the proposal fails
	if totalBalance.IsZero() {
		return false, false, tallyResults
	}

	// If there is not enough quorum of votes, the proposal fails
	percentVoting := totalVotingPower.Quo(totalBalance)
	if percentVoting.LT(tallyParams.Quorum) {
		return false, true, tallyResults
	}

	// If no one votes (everyone abstains), proposal fails
	if totalVotingPower.Sub(results[types.OptionAbstain]).Equal(sdk.ZeroDec()) {
		return false, false, tallyResults
	}

	// If more than 1/3 of voters veto, proposal fails
	if results[types.OptionNoWithVeto].Quo(totalVotingPower).GT(tallyParams.VetoThreshold) {
		return false, true, tallyResults
	}

	// If more than 1/2 of non-abstaining voters vote Yes, proposal passes
	if results[types.OptionYes].Quo(totalVotingPower.Sub(results[types.OptionAbstain])).GT(tallyParams.Threshold) {
		return true, false, tallyResults
	}

	// If more than 1/2 of non-abstaining voters vote No, proposal fails
	return false, false, tallyResults
}
