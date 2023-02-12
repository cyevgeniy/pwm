package cmd

import (
	"errors"
	"fmt"
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/cyevgeniy/pwm/dialog"
	"github.com/cyevgeniy/pwm/store"
	"github.com/cyevgeniy/pwm/ui"
	"github.com/cyevgeniy/pwm/utils"
	"github.com/cyevgeniy/pwm/pwmerrors"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updCmd)
}

func runUpd(cmd *cobra.Command, args []string) {
	s, err := store.GetStore()
	utils.CheckErr(err)

	if len(args) == 0 {
		utils.CheckErr(pwmerrors.ErrNoPassFileProvided)
	}

	fname := args[0]

	// If exists, show alert
	if !s.IsFileExists(fname) {
		utils.CheckErr(errors.New("There's no such password file"))
	}

	key, err := dialog.PromptForMasterPassword(false)

	utils.CheckErr(err)

	// Try to decrypt
	_, err = s.Decrypt(args[0], key)
	utils.CheckErr(err)

	pass, err := dialog.PromptForPassword("Enter your password: ")

	utils.CheckErr(err)

	// Encrypt with password:
	armor, err := helper.EncryptMessageWithPassword(key, string(pass))

	utils.CheckErr(err)
	err = s.WriteFile(fname, []byte(armor))

	utils.CheckErr(err)

	if meta != "" {
		armMeta, err := helper.EncryptMessageWithPassword(key, string(meta))

		if err != nil {
			err = fmt.Errorf("Can't encrypt meta information. Error: %s", err.Error())
			utils.CheckErr(err)
		}

		err = s.WriteFile(fname+".meta", []byte(armMeta))
		utils.CheckErr(err)
	}

	ui.Println("Done!")

}

var updCmd = &cobra.Command{
	Use:   "upd",
	Short: "Update password or metainformation",
	Long:  `Update password or metainformation`,
	Run:   runUpd,
}
