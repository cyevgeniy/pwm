package cmd

import (
	"fmt"
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/cyevgeniy/pwm/dialog"
	"github.com/cyevgeniy/pwm/store"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new password to the store",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		s, err := store.GetStore()

		if err != nil {
			log.Fatal("Can't reach password store")
		}

		fname := args[0]

		// If exists, show alert
		if s.IsFileExists(fname) {
			log.Fatal("Such password file already exists")
		}

		// If it's not exist, ask for password
		pass, err := dialog.PromptForPassword("Enter your password: ")

		if err != nil {
			log.Fatal("Error during reading password")
		}

		key, err := dialog.PromptForMasterPassword(true)

		if err != nil {
			log.Fatal(err.Error())
		}

		// Encrypt with password:
		armor, err := helper.EncryptMessageWithPassword(key, string(pass))

		if err != nil {
			log.Fatal("Error during password's encryption")
		}

		err = s.WriteFile(fname, []byte(armor))

		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("Done!")
	},
}
