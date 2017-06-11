package main

import (
	"log"

	"github.com/robertkrimen/otto"
)

type Otto struct {
	vm      *otto.Otto
	program *otto.Script
}

func NewOtto() CAS {
	vm := otto.New()
	_, err := vm.Run("window = {};")
	if err != nil {
		log.Panic(err)
	}
	return &Otto{
		vm: vm,
	}
}

func (o *Otto) Compile(algebrite []byte) error {
	program, err := o.vm.Compile("", algebrite)
	if err != nil {
		return err
	}
	o.program = program
	return nil
}

func (o *Otto) Load() error {
	_, err := o.vm.Run(o.program)
	if err != nil {
		return err
	}
	return nil
}

func (o *Otto) Run(line string) (string, error) {
	result, err := o.vm.Call("window.Algebrite.run", nil, line)
	if err != nil {
		return "", err
	}
	return result.String(), nil
}
