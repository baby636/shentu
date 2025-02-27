package gov

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/certikfoundation/shentu/v2/x/gov/keeper"
	"github.com/certikfoundation/shentu/v2/x/gov/types"
)

// InitGenesis stores genesis parameters.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, ak govTypes.AccountKeeper, bk govTypes.BankKeeper, data types.GenesisState) {
	k.SetProposalID(ctx, data.StartingProposalId)
	k.SetDepositParams(ctx, data.DepositParams)
	k.SetVotingParams(ctx, data.VotingParams)
	k.SetTallyParams(ctx, data.TallyParams)

	// check if the deposits pool account exists
	moduleAcc := k.GetGovernanceAccount(ctx)
	if moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", govTypes.ModuleName))
	}

	var totalDeposits sdk.Coins
	for _, deposit := range data.Deposits {
		k.SetDeposit(ctx, deposit)
		totalDeposits = totalDeposits.Add(deposit.Amount...)
	}

	for _, vote := range data.Votes {
		k.SetVote(ctx, vote)
	}

	for _, proposal := range data.Proposals {
		switch proposal.Status {
		case types.StatusDepositPeriod:
			k.InsertInactiveProposalQueue(ctx, proposal.ProposalId, proposal.DepositEndTime)
		case types.StatusCertifierVotingPeriod, types.StatusValidatorVotingPeriod:
			k.InsertActiveProposalQueue(ctx, proposal.ProposalId, proposal.VotingEndTime)
		}
		k.SetProposal(ctx, proposal)
	}

	// add coins if not provided on genesis
	if bk.GetAllBalances(ctx, moduleAcc.GetAddress()).IsZero() {
		if err := bk.SetBalances(ctx, moduleAcc.GetAddress(), totalDeposits); err != nil {
			panic(err)
		}

		ak.SetModuleAccount(ctx, moduleAcc)
	}
}

// ExportGenesis writes the current store values to a genesis file, which can be imported again with InitGenesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	startingProposalID, _ := k.GetProposalID(ctx)
	depositParams := k.GetDepositParams(ctx)
	votingParams := k.GetVotingParams(ctx)
	tallyParams := k.GetTallyParams(ctx)
	proposals := k.GetProposals(ctx)

	var genState types.GenesisState

	for _, proposal := range proposals {
		genState.Deposits = append(genState.Deposits, k.GetDepositsByProposalID(ctx, proposal.ProposalId)...)
		genState.Votes = append(genState.Votes, k.GetVotes(ctx, proposal.ProposalId)...)
	}
	genState.StartingProposalId = startingProposalID
	genState.Proposals = proposals
	genState.DepositParams = depositParams
	genState.VotingParams = votingParams
	genState.TallyParams = tallyParams

	return &genState
}
