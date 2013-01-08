// Copyright 2013, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postkr

import (
	"io"
	"log"
)

const (
	// Log level flags for flag of logger
	LLerror = 1 << (iota + 8)
	LLwarnning
	LLdebug
	LLall = LLerror | LLwarnning | LLdebug
)

var logger *log.Logger

// Set postkr's log out.
// For example, to log on stderr;
//
//     postkr.LogTo(os.Stderr)
//
func LogTo(out io.Writer) {
	logger = log.New(out, "postkr ", log.LstdFlags|LLall)
}

func logE(format string, v ...interface{}) {
	if logger == nil {
		return
	}
	if logger.Flags()&LLerror != 0 {
		logger.Printf("E:"+format+"\n", v...)
	}
}

func logW(format string, v ...interface{}) {
	if logger == nil {
		return
	}
	if logger.Flags()&LLwarnning != 0 {
		logger.Printf("W:"+format+"\n", v...)
	}
}

func logD(format string, v ...interface{}) {
	if logger == nil {
		return
	}
	if logger.Flags()&LLdebug != 0 {
		logger.Printf("D:"+format+"\n", v...)
	}
}
