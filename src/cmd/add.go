package cmd

import (
	"fmt"
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/cyevgeniy/pwm/dialog"
	"github.com/cyevgeniy/pwm/store"
	"github.com/cyevgeniy/pwm/utils"
	"github.com/cyevgeniy/pwm/ui"
	"github.com/spf13/cobra"
	"log"
)

var meta string

func init() {
	rootCmd.AddCommand(addCmd)
    addCmd.Flags().StringVarP(&meta, "meta", "m", "", "Add additional meta-informatino")
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new password to the store",
	Long:  "Add new password to the store. Will not replace already existed passwords",
	Run: func(cmd *cobra.Command, args []string) {
		s, err := store.GetStore()

        utils.CheckErr(err)

        if len(args) == 0 {
            err := fmt.Errorf("You should specify filename for password")
            utils.CheckErr(err)
        }

		fname := args[0]

		// If exists, show alert
		if s.IsFileExists(fname) {
			log.Fatal("Such password file already exists")
		}

		// If it's not exist, ask for password
		pass, err := dialog.PromptForPassword("Enter your password: ")

        utils.CheckErr(err)

		key, err := dialog.PromptForMasterPassword(true)

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

            err = s.WriteFile(fname + ".meta", []byte(armMeta))
            utils.CheckErr(err)
        }


		ui.Println("Done!")
	},
}
