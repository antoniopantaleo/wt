//go:build !debug

package debuglog

func Print(v ...any) {}

func Printf(format string, v ...any) {}
