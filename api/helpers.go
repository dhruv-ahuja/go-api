package api

import "log"

// checkErr checks for error in given functions/methods. It also outputs an
// error message, if given one.
func checkErr(errMsg string, err error) {
	if err != nil {
		log.Fatal(errMsg, err)
	}
}
