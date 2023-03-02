package main

import (
	"os"

	"github.com/althea-net/bech32-ibc/app"
	"github.com/althea-net/bech32-ibc/cmd/bech32ibcd/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
