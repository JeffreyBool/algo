/**
  * Author: JeffreyBool
  * Date: 2019/4/21
  * Time: 13:45
  * Software: GoLand
*/

package array_test

import (
	"fmt"
	"algo/array"
	"log"
)

//实例化
func ExampleNewArray() {
	arr := array.NewArray(2)
	fmt.Println(arr.GetSize())
	fmt.Println(arr.GetCapacity())
	//Output:
	//0
	//2
}

//获取数组容量
func ExampleGetCapacity() {
	array := array.NewArray(10)
	fmt.Printf("cap: %d \n", array.GetCapacity())
}

//获取数组长度
func ExampleArray_GetSize() {
	array := array.NewArray(10)
	fmt.Printf("len: %d \n", array.GetSize())
}

//判断是否为空
func ExampleArray_IsEmpty() {
	array := array.NewArray(10)
	fmt.Printf("empty: %t \n", array.IsEmpty())
}

//设置某个索引的值
func ExampleArray_Set() {
	array := array.NewArray(10)
	if err := array.Set(0,10); err != nil{
		log.Fatal(err)
	}

	fmt.Printf("test array cap: %d \n", array.GetCapacity())
	fmt.Printf("test array count: %d \n", array.GetSize())
	array.PrintIn()
}

//根据索引获取某个值
func ExampleArray_Get() {
	array := array.NewArray(10)
	array.Set(0,10)
	if err, value := array.Get(0);err != nil{
		log.Fatal(err)
	}else{
		fmt.Println(value)
	}
}

//查找一个值是否存在
func ExampleArray_Find() {
	array := array.NewArray(10)
	find := array.Find(0)
	fmt.Println(find)
}

//指定位置添加元素
func  ExampleArray_Add() {
	array := array.NewArray(10)

	for i:= 0; i < 10; i++{
		array.Add(i,i+1)
	}

	fmt.Printf("test array cap: %d \n", array.GetCapacity())
	fmt.Printf("test array count: %d \n", array.GetSize())
	array.PrintIn()
}

//向数组头部添加元素
func ExampleArray_AddFirst() {
	array := array.NewArray(10)
	array.AddFirst(1)
	fmt.Printf("test array cap: %d \n", array.GetCapacity())
	fmt.Printf("test array count: %d \n", array.GetSize())
	array.PrintIn()
}

//向数组末尾添加元素
func ExampleArray_AddLast() {
	array := array.NewArray(10)
	array.Add(0,19)
	array.Add(1,19)
	array.Add(2,19)
	array.Add(3,19)
	array.Add(4,19)
	array.AddLast(20)
	fmt.Printf("test array cap: %d \n", array.GetCapacity())
	fmt.Printf("test array count: %d \n", array.GetSize())
	array.PrintIn()
	fmt.Println(array)
}

