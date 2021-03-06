package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tliron/puccini/common"
	"github.com/tliron/puccini/format"
	"github.com/tliron/puccini/js"
	"github.com/tliron/puccini/url"
)

func init() {
	rootCmd.AddCommand(putCmd)
	putCmd.Flags().StringVarP(&output, "output", "o", "", "output Clout to file (instead of stdout)")
}

var putCmd = &cobra.Command{
	Use:   "put [COMMAND] [JavaScript PATH or URL] [[Clout PATH or URL]]",
	Short: "Put JavaScript in Clout",
	Long:  ``,
	Args:  cobra.RangeArgs(2, 3),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		jsUrl := args[1]

		var cloutPath string
		if len(args) == 3 {
			cloutPath = args[2]
		}

		c, err := ReadClout(cloutPath)
		common.ValidateError(err)

		url_, err := url.NewValidURL(jsUrl, nil)
		common.ValidateError(err)

		sourceCode, err := url.Read(url_)
		common.ValidateError(err)

		err = js.SetScriptSourceCode(name, js.Cleanup(sourceCode), c)
		common.ValidateError(err)

		if !common.Quiet || (output != "") {
			err = format.WriteOrPrint(c, ardFormat, true, output)
			common.ValidateError(err)
		}
	},
}
