// +build ignore
// Copyright 2016 The CAS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
)

func build() {
	gopath := os.Getenv("GOPATH")
	gocmd, err := exec.LookPath("go")
	if err != nil {
		panic(err)
	}
	fileb0x := path.Join(gopath, "bin", "fileb0x")

	_, err = os.Stat(fileb0x)
	if err != nil {
		fmt.Println("installing github.com/UnnoTed/fileb0x")
		cmd := exec.Command(gocmd, "get", "-u", "github.com/UnnoTed/fileb0x")
		err = cmd.Run()
		if err != nil {
			panic(err)
		}
	}

	_, err = os.Stat("algebrite.bundle-for-browser.js")
	if err != nil {
		fmt.Println("fetching algebrite.bundle-for-browser.js")
		resp, err1 := http.Get("https://raw.githubusercontent.com/davidedc/Algebrite/master/dist/algebrite.bundle-for-browser.js")
		if err1 != nil {
			panic(err1)
		}
		out, err1 := os.Create("algebrite.bundle-for-browser.js")
		if err1 != nil {
			panic(err1)
		}
		_, err1 = io.Copy(out, resp.Body)
		if err1 != nil {
			panic(err1)
		}
		err1 = out.Close()
		if err1 != nil {
			panic(err1)
		}
	}

	fmt.Println(fileb0x)
	cmd := exec.Command(fileb0x, "fileb0x.json")
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println("building...")
	cmd = exec.Command(gocmd, "build")
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}

func clean() {
	os.Remove("algebrite.bundle-for-browser.js")
	os.Remove("cas")
}

func main() {
	cmd := ""
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}
	switch cmd {
	case "build":
		build()
	case "clean":
		clean()
	default:
		build()
	}
}
