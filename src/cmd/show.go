package cmd

import (
    "github.com/cyevgeniy/pwm/dialog"
    "github.com/cyevgeniy/pwm/pwmerrors"
    "github.com/cyevgeniy/pwm/store"
    "github.com/cyevgeniy/pwm/ui"
    "github.com/cyevgeniy/pwm/utils"
    "github.com/spf13/cobra"
)

type showCmd struct {
    keyFile string
}


func init() {
    rootCmd.AddCommand(newShowCmd())
}

func newShowCmd() *cobra.Command {
    cc := &showCmd{}

    cmd := &cobra.Command{
        Use:   "show",
        Short: "Show",
        Long:  "Show password",
        Run: func(cmd *cobra.Command, args []string) {
            s, err := store.GetStore()

            utils.CheckErr(err)

            if len(args) == 0 {
                utils.CheckErr(pwmerrors.ErrNoPassFileProvided)
            }

            var key []byte

            if cc.keyFile == "" {
                key, err = dialog.PromptForMasterPassword(false)
            } else {
                key, err = utils.ReadKeyFile(cc.keyFile)
            }

            utils.CheckErr(err)

            message, err := s.Decrypt(args[0], key)

            utils.CheckErr(err)

            ui.Println(message)

        },
    }

    cmd.Flags().StringVarP(&cc.keyFile, "input", "i", "", "File with a master password")

    return cmd
}
