package main

import (
	"fmt"
	"refl-02/protogo"
	"reflect"
)

type handler struct {
	sdkOfBaseContract  reflect.Value
	baseContractOfUser reflect.Value
	userContractMethod map[string]reflect.Value
}

func (h *handler) parsUserContract(c contract) error {
	// 1. 检查是否有 BaseContract字段
	contractType := reflect.TypeOf(c).Elem()
	//fmt.Printf("===field===\n")
	fieldNumOfType := contractType.NumField()
	var extendBaseContract bool
	if fieldNumOfType <= 0 {
		panic("UserContract must extend BaseContract")
	}

	//fmt.Printf("field num: %d\n", fieldNumOfType)
	for i := 0; i < fieldNumOfType; i++ {
		fieldI := contractType.Field(i)
		if fieldI.Name == "BaseContract" {
			extendBaseContract = true
		}
		//fmt.Printf("fieldI: %s\n", fieldI.Name)
	}

	if !extendBaseContract {
		panic("UserContract must extend BaseContract")
	}

	if extendBaseContract && fieldNumOfType > 1 {
		panic("UserContract cannot contain fields other than BaseContract")
	}

	// 2. 检查该字段是否为空
	//		2.1 为空设置上BaseContract值
	//		2.2 不为空continue
	contractValue := reflect.ValueOf(c).Elem()
	baseContractOfUser := contractValue.Field(0)
	h.baseContractOfUser = contractValue

	// 3. 设置SDK实例
	sdkOfBaseContract := baseContractOfUser.Field(0)
	h.sdkOfBaseContract = sdkOfBaseContract

	// 4. 解析可导出函数
	methodNum := contractValue.Addr().NumMethod()
	contractTypePtr := reflect.PtrTo(contractType)

	var method reflect.Value
	var name string
	for i := 0; i < methodNum; i++ {
		method = contractValue.Addr().Method(i)
		name = contractTypePtr.Method(i).Name
		// 检查参数列表和返回值列表
		err := h.check(contractTypePtr.Method(i))
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		h.userContractMethod[name] = method
	}

	return nil
}

func (h *handler) SetSdk1() {
	h.sdkOfBaseContract.Set(reflect.ValueOf(&Sdk{name: "sdk-01", args: map[string][]byte{"msg": []byte("value of sdk-01")}}))
}

func (h *handler) SetSdk(sdk SDKInterface) {
	h.sdkOfBaseContract.Set(reflect.ValueOf(sdk))
}

func (h *handler) SetSdk2() {
	h.sdkOfBaseContract.Set(reflect.ValueOf(&Sdk{name: "sdk-02", args: map[string][]byte{"msg": []byte("value of sdk-01")}}))
}

func (h *handler) call(method string) protogo.Response {
	results := h.userContractMethod[method].Call([]reflect.Value{})
	return results[0].Interface().(protogo.Response)
}

func (h *handler) check(method reflect.Method) error {
	if method.Type.NumIn() > 1 {
		return fmt.Errorf("invalid contract method: %s, "+
			"contract method parameter list must be empty", method.Name)
	}

	if method.Type.NumOut() == 0 {
		return fmt.Errorf("invalid contract method: %s, "+
			"contract method must have a return value and must be of response type", method.Name)

	}

	if method.Type.NumOut() > 1 {
		return fmt.Errorf("invalid contract method: %s, "+
			"contract methods can only have one return value and must be of response type", method.Name)
	}

	out := method.Func.Type().Out(0)
	tt := reflect.TypeOf(protogo.Response{})
	if out != tt {
		return fmt.Errorf("invalid contract method: %s, "+
			"contract method return value must be response type", method.Name)
	}

	return nil
}
