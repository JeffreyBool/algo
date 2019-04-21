/**
  * Author: JeffreyBool
  * Date: 2019/4/12
  * Time: 11:36
  * Software: GoLand
*/

package array_test

import (
	"testing"
	"algo/array"
)

//测试实例化
func TestNewArray(t *testing.T) {
	array := array.NewArray(2)
	for i := 0; i < 10; i++ {
		if err := array.Add(i,i+1); err != nil {
			t.Error(err)
			break
		}
	}

	array.PrintIn()
	if err := array.Add(10, 11); err != nil {
		t.Error(err)
	}

	if err := array.Add(10, 12); err != nil {
		t.Error(err)
	}

	array.PrintIn()
}

//获取数组容量
func TestArray_GetCapacity(t *testing.T) {
	array := array.NewArray()
	array.AddLast("我是张高元")
	array.PrintIn()
	t.Logf("array cap:%d \n",array.GetCapacity())
}

//获取数组长度
func TestArray_GetSize(t *testing.T) {
	array := array.NewArray()
	t.Logf("array len: %d \n", array.GetSize())
}

//判断是否为空
func TestArray_IsEmpty(t *testing.T) {
	array := array.NewArray()
	t.Logf("array empty: %t \n", array.IsEmpty())
}

//向数组头部添加元素
func TestArray_AddFirst(t *testing.T) {
	array := array.NewArray()
	array.AddFirst("array add first")
	array.AddLast("array add last")
	array.Add(1,"array add value")
	array.PrintIn()
}

//向数组末尾添加元素
func TestArray_AddLast(t *testing.T) {
	array := array.NewArray()
	array.Add(0,19)
	array.Add(1,19)
	array.Add(2,19)
	array.Add(3,19)
	array.Add(4,19)
	array.AddLast(20)
	array.PrintIn()
}

//动态添加元素
func TestArray_Add(t *testing.T) {
	strs := []string{"A","B","C","D","E","F"}
	array := array.NewArray()
	for i:= 0; i < len(strs); i++{
		if err := array.Add(i,strs[i]); err != nil{
			t.Error(err)
			break
		}
	}
	array.PrintIn()
}

//动态添加元素性能测试
func BenchmarkArray_Add(b *testing.B) {
	array := array.NewArray(b.N)

	//重置时间点
	b.ResetTimer()
	for i:= 0; i < b.N; i++{
		array.Add(i,i+1)
	}

	array.PrintIn()
}

//根据索引获取某个值
func TestArray_Get(t *testing.T) {
	array := array.NewArray(10)
	array.Add(0,10)
	if value,err := array.Get(0);err != nil{
		t.Error(err)
	}else{
		t.Log(value)
	}
}

//根据索引修改某个值
func TestArray_Set(t *testing.T) {
	array := array.NewArray()
	array.AddFirst("array add first")
	array.PrintIn()
	array.Set(0,"array set value")
	array.PrintIn()
}

//判断数组是否存在某个值
func TestArray_Contains(t *testing.T) {
	array := array.NewArray()
	if array.Contains("我是张三"){
		t.Log("找到了")
	}
	array.AddFirst("我是张三")
	if array.Contains("我是张三"){
		t.Log("找到了")
	}
	array.PrintIn()
}

//查询一个值的索引位置
func TestArray_Find(t *testing.T) {
	array := array.NewArray(10)
	array.AddFirst("我是张三")
	array.AddFirst("我是张三")
	array.AddLast("我是李四")
	find := array.Find("我是李四")
	t.Log(find)
	array.PrintIn()
}

//根据索引删除元素
func TestArray_Remove(t *testing.T) {
	array := array.NewArray()
	strs := []string{"A","B","C","D","E","F"}
	for index, str := range strs{
		array.Add(index,str)
	}

	array.PrintIn()
	array.Remove(2)
	array.PrintIn()
}

//删除头部元素
func TestArray_RemoveFirst(t *testing.T) {
	array := array.NewArray()
	strs := []string{"A","B","C","D","E","F"}
	for index, str := range strs{
		array.Add(index,str)
	}

	array.PrintIn()
	array.RemoveFirst()
	array.PrintIn()
}

//删除末尾元素
func TestArray_RemoveLast(t *testing.T) {
	array := array.NewArray()
	strs := []string{"A","B","C","D","E","F"}
	for index, str := range strs{
		array.Add(index,str)
	}

	array.PrintIn()
	array.RemoveLast()
	array.PrintIn()
}

//删除指定元素
func TestArray_RemoveElement(t *testing.T) {
	array := array.NewArray()
	strs := []string{"A","B","C","D","E","F"}
	for index, str := range strs{
		array.Add(index,str)
	}

	array.PrintIn()
	array.RemoveElement("B")
	array.PrintIn()
}

//清空元素
func TestArray_Clear(t *testing.T) {
	array := array.NewArray()
	strs := []string{"A","B","C","D","E","F"}
	for index, str := range strs{
		array.Add(index,str)
	}

	array.PrintIn()
	array.Clear()
	array.PrintIn()
}

//打印输出
func TestArray_PrintIn(t *testing.T) {
	array := array.NewArray()
	strs := []string{"A","B","C","D","E","F"}
	for index, str := range strs{
		array.Add(index,str)
	}
	array.PrintIn()
}

