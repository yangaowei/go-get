package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

// func (user User) HK() int {
// 	return 10
// }

func (u User) Hello(m string) (int, string) {
	fmt.Println("Hello", m, ", I'm ", m)
	return u.Age + u.Age, m
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())
	fmt.Println("Type:", t.Kind())
	v := reflect.ValueOf(o) //获取接口的值类型
	fmt.Println("Fields:", v)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s:%v=%v\n", f.Name, f.Type, val)
	}

	fmt.Println(">>>>>>")
	args := []reflect.Value{reflect.ValueOf("yan")}
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		//fmt.Println(m.Call())
		fmt.Printf("%6s: %v\n", m.Name, m.Type) //获取方法的名称和类型
		mv := v.MethodByName(m.Name)
		fmt.Println(mv.Call(args)[0])
	}
	fmt.Println(">>>>>>")

	//v = v.Elem()
	e := reflect.ValueOf(&o).Elem()
	fmt.Println(e.CanSet())
	fmt.Println(v.Kind() == reflect.Ptr)
}

func main() {
	var user User
	user.Id = 10
	Info(user)
	fmt.Println("............")
	//Info(0)

}
