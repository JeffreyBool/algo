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
	data []int
	length uint
}

//数组初始化内存
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}

	return &Array{
		data:make([]int,capacity,capacity),
		length:0,
	}
}

//数组长度
func (this *Array) Len()  uint {
	return this.length
}

//判断索引是否越界
func (this *Array) isIndexOutOfRange(index uint) bool {
	/*
	**判断当前索引是否大于 数组的分配空间大小
	**cap()函数返回的是数组切片分配的空间大小
	*/
	if index >= uint(cap(this.data)) {
		return true
	}

	return false
}

//通过索引查找数组，索引范围[0,n-1]
func (this *Array) Find(index uint) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}

	return this.data[index],nil
}

