package main

import "refl-02/protogo"

type contract interface {
	Init() protogo.Response
	Upgrade() protogo.Response
}

type BaseContract struct {
	Sdk SDKInterface
}

func (bc *BaseContract) Init() protogo.Response {
	return protogo.Response{
		Status:  protogo.Status_OK,
		Message: "Success",
		Payload: []byte("BaseContract Init Success"),
	}
}

func (bc *BaseContract) Upgrade() protogo.Response {
	return protogo.Response{
		Status:  protogo.Status_OK,
		Message: "Success",
		Payload: []byte("BaseContract Upgrade Success"),
	}
}
