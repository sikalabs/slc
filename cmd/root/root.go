package root

import (
	"github.com/sikalabs/slc/version"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "slc",
	Short: "SikaLabs Cloud CLI, " + version.Version,
}
