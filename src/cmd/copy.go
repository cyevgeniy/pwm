package cmd

import (
	"github.com/cyevgeniy/pwm/dialog"
	"github.com/cyevgeniy/pwm/store"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
	"log"
)

func init() {
	rootCmd.AddCommand(copyCmd)
}

func runCopy(cmd *cobra.Command, args []string) {
	s, err := store.GetStore()

	if err != nil {
		log.Fatal("Can't reach password store")
	}

	key, err := dialog.PromptForMasterPassword(false)

	if err != nil {
		log.Fatal(err.Error())
	}

	message, err := s.Decrypt(args[0], key)

	if err != nil {
		log.Fatal(err.Error())
	}

	// TODO: Rewrite this command - clear clipboard after some time
	clipboard.Write(clipboard.FmtText, []byte(message))

}

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy password to clipboard",
	Long:  `Copy password to clipboard`,
	Run:   runCopy,
}
