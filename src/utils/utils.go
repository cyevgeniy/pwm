package utils

import (
	"github.com/cyevgeniy/pwm/ui"
	"os"
)

func CheckErr(e error) {
	if e != nil {
		ui.ErrPrintln(e.Error())

		os.Exit(1)
	}
}
