package cmd

import (
	"github.com/cyevgeniy/pwm/dialog"
	"github.com/cyevgeniy/pwm/pwmerrors"
	"github.com/cyevgeniy/pwm/store"
	"github.com/cyevgeniy/pwm/ui"
	"github.com/cyevgeniy/pwm/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show",
	Long:  "Show password",
	Run: func(cmd *cobra.Command, args []string) {
		s, err := store.GetStore()

		utils.CheckErr(err)

		if len(args) == 0 {
			utils.CheckErr(pwmerrors.ErrNoPassFileProvided)
		}

		key, err := dialog.PromptForMasterPassword(false)

		utils.CheckErr(err)

		message, err := s.Decrypt(args[0], key)

		utils.CheckErr(err)

		ui.Println(message)

	},
}
