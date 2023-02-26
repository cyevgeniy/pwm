package cmd

import (
    "github.com/cyevgeniy/pwm/dialog"
    "github.com/cyevgeniy/pwm/store"
    "github.com/cyevgeniy/pwm/ui"
    "github.com/cyevgeniy/pwm/utils"
    "github.com/spf13/cobra"
    "golang.design/x/clipboard"
)

type copyCmd struct {
    useMeta bool
    keyFile string
}

func init() {
    rootCmd.AddCommand(newCopyCmd())
}


func newCopyCmd() *cobra.Command {
    cc := copyCmd{}

    cmd := &cobra.Command{
        Use:   "copy",
        Short: "Copy password to the clipboard",
        Long:  `Copy password to the clipboard`,
        Run:   func(cmd *cobra.Command, args []string) {
            s, err := store.GetStore()
            utils.CheckErr(err)

            var key []byte

            if cc.keyFile == "" {
                key, err = dialog.PromptForMasterPassword(false)
            } else {
                key, err = utils.ReadKeyFile(cc.keyFile)
            }

            utils.CheckErr(err)

            fname := args[0]

            message, err := s.Decrypt(fname, key)
            utils.CheckErr(err)

            // TODO: Rewrite this command - clear clipboard after some time
            clipboard.Write(clipboard.FmtText, []byte(message))

            if cc.useMeta {
                message, err := s.Decrypt(fname+".meta", key)
                utils.CheckErr(err)

                ui.Println(message)
            }
        },

    }

    cmd.Flags().BoolVarP(&cc.useMeta, "meta", "m", false, "Show metadata after copying")
    cmd.Flags().StringVarP(&cc.keyFile, "input", "i", "", "File with a master password")

    return cmd
}
