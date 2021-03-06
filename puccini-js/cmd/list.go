package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tliron/puccini/ard"
	"github.com/tliron/puccini/clout"
	"github.com/tliron/puccini/common"
	"github.com/tliron/puccini/js"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list [[Clout PATH or URL]]",
	Short: "List JavaScript in Clout",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var path string
		if len(args) == 1 {
			path = args[0]
		}

		c, err := ReadClout(path)
		common.ValidateError(err)

		List(c)
	},
}

func List(c *clout.Clout) {
	metadata, err := js.GetMetadata(c)
	common.ValidateError(err)

	ListValue(metadata, nil)
}

func ListValue(value interface{}, path []string) {
	switch v := value.(type) {
	case string:
		if !common.Quiet {
			fmt.Printf("%s\n", strings.Join(path, "."))
		}
	case ard.Map:
		for key, vv := range v {
			ListValue(vv, append(path, key))
		}
	}
}
