package client

import (
	govclient "github.com/ivansukach/modified-cosmos-sdk/x/gov/client"
	"github.com/ivansukach/modified-cosmos-sdk/x/params/client/cli"
	"github.com/ivansukach/modified-cosmos-sdk/x/params/client/rest"
)

// ProposalHandler handles param change proposals
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
