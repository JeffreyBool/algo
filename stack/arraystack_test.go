/**
  * Author: JeffreyBool
  * Date: 2019/4/22
  * Time: 18:13
  * Software: GoLand
*/

package stack_test

import (
	"testing"
	"algo/stack"
)

//实例化
func TestNewArrayStack(t *testing.T) {
	stack := stack.NewArrayStack()
	t.Log(stack)
}

//获取栈长度
func TestArrayStack_GetSize(t *testing.T) {
	stack := stack.NewArrayStack()
	t.Logf("stack len: %d \n",stack.GetSize())
}

//判断栈是否为空
func TestArrayStack_IsEmpty(t *testing.T) {
	stack := stack.NewArrayStack()
	t.Logf("stack empty: %t \n",stack.IsEmpty())
}

//压入栈顶元素
func TestArrayStack_Push(t *testing.T) {
	stack := stack.NewArrayStack()
	if err := stack.Push("你好呀？"); err != nil{
		t.Error(err)
	}

	stack.PrintIn()
}

//弹出栈顶元素
func TestArrayStack_Pop(t *testing.T) {
	stack := stack.NewArrayStack()
	if err := stack.Push("你好呀？"); err != nil{
		t.Error(err)
	}
	if err := stack.Push("hello"); err != nil{
		t.Error(err)
	}
	stack.PrintIn()
	if pop, err := stack.Pop(); err != nil{
		t.Error(err)
	}else{
		t.Log(pop)
	}
}

//查看栈顶元素
func TestArrayStack_Peek(t *testing.T) {
	stack := stack.NewArrayStack()
	strs := []string{"A","B","C","D","E","F"}
	for _,str := range strs{
		stack.Push(str)
	}
	peek := stack.Peek()
	t.Logf("stack peek: %+v \n",peek)
	stack.PrintIn()
}

//格式化输出
func TestArrayStack_PrintIn(t *testing.T) {
	stack := stack.NewArrayStack()
	strs := []string{"A","B","C","D","E","F"}
	for _,str := range strs{
		stack.Push(str)
	}

	stack.PrintIn()
}

