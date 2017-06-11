package main

import (
	"errors"
	"log"

	"github.com/dop251/goja"
)

type GOJA struct {
	vm      *goja.Runtime
	program *goja.Program
}

func NewGOJA() CAS {
	vm := goja.New()
	_, err := vm.RunString("window = {};")
	if err != nil {
		log.Panic(err)
	}
	return &GOJA{
		vm: vm,
	}
}

func (g *GOJA) Compile(algebrite []byte) error {
	program, err := goja.Compile("algebrite", string(algebrite), true)
	if err != nil {
		return err
	}
	g.program = program
	return nil
}

func (g *GOJA) Load() error {
	_, err := g.vm.RunProgram(g.program)
	if err != nil {
		return err
	}
	return nil
}

func (g *GOJA) Run(line string) (string, error) {
	vm := g.vm
	window := vm.Get("window").ToObject(vm)
	alg := window.Get("Algebrite").ToObject(vm)
	run, valid := goja.AssertFunction(alg.Get("run"))
	if !valid {
		return "", errors.New("window.Algebrite.run is not a function")
	}
	result, err := run(goja.Null(), vm.ToValue(line))
	if err != nil {
		return "", err
	}
	return result.String(), nil
}
