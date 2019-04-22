/**
  * Author: JeffreyBool
  * Date: 2019/4/22
  * Time: 12:29
  * Software: GoLand
*/

package stack

import (
	"algo/array"
	"fmt"
	"bytes"
)

type ArrayStack struct {
	array *array.Array
}

//初始化栈
func NewArrayStack(capacity ...int) (stack *ArrayStack) {
	if len(capacity) >= 1 && capacity[0] != 0 {
		stack = &ArrayStack{
			array:array.NewArray(capacity[0]),
		}
	} else {
		stack = &ArrayStack{
			array:array.NewArray(10),
		}
	}

	return
}

//获取栈元素长度
func (stack *ArrayStack) GetSize() int {
	return stack.array.GetSize()
}

//判断栈是否为空
func (stack *ArrayStack) IsEmpty() bool {
	return stack.array.IsEmpty()
}

//压入栈顶元素
func (stack *ArrayStack) Push(v interface{}) error {
	return stack.array.AddLast(v)
}

//弹出栈顶元素
func (stack *ArrayStack) Pop() (interface{}, error) {
	return stack.array.RemoveLast()
}

//查看栈顶元素
func (stack *ArrayStack) Peek() interface{} {
	last, _ := stack.array.GetLast()
	return last
}

//格式化输出
func (stack *ArrayStack) PrintIn()  {
	var buffer bytes.Buffer
	buffer.WriteString("Stack: ")
	format := fmt.Sprintf("Array: size = %d , capacity = %d\n",stack.array.GetSize(), stack.array.GetCapacity())
	buffer.WriteString(format)
	buffer.WriteString("[")
	for i := 0; i < stack.array.GetSize(); i++ {
		value, _ := stack.array.Get(i)
		buffer.WriteString(fmt.Sprint(value))
		if i != stack.array.GetSize()-1 {
			buffer.WriteString(", ")
		}
	}

	buffer.WriteString("] top")
	fmt.Println(buffer.String())
}
