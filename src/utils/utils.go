package utils

import (
        "os"
        "github.com/cyevgeniy/pwm/ui"
)

func CheckErr(e error) {
    if e != nil {
        ui.ErrPrintln(e.Error())

        os.Exit(1)
    }
}
