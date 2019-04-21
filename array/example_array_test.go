/**
  * Author: JeffreyBool
  * Date: 2019/4/21
  * Time: 13:45
  * Software: GoLand
*/

package array_test

import (
	"algo/array"
	"fmt"
	"log"
)

/**
* 结构体和方法实现:
type Array struct {
	// Has unexported fields.
}

func NewArray(capacity ...int) (array *Array)
func (array *Array) Add(index int, value interface{}) (err error)
func (array *Array) AddFirst(value interface{}) error
func (array *Array) AddLast(value interface{}) error
func (array *Array) Clear()
func (array *Array) Contains(value interface{}) bool
func (array *Array) Find(value interface{}) int
func (array *Array) Get(index int) (value interface{}, err error)
func (array *Array) GetCapacity() int
func (array *Array) GetSize() int
func (array *Array) IsEmpty() bool
func (array *Array) PrintIn()
func (array *Array) Remove(index int) (value interface{}, err error)
func (array *Array) RemoveElement(value interface{}) (e interface{}, err error)
func (array *Array) RemoveFirst() (interface{}, error)
func (array *Array) RemoveLast() (interface{}, error)
func (array *Array) Set(index int, value interface{}) (err error)
*/

/*
 * 启动文档本地服务命令
 * 进入目录: cd $GOPATH/algo/array
 * 启动文档服务: godoc -http=:6060 -play
 * 查看: http://127.0.0.1:6060/pkg/algo/array/
*/

//测试实例化
func ExampleNewArray() {
	array := array.NewArray(2)
	for i := 0; i < 10; i++ {
		if err := array.Add(i, i+1); err != nil {
			log.Fatal(err)
			break
		}
	}

	array.PrintIn()
	if err := array.Add(10, 11); err != nil {
		log.Fatal(err)
	}

	if err := array.Add(10, 12); err != nil {
		log.Fatal(err)
	}

	array.PrintIn()

	//Output:
	//Array: size = 10 , capacity = 16
	//[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
	//Array: size = 12 , capacity = 16
	//[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 11]
}

//获取数组容量
func ExampleArray_GetCapacity() {
	array := array.NewArray()
	array.AddLast("我是张高元")
	array.PrintIn()
	fmt.Printf("array cap:%d \n", array.GetCapacity())
	//Output:
	//Array: size = 1 , capacity = 10
	//[我是张高元]
	//array cap:10
}

//获取数组长度
func ExampleArray_GetSize() {
	array := array.NewArray()
	fmt.Printf("array len: %d \n", array.GetSize())
	//Output:
	//array len: 0
}

//判断是否为空
func ExampleArray_IsEmpty() {
	array := array.NewArray()
	fmt.Printf("array empty: %t \n", array.IsEmpty())
	//Output:
	//array empty: true
}

//向数组头部添加元素
func ExampleArray_AddFirst() {
	array := array.NewArray()
	array.AddFirst("array add first")
	array.AddLast("array add last")
	array.Add(1, "array add value")
	array.PrintIn()
	//Output:
	//Array: size = 3 , capacity = 10
	//[array add first, array add value, array add last]
}

//向数组末尾添加元素
func ExampleArray_AddLast() {
	array := array.NewArray()
	array.Add(0, 19)
	array.Add(1, 19)
	array.Add(2, 19)
	array.Add(3, 19)
	array.Add(4, 19)
	array.AddLast(20)
	array.PrintIn()
	//Output:
	//Array: size = 6 , capacity = 10
	//[19, 19, 19, 19, 19, 20]
}

//动态添加元素
func ExampleArray_Add() {
	strs := []string{"A", "B", "C", "D", "E", "F"}
	array := array.NewArray()
	for i := 0; i < len(strs); i++ {
		if err := array.Add(i, strs[i]); err != nil {
			log.Fatal(err)
			break
		}
	}
	array.PrintIn()
	//Output:
	//Array: size = 6 , capacity = 10
	//[A, B, C, D, E, F]
}

//根据索引获取某个值
func ExampleArray_Get() {
	array := array.NewArray(10)
	array.Add(0, 10)
	if value, err := array.Get(0); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(value)
	}
	//Output:
	//10
}

//根据索引修改某个值
func ExampleArray_Set() {
	array := array.NewArray()
	array.AddFirst("array add first")
	array.PrintIn()
	array.Set(0, "array set value")
	array.PrintIn()
	//Output:
	//Array: size = 1 , capacity = 10
	//[array add first]
	//Array: size = 1 , capacity = 10
	//[array set value]
}

//判断数组是否存在某个值
func ExampleArray_Contains() {
	array := array.NewArray()
	if array.Contains("我是张三") {
		fmt.Println("找到了")
	}
	array.AddFirst("我是张三")
	if array.Contains("我是张三") {
		fmt.Println("找到了")
	}
	array.PrintIn()
	//Output:
	//找到了
	//Array: size = 1 , capacity = 10
	//[我是张三]
}

//查询一个值的索引位置
func ExampleArray_Find() {
	array := array.NewArray(10)
	array.AddFirst("我是张三")
	array.AddFirst("我是张三")
	array.AddLast("我是李四")
	find := array.Find("我是李四")
	fmt.Println(find)
	array.PrintIn()
	//Output:
	//2
	//Array: size = 3 , capacity = 10
	//[我是张三, 我是张三, 我是李四]
}

//根据索引删除元素
func ExampleArray_Remove() {
	array := array.NewArray()
	strs := []string{"A", "B", "C", "D", "E", "F"}
	for index, str := range strs {
		array.Add(index, str)
	}

	array.PrintIn()
	array.Remove(2)
	array.PrintIn()
	//Output:
	//Array: size = 6 , capacity = 10
	//[A, B, C, D, E, F]
	//Array: size = 5 , capacity = 10
	//[A, B, D, E, F]
}

//删除头部元素
func ExampleArray_RemoveFirst() {
	array := array.NewArray()
	strs := []string{"A", "B", "C", "D", "E", "F"}
	for index, str := range strs {
		array.Add(index, str)
	}

	array.PrintIn()
	array.RemoveFirst()
	array.PrintIn()
	//Output:
	//array: size = 6 , capacity = 10
	//[A, B, C, D, E, F]
	//Array: size = 5 , capacity = 10
	//[B, C, D, E, F]
}

//删除末尾元素
func ExampleArray_RemoveLast() {
	array := array.NewArray()
	strs := []string{"A", "B", "C", "D", "E", "F"}
	for index, str := range strs {
		array.Add(index, str)
	}

	array.PrintIn()
	array.RemoveLast()
	array.PrintIn()
	//Output:
	//Array: size = 6 , capacity = 10
	//[A, B, C, D, E, F]
	//Array: size = 5 , capacity = 10
	//[A, B, C, D, E]
}

//删除指定元素
func ExampleArray_RemoveElement() {
	array := array.NewArray()
	strs := []string{"A", "B", "C", "D", "E", "F"}
	for index, str := range strs {
		array.Add(index, str)
	}

	array.PrintIn()
	array.RemoveElement("B")
	array.PrintIn()
	//Output:
	//Array: size = 6 , capacity = 10
	//[A, B, C, D, E, F]
	//Array: size = 5 , capacity = 10
	//[A, C, D, E, F]
}

//清空元素
func ExampleArray_Clear() {
	array := array.NewArray()
	strs := []string{"A", "B", "C", "D", "E", "F"}
	for index, str := range strs {
		array.Add(index, str)
	}

	array.PrintIn()
	array.Clear()
	array.PrintIn()
	//Output:
	//Array: size = 6 , capacity = 10
	//[A, B, C, D, E, F]
	//Array: size = 0 , capacity = 6
	//[]
}

//打印输出
func ExampleArray_PrintIn() {
	array := array.NewArray()
	strs := []string{"A", "B", "C", "D", "E", "F"}
	for index, str := range strs {
		array.Add(index, str)
	}
	array.PrintIn()
	//Output:
	//Array: size = 6 , capacity = 10
	//[A, B, C, D, E, F]
}
