package cmd

import (
	"github.com/cyevgeniy/pwm/store"
	"github.com/cyevgeniy/pwm/ui"
	"github.com/cyevgeniy/pwm/utils"
	"github.com/spf13/cobra"
	"strings"
)

var filter string

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().StringVarP(&filter, "filter", "f", "", "Filter for filenames")
}

func runLs(cmd *cobra.Command, args []string) {
	s, err := store.GetStore()
	utils.CheckErr(err)

	files, err := s.ListFiles()
	utils.CheckErr(err)

	for i := range files {
		fname := files[i].Name()
		if (filter == "") || (filter != "" && strings.Contains(fname, filter)) {
			ui.Println(fname)
		}
	}
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Print list of passwords",
	Long:  `Prints lists of all password files in the store`,
	Run:   runLs,
}
