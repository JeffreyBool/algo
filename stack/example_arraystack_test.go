/**
  * Author: JeffreyBool
  * Date: 2019/4/22
  * Time: 19:03
  * Software: GoLand
*/

package stack_test

import (
	"algo/stack"
	"fmt"
	"log"
)

//实例化
func ExampleNewArrayStack() {
	arrayStack := stack.NewArrayStack()
	arrayStack.PrintIn()
	//Output:
	//Stack: Array: size = 0 , capacity = 10
	//[] top
}

//获取栈长度
func ExampleArrayStack_GetSize() {
	stack := stack.NewArrayStack()
	fmt.Printf("stack len: %d \n",stack.GetSize())
	//Output:
	//stack len: 0
}

//判断栈是否为空
func ExampleArrayStack_IsEmpty() {
	stack := stack.NewArrayStack()
	fmt.Printf("stack empty: %t \n",stack.IsEmpty())
	//Output:
	//stack empty: true
}

//压入栈顶元素
func ExampleArrayStack_Push() {
	stack := stack.NewArrayStack()
	if err := stack.Push("你好呀？"); err != nil{
		log.Fatal(err)
	}

	stack.PrintIn()
	//Output:
	//Stack: Array: size = 1 , capacity = 10
	//[你好呀？] top
}

//弹出栈顶元素
func ExampleArrayStack_Pop() {
	stack := stack.NewArrayStack()
	if err := stack.Push("你好呀？"); err != nil{
		log.Fatal(err)
	}
	if err := stack.Push("hello"); err != nil{
		log.Fatal(err)
	}
	stack.PrintIn()
	if pop, err := stack.Pop(); err != nil{
		log.Fatal(err)
	}else{
		fmt.Println(pop)
	}
	//Output:
	//Stack: Array: size = 2 , capacity = 10
	//[你好呀？, hello] top
	//hello
}

//查看栈顶元素
func ExampleArrayStack_Peek() {
	stack := stack.NewArrayStack()
	strs := []string{"A","B","C","D","E","F"}
	for _,str := range strs{
		stack.Push(str)
	}
	peek := stack.Peek()
	fmt.Printf("stack peek: %+v \n",peek)
	stack.PrintIn()
	//Output:
	//stack peek: F
	//Stack: Array: size = 6 , capacity = 10
	//[A, B, C, D, E, F] top
}

//格式化输出
func ExampleArrayStack_PrintIn() {
	stack := stack.NewArrayStack()
	strs := []string{"A","B","C","D","E","F"}
	for _,str := range strs{
		stack.Push(str)
	}

	stack.PrintIn()
	//Output:
	//Stack: Array: size = 6 , capacity = 10
	//[A, B, C, D, E, F] top
}
