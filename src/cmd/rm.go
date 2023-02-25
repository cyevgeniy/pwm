package cmd

import (
	"github.com/cyevgeniy/pwm/pwmerrors"
	"github.com/cyevgeniy/pwm/store"
	"github.com/cyevgeniy/pwm/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rmCmd)
}

func runRm(cmd *cobra.Command, args []string) {
	s, err := store.GetStore()
	utils.CheckErr(err)

	if len(args) == 0 {
		utils.CheckErr(pwmerrors.ErrNoPassFileProvided)
	}

	utils.CheckErr(s.RemoveFile(args[0]))
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete specified file",
	Long:  `Delete password file. This command does not delete metadata`,
	Run:   runRm,
}
