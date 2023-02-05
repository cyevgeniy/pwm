package cmd

import (
	"fmt"
	"github.com/cyevgeniy/pwm/dialog"
	"github.com/cyevgeniy/pwm/store"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show",
	Long:  `Show password`,
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

		fmt.Println(message)

	},
}
