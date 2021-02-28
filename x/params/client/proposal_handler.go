package client

import (
	govclient "github.com/ivansukach/modified-cosmos-sdk/x/gov/client"
	"github.com/ivansukach/modified-cosmos-sdk/x/params/client/cli"
	"github.com/ivansukach/modified-cosmos-sdk/x/params/client/rest"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd, rest.ProposalRESTHandler)
