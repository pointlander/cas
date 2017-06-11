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
	"github.com/dop251/goja"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	vm := goja.New()
	_, err := vm.RunString("window = {};")
	if err != nil {
		log.Fatal(err)
	}
	// for license see: https://github.com/davidedc/Algebrite
	algebrite, err := algebriteBundleForBrowserJsBytes()
	if err != nil {
		log.Fatal(err)
	}

	program, err := goja.Compile("algebrite", string(algebrite), true)
	if err != nil {
		log.Fatal(err)
	}
	_, err = vm.RunProgram(program)
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

	window := vm.Get("window").ToObject(vm)
	alg := window.Get("Algebrite").ToObject(vm)
	run, valid := goja.AssertFunction(alg.Get("run"))
	if !valid {
		log.Fatal("window.Algebrite.run is not a function")
	}
	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}

		result, err := run(goja.Null(), vm.ToValue(line))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result.Export())
	}
}
