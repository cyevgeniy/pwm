package cmd

import (
	"github.com/cyevgeniy/pwm/dialog"
	"github.com/cyevgeniy/pwm/store"
	"github.com/cyevgeniy/pwm/utils"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

func init() {
	rootCmd.AddCommand(copyCmd)
}

func runCopy(cmd *cobra.Command, args []string) {
	s, err := store.GetStore()
    utils.CheckErr(err)

	key, err := dialog.PromptForMasterPassword(false)
    utils.CheckErr(err)

	message, err := s.Decrypt(args[0], key)
    utils.CheckErr(err)

	// TODO: Rewrite this command - clear clipboard after some time
	clipboard.Write(clipboard.FmtText, []byte(message))

}

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy password to the clipboard",
	Long:  `Copy password to the clipboard`,
	Run:   runCopy,
}
