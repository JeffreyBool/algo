/**
  * Author: JeffreyBool
  * Date: 2019/4/12
  * Time: 11:28
  * Software: GoLand
*/

package array

import (
	"errors"
	"fmt"
)

/**
 * 1) 数组的插入、删除、按照下标随机访问操作；
 * 2）数组中的数据是interface类型的；
 */

type Array struct {
	data []interface{}
	size int
}

//数组初始化内存 (可以指定长度,默认为10个长度)
func NewArray(capacity ...int) (array *Array) {
	if len(capacity) >= 1 && capacity[0] != 0 {
		array = &Array{
			data: make([]interface{}, capacity[0]),
			size: 0,
		}
	} else {
		array = &Array{
			data: make([]interface{}, 10),
			size: 0,
		}
	}

	return
}

//判断索引是否越界
func (array *Array) checkIndex(index int) bool {
	if index < 0 || index >= array.size {
		return true
	}

	return false
}

//数组扩容
func (array *Array) resize(capacity int) {
	newArray := make([]interface{}, capacity)
	for i := 0; i < array.size; i ++ {
		newArray[i] = array.data[i]
	}
	array.data = newArray
	newArray = nil
}

//获取数组容量
func (array *Array) GetCapacity() int {
	return cap(array.data)
}

//获取数组长度
func (array *Array) GetSize() int {
	return array.size
}

//判断数组是否为空
func (array *Array) IsEmpty() bool {
	return array.size == 0
}

//向数组头插入元素
func (array *Array) AddFirst(value interface{}) error {
	return array.Add(0, value)
}

//向数组尾插入元素
func (array *Array) AddLast(value interface{}) error {
	return array.Add(array.size, value)
}

//在 index 位置，插入元素e, 时间复杂度 O(m+n)
func (array *Array) Add(index int, value interface{}) (err error) {
	if index < 0 || index > array.size {
		err = errors.New("Add failed. Require index >= 0 and index <= size.")
		return
	}

	// 如果当前元素个数等于数组容量，则将数组扩容为原来的2倍
	cap := array.GetCapacity()
	if array.size == cap {
		array.resize(cap * 2)
	}

	for i := array.size - 1; i >= index; i-- {
		array.data[i+1] = array.data[i]
	}

	array.data[index] = value
	array.size++
	return
}

//获取对应 index 位置的元素
func (array *Array) Get(index int) (value interface{}, err error) {
	if array.checkIndex(index) {
		err = errors.New("Get failed. Index is illegal.")
		return
	}

	value = array.data[index]
	return
}

//修改 index 位置的元素
func (array *Array) Set(index int, value interface{}) (err error) {
	if array.checkIndex(index) {
		err = errors.New("Set failed. Index is illegal.")
		return
	}

	array.data[index] = value
	return
}

//查找数组中是否有元素
func (array *Array) Contains(value interface{}) bool {
	for i := 0; i < array.size; i++ {
		if array.data[i] == value {
			return true
		}
	}

	return false
}

//通过索引查找数组，索引范围[0,n-1](未找到，返回 -1)
func (array *Array) Find(value interface{}) int {
	for i := 0; i < array.size; i++ {
		if array.data[i] == value {
			return i
		}
	}

	return -1
}

// 删除 index 位置的元素，并返回
func (array *Array) Remove(index int) (value interface{}, err error) {
	if array.checkIndex(index) {
		err = errors.New("Remove failed. Index is illegal.")
		return
	}

	value = array.data[index]
	for i := index + 1; i < array.size; i++ {
		//数据全部往前挪动一位,覆盖需要删除的元素
		array.data[i-1] = array.data[i]
	}

	array.size--
	array.data[array.size] = nil //loitering objects != memory leak

	cap := array.GetCapacity()
	if array.size == cap/4 && cap/2 != 0 {
		array.resize(cap / 2)
	}
	return
}

//删除数组首个元素
func (array *Array) RemoveFirst() (interface{}, error) {
	return array.Remove(0)
}

//删除末尾元素
func (array *Array) RemoveLast() (interface{}, error) {
	return array.Remove(int(array.size - 1))
}

//从数组中删除指定元素
func (array *Array) RemoveElement(value interface{}) (e interface{}, err error) {
	index := array.Find(value)
	if index != -1 {
		e, err = array.Remove(index)
	}
	return
}

//清空数组
func (array *Array) Clear() {
	array.data = make([]interface{}, array.size)
	array.size = 0
}

//打印数列
func (array *Array) PrintIn() {
	var format string
	format = fmt.Sprintf("Array: size = %d , capacity = %d\n",array.size, cap(array.data))
	format += "["
	for i := 0; i < array.GetSize(); i++ {
		format += fmt.Sprintf("%+v", array.data[i])
		if i != array.size -1 {
			format += ", "
		}
	}
	format += "]"
	fmt.Println(format)
}
