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

func TestNewArray(t *testing.T) {
	array := array.NewArray(2)
	//t.Logf("array cap: %d \n",array.GetCapacity())
	//t.Logf("array count: %d \n",array.Count())
	if err := array.Set(0, 2); err != nil {
		t.Error(err)
	}

	//t.Logf("array cap: %d \n",array.GetCapacity())
	//t.Logf("array count: %d \n",array.Count())
	//t.Log(array)

	for i := 0; i < 10; i++ {
		if err := array.Add(uint(i), i+1); err != nil {
			t.Error(err)
			break
		}
	}

	if err := array.Add(10, 11); err != nil {
		t.Error(err)
	}
	t.Log(array)

	if err := array.Add(10, 12); err != nil {
		t.Error(err)
	}

	t.Logf("test array cap: %d \n", array.GetCapacity())
	t.Logf("test array count: %d \n", array.Count())
	t.Log(array)
	array.Print()
}

//获取数组容量
func TestArray_GetCapacity(t *testing.T) {
	array := array.NewArray(10)
	t.Logf("cap: %d \n", array.GetCapacity())
}

//获取数组长度
func TestArray_Count(t *testing.T) {
	array := array.NewArray(10)
	t.Logf("len: %d \n", array.Count())
}

//判断是否为空
func TestArray_IsEmpty(t *testing.T) {
	array := array.NewArray(10)
	t.Logf("empty: %t \n", array.IsEmpty())
}

//设置某个索引的值
func TestArray_Set(t *testing.T) {
	array := array.NewArray(10)
	if err := array.Set(0,10); err != nil{
		t.Error(err)
	}

	t.Logf("test array cap: %d \n", array.GetCapacity())
	t.Logf("test array count: %d \n", array.Count())
	array.Print()
}

//根据索引获取某个值
func TestArray_Get(t *testing.T) {
	array := array.NewArray(10)
	array.Set(0,10)
	if err, value := array.Get(0);err != nil{
		t.Error(err)
	}else{
		t.Log(value)
	}
}

//查找一个值是否存在
func TestArray_Find(t *testing.T) {
	array := array.NewArray(10)
	find := array.Find(0)
	t.Log(find)
}

//指定位置添加元素
func BenchmarkArray_Add(b *testing.B) {
	array := array.NewArray(10)

	//重置时间点
	b.ResetTimer()
	for i:= 0; i < b.N; i++{
		array.Add(uint(i),i+1)
	}

	b.Logf("test array cap: %d \n", array.GetCapacity())
	b.Logf("test array count: %d \n", array.Count())
	array.Print()
}

//向数组头部添加元素
func TestArray_AddFirst(t *testing.T) {
	array := array.NewArray(10)
	array.AddFirst(1)
	t.Logf("test array cap: %d \n", array.GetCapacity())
	t.Logf("test array count: %d \n", array.Count())
	array.Print()
}

//向数组末尾添加元素
func TestArray_AddLast(t *testing.T) {
	array := array.NewArray(10)
	array.Add(0,19)
	array.Add(1,19)
	array.Add(2,19)
	array.Add(3,19)
	array.Add(4,19)
	array.AddLast(20)
	t.Logf("test array cap: %d \n", array.GetCapacity())
	t.Logf("test array count: %d \n", array.Count())
	array.Print()
	t.Log(array)
}
