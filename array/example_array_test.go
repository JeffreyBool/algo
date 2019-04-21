/**
  * Author: JeffreyBool
  * Date: 2019/4/21
  * Time: 13:45
  * Software: GoLand
*/

package array_test

import (
	"algo/array"
	"log"
	"fmt"
)

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
func ExampleArrayGetCapacity() {
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
func ExampleArrayGetSize() {
	array := array.NewArray()
	fmt.Printf("array len: %d \n", array.GetSize())
	//Output:
	//array len: 0
}

//判断是否为空
func ExampleArrayIsEmpty() {
	array := array.NewArray()
	fmt.Printf("array empty: %t \n", array.IsEmpty())
	//Output:
	//array empty: true
}

//向数组头部添加元素
func ExampleArrayAddFirst() {
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
func ExampleArrayAddLast() {
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
func ExampleArrayAdd() {
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
func ExampleArrayGet() {
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
func ExampleArraySet() {
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
func ExampleArrayContains() {
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
func ExampleArrayFind() {
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
func ExampleArrayRemove() {
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
func ExampleArrayRemoveFirst() {
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
func ExampleArrayRemoveLast() {
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
func ExampleArrayRemoveElement() {
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
func ExampleArrayClear() {
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
func ExampleArrayPrintIn() {
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
