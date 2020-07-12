package util

import "log"

// DieIf logs and panics if err is not nil.
func DieIf(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// LogIf logs and panics if err is not nil.
func LogIf(err error) {
	if err != nil {
		log.Panic(err)
	}
}
