package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/tendermint/tmlibs/cli"

	cmd "github.com/tendermint/tendermint/cmd/tendermint/commands"
	cfg "github.com/tendermint/tendermint/config"
	nm "github.com/tendermint/tendermint/node"
	_ "net/http/pprof"
)

func main() {
	go func() {
		http.ListenAndServe(":2020", nil)
	}()

	cmd.AddInitFlags(cmd.InitFilesCmd)

	rootCmd := cmd.RootCmd
	rootCmd.AddCommand(
		cmd.InitFilesCmd,
		cmd.ProbeUpnpCmd,
		cmd.ResetAllCmd,
		cmd.ResetPrivValidatorCmd,
		cmd.ShowValidatorCmd,
		cmd.ShowNodeIDCmd,
		cmd.VersionCmd)

	// NOTE:
	// Users wishing to:
	//	* Use an external signer for their validators
	//	* Supply an in-proc abci app
	//	* Supply a genesis doc file from another source
	//	* Provide their own DB implementation
	// can copy this file and use something other than the
	// DefaultNewNode function
	nodeFunc := nm.DefaultNewNode

	// Create & start node
	rootCmd.AddCommand(
		cmd.NewSyncNodeCmd(nodeFunc),
		cmd.NewRunNodeCmd(nodeFunc))

	command := cli.PrepareBaseCmd(rootCmd, "TM", os.ExpandEnv(filepath.Join("$HOME", cfg.DefaultTendermintDir)))
	if err := command.Execute(); err != nil {
		panic(err)
	}
}
