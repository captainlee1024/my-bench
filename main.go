package main

import (
	"fmt"
	"reflect"
)

const (
	Method_Print   = "Print"
	Method_Save    = "Save"
	Method_Find    = "Find"
	Method_Init    = "Init"
	Method_Upgrade = "Upgrade"
)

func main() {
	Test2()
}

func Test2() {

	c := &SmartContract{}
	h := &handler{
		userContractMethod: make(map[string]reflect.Value),
	}
	err := h.parsUserContract(c)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("\n======================s1==========================\n")
	s1 := &Sdk{
		name: "s1",
		args: map[string][]byte{
			"msg": []byte("msg from s1"),
		},
	}
	h.SetSdk(s1)
	h.call(Method_Print)
	//
	//fmt.Printf("\n======================s2==========================\n")
	//s2 := &Sdk{
	//	name: "s2",
	//	args: map[string][]byte{
	//		"msg":   []byte("msg from s2"),
	//		"key":   []byte("s2#key01#field01"),
	//		"value": []byte("value of s2"),
	//	},
	//}
	//h.SetSdk(s2)
	//h.call(Method_Print)
	//h.call(Method_Save)
	//
	//response := h.call(Method_Find)
	//if response.Status != protogo.Status_OK {
	//	fmt.Printf("call find failed, err: %s\n", response.Message)
	//} else {
	//	fmt.Printf("find: %s\n", string(response.Payload))
	//
	//}
	//
	//fmt.Printf("\n======================s3==========================\n")
	//s3 := &Sdk{
	//	name: "s3",
	//	args: map[string][]byte{
	//		"msg": []byte("msg from s3"),
	//		"key": []byte("s3#key01#field01######"),
	//	},
	//}
	//h.SetSdk(s3)
	//h.call(Method_Print)
	//
	//response = h.call(Method_Find)
	//if response.Status != protogo.Status_OK {
	//	fmt.Printf("call find failed, err: %s\n", response.Message)
	//} else {
	//	fmt.Printf("find: %s\n", string(response.Payload))
	//
	//}
	//
	//response = h.call(Method_Init)
	//fmt.Printf("\n===Init Response:\nStatus: %d\nMessage: %s\nPayload: %s\n",
	//	response.Status,
	//	response.Message,
	//	response.Payload,
	//)
	//response = h.call(Method_Upgrade)
	//fmt.Printf("\n===Upgrade Response:\nStatus: %d\nMessage: %s\nPayload: %s\n",
	//	response.Status,
	//	response.Message,
	//	response.Payload,
	//)
}

//results := h.call(Method_Find)
//resultErr := results[1].Interface()
//if resultErr != nil {
//	err := resultErr.(error)
//	fmt.Printf("call find failed, err: %s\n", err.Error())
//} else {
//	successResult := results[0].Interface()
//	value := successResult.([]byte)
//	fmt.Printf("find: %s\n", string(value))
//}
//results = h.call(Method_Find)
//resultErr = results[1].Interface()
//if resultErr != nil {
//	err := resultErr.(error)
//	fmt.Printf("call find failed, err: %s\n", err.Error())
//} else {
//	successResult := results[0].Interface()
//	value := successResult.([]byte)
//	fmt.Printf("find: %s\n", string(value))
//}

func test1() {

	var c contract
	c = &SmartContract{}

	// TypeOf
	// TypeOf 返回表示 i 的动态类型的反射Type 。如果 i 是一个 nil 接口值，TypeOf 返回 nil。
	contractType := reflect.TypeOf(c).Elem()
	fmt.Printf("Type: %+v\n", contractType)
	fmt.Printf("===field===\n")
	fieldNumOfType := contractType.NumField()
	for i := 0; i < fieldNumOfType; i++ {
		fieldI := contractType.Field(i)
		fmt.Printf("fieldI: %s\n", fieldI.Name)
	}

	contractType = reflect.PtrTo(reflect.TypeOf(c).Elem())
	fmt.Printf("===method===\n")
	methodNumOfType := contractType.NumMethod()
	for i := 0; i < methodNumOfType; i++ {
		methodI := contractType.Method(i)
		fmt.Printf("fieldI: %s\n", methodI.Name)
	}
	methodSave, ok := contractType.MethodByName("Save")
	if !ok {
		panic("save not found")
	}
	fmt.Printf("type: %s, kind: %s, method: %s, index: %d\n",
		methodSave.Type,
		methodSave.Type.Kind(),
		methodSave.Name,
		methodSave.Index,
	)

	returnsOfSave := methodSave.Func.Call([]reflect.Value{reflect.ValueOf(c), reflect.ValueOf("key-01"), reflect.ValueOf("value-01")})
	fmt.Printf("\nreturus of save: %+v\n", returnsOfSave)
	errResult := returnsOfSave[0].Interface()

	if errResult != nil {
		err := errResult.(error)
		fmt.Printf("save failed, err: %s\n", err.Error())
	}

	// ValueOf
	// ValueOf 返回一个新值，初始化为存储在接口 i 中的具体值。 ValueOf(nil) 返回零值。
	contractValue := reflect.ValueOf(c).Elem()
	fmt.Printf("\n\nValue: %+v\n", contractValue)
	fmt.Printf("===field===\n")
	fieldNumOfValue := contractValue.NumField()
	for i := 0; i < fieldNumOfValue; i++ {
		fieldI := contractValue.Field(i)
		fmt.Printf("fieldI: %s\n", fieldI)
	}
	fmt.Printf("\nbaseContract: %+v\n", contractValue.Field(0))
	contractValue.Field(0).Set(reflect.ValueOf(BaseContract{}))
	fmt.Printf("baseContract: %+v\n\n", contractValue.Field(0))

	contractValue = reflect.ValueOf(c).Elem().Addr()
	fmt.Printf("===method===\n")
	methodNumOfValue := contractValue.NumMethod()
	for i := 0; i < methodNumOfValue; i++ {
		methodI := contractValue.Method(i)
		fmt.Printf("fieldI: %s\n", methodI.Type())
	}
	methodFind := contractValue.Method(0)
	returnsOfFind := methodFind.Call([]reflect.Value{reflect.ValueOf("key-01")})

	fmt.Printf("\nreturus: %+v\n", returnsOfFind)
	errResult = returnsOfFind[1].Interface()
	if errResult != nil {
		err := errResult.(error)
		fmt.Printf("find failed, err: %s\n\n\n", err.Error())
	} else {
		successResult := returnsOfFind[0].Interface()
		value := successResult.(string)
		fmt.Printf("find: %s\n\n\n", value)
	}

	//returnsOfFind = methodFind.Call([]reflect.Value{reflect.ValueOf("key-02")})
	//
	//fmt.Printf("returus: %+v\n", returnsOfFind)
	//errResult = returnsOfFind[1].Interface()
	//if errResult != nil {
	//	err := errResult.(error)
	//	fmt.Printf("find failed, err: %s\n\n\n", err.Error())
	//} else {
	//	successResult := returnsOfFind[0].Interface()
	//	value := successResult.(string)
	//	fmt.Printf("find: %s\n\n\n", value)
	//
}
