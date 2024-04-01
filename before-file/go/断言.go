package main

import (
	"fmt"
	"strconv"
)

// 前提
// 断言是接口类型才有的
// 非接口类型使用断言会报错

// Element 定义空接口类型 Element
type Element interface{}

// List 定义接口类型切片 List
type List []Element

// 定义结构体 Person
type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func main() {
	list := make([]Element, 3, 6)

	list[0] = 111
	list[1] = "GG"
	list[2] = Person{"Polar", 27}

	for i, Element := range list {

		if v, ok := Element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d \n", i, v)
		} else if v, ok := Element.(string); ok {
			fmt.Printf("list[%d] is an string and its value is %s \n", i, v)
		} else if v, ok := Element.(Person); ok {
			fmt.Printf("list[%d] is an string and its value is %s \n", i, v)
		} else {
			fmt.Printf("list[%d] is of a different type \n", i)
		}
	}

	// 使用element.(type) 语法断言元素类型
	// 注意：element.(type) 不能再 switch 外的任何逻辑中使用
	// 如果要在switch 外面判断一个类型就要使用 comma-ok
	for i, Element := range list {

		switch v := Element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d \n", i, v)
		case string:
			fmt.Printf("list[%d] is an string and its value is %s \n", i, v)
		case Person:
			fmt.Printf("list[%d] is an string and its value is %s \n", i, v)
		default:
			fmt.Printf("list[%d] is of a different type \n", i)
		}
	}
}
