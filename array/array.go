/**
  * Author: JeffreyBool
  * Date: 2019/4/12
  * Time: 11:28
  * Software: GoLand
*/

package array

import (
	"errors"
)

/**
 * 1) 数组的插入、删除、按照下标随机访问操作；
 * 2）数组中的数据是int类型的；
 */

type Array struct {
	data   []int
	length int
}

//数组初始化内存 (可以指定长度,默认为50个长度)
func NewArray(capacity ...uint) (array *Array) {
	if len(capacity) >= 1 && capacity[0] != 0 {
		array = &Array{
			data:   make([]int,capacity[0]),
			length: 0,
		}
	} else {
		array = &Array{
			data:   make([]int,50),
			length: 0,
		}
	}

	return
}

//判断索引是否越界
func (array *Array) checkIndex(index uint) (err error) {
	/*
	**判断当前索引是否大于 数组的分配空间大小
	**cap()函数返回的是数组切片分配的空间大小
	*/
	if index < 0 || int(index) > array.length {
		err = errors.New("Add failed! Require index >=0 and index <= size.")
	}

	return err
}

func (array *Array) resize(capacity uint) {
	newArray := NewArray(capacity)
	for i := 0; i < array.length; i ++ {
		newArray.data[i] = array.data[i]
	}

	array.data = make([]int,capacity)
	array.data = newArray.data
}

//获取数组容量
func (array *Array) GetCapacity() int {
	return cap(array.data)
}

//获取数组长度
func (array *Array) Count() int {
	return array.length
}

//判断数组是否为空
func (array *Array) IsEmpty() bool {
	return array.length == 0
}

//修改 index 位置的元素
func (array *Array) Set(index uint, value int) (err error) {
	if err = array.checkIndex(index); err != nil {
		return
	}

	array.data[index] = value
	return
}

//获取对应 index 位置的元素
func (array *Array) Get(index uint) (err error, value int) {
	if err = array.checkIndex(index); err != nil {
		return
	}

	value = array.data[index]
	return
}

//通过索引查找数组，索引范围[0,n-1](未找到，返回 -1)
func (array *Array) Find(value int) int {
	for i := 0; i < array.length; i++{
		if array.data[i] == value {
			return array.data[i]
		}
	}

	return -1
}

//在 index 位置，插入元素e, 时间复杂度 O(m+n)
func (array *Array) Add(index uint, value int) (err error) {
	if err = array.checkIndex(index); err != nil {
		return
	}

	// 如果当前元素个数等于数组容量，则将数组扩容为原来的2倍
	if array.length == array.GetCapacity() {
		array.resize(uint(array.length * 2))
	}


	for i:= array.length-1; i >= int(index); i-- {
		array.data[i+1] = array.data[i]
	}

	array.data[index] = value
	array.length++
	return
}

//向数组头插入元素
func (array *Array) AddFirst(value int) error {
	return array.Add(0,value)
}

//向数组尾插入元素
func (array *Array) AddLast(value int) error {
	return array.Add(uint(array.length),value)
}

// 删除 index 位置的元素，并返回
