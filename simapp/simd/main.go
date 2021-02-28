package main

import (
	"os"

	"github.com/ivansukach/modified-cosmos-sdk/server"
	svrcmd "github.com/ivansukach/modified-cosmos-sdk/server/cmd"
	"github.com/ivansukach/modified-cosmos-sdk/simapp"
	"github.com/ivansukach/modified-cosmos-sdk/simapp/simd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, simapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
