// Copyright 2016 The CAS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/chzyer/readline"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

type CAS interface {
	Compile(algebrite []byte) error
	Load() error
	Run(line string) (string, error)
}

func main() {
	cas := NewGOJA()

	// for license see: https://github.com/davidedc/Algebrite
	algebrite, err := algebriteBundleForBrowserJsBytes()
	if err != nil {
		log.Panic(err)
	}

	err = cas.Compile(algebrite)
	if err != nil {
		log.Panic(err)
	}
	err = cas.Load()
	if err != nil {
		log.Panic(err)
	}

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	rl, err := readline.New("> ")
	if err != nil {
		log.Panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}

		result, err := cas.Run(line)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(result)
	}
}
