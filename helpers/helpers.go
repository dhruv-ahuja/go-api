package helpers

import "log"

// CheckErr checks for error in given functions/methods. It also outputs an
// error message, if given one.
func CheckErr(errMsg string, err error) {
	if err != nil {
		log.Fatal(errMsg, err)
	}
}
