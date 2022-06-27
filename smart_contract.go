package main

import (
	"fmt"
	"refl-02/protogo"
)

type SmartContract struct {
	BaseContract
}

func (s *SmartContract) Init() protogo.Response {
	return protogo.Response{
		Status:  protogo.Status_OK,
		Message: "Success",
		Payload: []byte("Init Success"),
	}
}

//func (s *SmartContract) Upgrade() *Response {
//	return &Response{
//		Status:  Status_OK,
//		Message: "Success",
//		Payload: []byte("Init Success"),
//	}
//}

func (s *SmartContract) Save() protogo.Response {
	args := s.Sdk.Args()

	key := string(args["key"])
	value := args["value"]
	rwSet[key] = value
	return protogo.Response{
		Status: protogo.Status_OK,
	}
}

func (s *SmartContract) Find() protogo.Response {
	args := s.Sdk.Args()
	key := string(args["key"])

	value, ok := rwSet[key]
	if !ok {
		return protogo.Response{
			Status:  protogo.Status_Fail,
			Message: fmt.Sprintf("not found value for %s", key),
		}
	}

	return protogo.Response{
		Status:  protogo.Status_OK,
		Payload: value,
	}
}

func (s *SmartContract) Print() protogo.Response {
	args := s.Sdk.Args()
	err := s.Sdk.Log(string(args["msg"]))
	if err != nil {
		return protogo.Response{
			Status:  protogo.Status_Fail,
			Message: err.Error(),
		}
	}

	return protogo.Response{
		Status: protogo.Status_OK,
	}
}

//func (s *SmartContract) T01(a string) protogo.Response {
//	return protogo.Response{}
//}
//
//func (s *SmartContract) T02() error {
//	return nil
//}
//
//func (s *SmartContract) T03() (string, error) {
//	return "", nil
//}
//
//func (s *SmartContract) t04() {}
