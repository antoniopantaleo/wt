//go:build debug

package debuglog

import "log"

func Print(v ...any) {
	log.Print(v...)
}

func Printf(format string, v ...any) {
	log.Printf(format, v...)
}
