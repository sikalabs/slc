package cmd

import (
	"github.com/sikalabs/slc/cmd/root"
	_ "github.com/sikalabs/slc/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.RootCmd.Execute())
}
