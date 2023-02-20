package cmd

import (
	"github.com/cyevgeniy/pwm/dialog"
	"github.com/cyevgeniy/pwm/store"
	"github.com/cyevgeniy/pwm/ui"
	"github.com/cyevgeniy/pwm/utils"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

var useMeta bool

func init() {
	rootCmd.AddCommand(copyCmd)
	copyCmd.Flags().BoolVarP(&useMeta, "meta", "m", false, "Show metadata after copying")
}

func runCopy(cmd *cobra.Command, args []string) {
	s, err := store.GetStore()
	utils.CheckErr(err)

	key, err := dialog.PromptForMasterPassword(false)
	utils.CheckErr(err)

	fname := args[0]

	message, err := s.Decrypt(fname, key)
	utils.CheckErr(err)

	// TODO: Rewrite this command - clear clipboard after some time
	clipboard.Write(clipboard.FmtText, []byte(message))

	if useMeta {
		message, err := s.Decrypt(fname+".meta", key)
		utils.CheckErr(err)

		ui.Println(message)
	}
}

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy password to the clipboard",
	Long:  `Copy password to the clipboard`,
	Run:   runCopy,
}
