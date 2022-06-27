package main

import "fmt"

type SDKInterface interface {
	Log(msg string) error
	Args() map[string][]byte
}

type Sdk struct {
	name string
	args map[string][]byte
}

func (s *Sdk) Log(msg string) error {
	fmt.Printf("[SDK-%s]: %s\n", s.name, msg)
	return nil
}

func (s *Sdk) Args() map[string][]byte {
	return s.args
}

var rwSet = make(map[string][]byte)
