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

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy password to clipboard",
	Long:  `Copy password to clipboard`,
	Run: func(cmd *cobra.Command, args []string) {
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

		clipboard.Write(clipboard.FmtText, []byte(message))

	},
}
