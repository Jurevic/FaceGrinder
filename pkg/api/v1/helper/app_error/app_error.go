package app_error

import "log"

func Fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}