package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/cascadiafoundation/cascadia/v1/x/erc20/client/cli"
	"github.com/cascadiafoundation/cascadia/v1/x/erc20/client/rest"
)

var (
	RegisterCoinProposalHandler          = govclient.NewProposalHandler(cli.NewRegisterCoinProposalCmd, rest.RegisterCoinProposalRESTHandler)
	RegisterERC20ProposalHandler         = govclient.NewProposalHandler(cli.NewRegisterERC20ProposalCmd, rest.RegisterERC20ProposalRESTHandler)
	ToggleTokenConversionProposalHandler = govclient.NewProposalHandler(cli.NewToggleTokenConversionProposalCmd, rest.ToggleTokenConversionRESTHandler)
)
