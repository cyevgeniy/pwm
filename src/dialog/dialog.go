package dialog

import (
	"os"
    "golang.org/x/term"
    "github.com/cyevgeniy/pwm/pwmerrors"
    "github.com/cyevgeniy/pwm/ui"
)

// Prompts for master password.
// Returns error if passwords mismatched
func PromptForMasterPassword(duplicate bool) ([]byte, error) {
	key, err := PromptForPassword("Enter your master password: ")
	if err != nil {
		return nil, err
	}

	if !duplicate {
		return key, nil
	}

	keyD, err := PromptForPassword("Repeat your master password: ")
	if err != nil {
		return nil, err
	}

	if string(key) != string(keyD) {
		return nil, pwmerrors.ErrPassMismatch
	}

	return key, nil
}

func PromptForPassword(msg string) ([]byte, error) {
    ui.Println(msg)

	return term.ReadPassword(int(os.Stdin.Fd()))
}
