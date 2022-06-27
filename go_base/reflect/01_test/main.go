package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id       int
	Username string
	Password string
}

func (u User) Calling() {
	fmt.Printf("%s is calling...", u.Username)
}

func DoFiledAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)
	fmt.Printf("type:%s, value:%v\n", inputType.Name(), inputValue)
	fieldNum := inputValue.NumField()
	for i := 0; i < fieldNum; i++ {
		fieldType := inputType.Field(i)
		fieldValue := inputValue.Field(i).Interface()
		fmt.Printf("Name:%s, Type:%v, Value:%v\n", fieldType.Name, fieldType.Type, fieldValue)
	}
	methodNum := inputType.NumMethod()
	fmt.Println(methodNum)
	for i := 0; i < methodNum; i++ {
		method := inputType.Method(i)

		fmt.Println(method)
	}
}

func main() {
	u := User{Id: 12, Username: "zhang3", Password: "123"}
	DoFiledAndMethod(u)
}
