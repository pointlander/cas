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
	"github.com/robertkrimen/otto"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	vm := otto.New()
	_, err := vm.Run("window = {};")
	if err != nil {
		log.Fatal(err)
	}
	// for license see: https://github.com/davidedc/Algebrite
	algebrite, err := algebriteBundleForBrowserJsBytes()
	if err != nil {
		log.Fatal(err)
	}
	script, err := vm.Compile("", algebrite)
	if err != nil {
		log.Fatal(err)
	}
	_, err = vm.Run(script)
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	rl, err := readline.New("> ")
	if err != nil {
		log.Fatal(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		result, err := vm.Call("window.Algebrite.run", nil, line)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
}
