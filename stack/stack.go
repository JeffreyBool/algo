/**
  * Author: JeffreyBool
  * Date: 2019/4/22
  * Time: 12:25
  * Software: GoLand
*/

package stack

type Stack interface {
	GetSize() int
	IsEmpty() bool
	Push(interface{}) error
	Pop() (interface{},error)
	Peek() interface{}
}
