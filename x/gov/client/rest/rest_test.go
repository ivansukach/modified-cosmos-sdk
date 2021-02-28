// +build norace

package rest_test

import (
	"fmt"

	"github.com/ivansukach/modified-cosmos-sdk/types/rest"
	"github.com/ivansukach/modified-cosmos-sdk/x/gov/types"
)

func (s *IntegrationTestSuite) TestLegacyGetVote() {
	val := s.network.Validators[0]
	voterAddressBech32 := val.Address.String()

	testCases := []struct {
		name      string
		url       string
		expErr    bool
		expErrMsg string
	}{
		{
			"get non existing proposal",
			fmt.Sprintf("%s/gov/proposals/%s/votes/%s", val.APIAddress, "10", voterAddressBech32),
			true, "proposalID 10 does not exist",
		},
		{
			"get proposal with wrong voter address",
			fmt.Sprintf("%s/gov/proposals/%s/votes/%s", val.APIAddress, "1", "wrongVoterAddress"),
			true, "decoding bech32 failed: string not all lowercase or all uppercase",
		},
		{
			"get proposal with id",
			fmt.Sprintf("%s/gov/proposals/%s/votes/%s", val.APIAddress, "1", voterAddressBech32),
			false, "",
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			respJSON, err := rest.GetRequest(tc.url)
			s.Require().NoError(err)

			if tc.expErr {
				var errResp rest.ErrorResponse
				s.Require().NoError(val.ClientCtx.LegacyAmino.UnmarshalJSON(respJSON, &errResp))

				s.Require().Equal(errResp.Error, tc.expErrMsg)
			} else {
				var resp = rest.ResponseWithHeight{}
				err = val.ClientCtx.LegacyAmino.UnmarshalJSON(respJSON, &resp)
				s.Require().NoError(err)

				// Check result is not empty.
				var vote types.Vote
				s.Require().NoError(val.ClientCtx.LegacyAmino.UnmarshalJSON(resp.Result, &vote))
				s.Require().Equal(val.Address.String(), vote.Voter)
				// Note that option is now an int.
				s.Require().Equal(types.VoteOption(1), vote.Option)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestLegacyGetVotes() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		url       string
		expErr    bool
		expErrMsg string
	}{
		{
			"votes with empty proposal id",
			fmt.Sprintf("%s/gov/proposals/%s/votes", val.APIAddress, ""),
			true, "'votes' is not a valid uint64",
		},
		{
			"get votes with valid id",
			fmt.Sprintf("%s/gov/proposals/%s/votes", val.APIAddress, "1"),
			false, "",
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			respJSON, err := rest.GetRequest(tc.url)
			s.Require().NoError(err)

			if tc.expErr {
				var errResp rest.ErrorResponse
				s.Require().NoError(val.ClientCtx.LegacyAmino.UnmarshalJSON(respJSON, &errResp))

				s.Require().Equal(errResp.Error, tc.expErrMsg)
			} else {
				var resp = rest.ResponseWithHeight{}
				err = val.ClientCtx.LegacyAmino.UnmarshalJSON(respJSON, &resp)
				s.Require().NoError(err)

				// Check result is not empty.
				var votes []types.Vote
				s.Require().NoError(val.ClientCtx.LegacyAmino.UnmarshalJSON(resp.Result, &votes))
				s.Require().Greater(len(votes), 0)
			}
		})
	}
}
