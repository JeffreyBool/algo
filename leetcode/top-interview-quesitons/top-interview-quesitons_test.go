/**
  * Author: JeffreyBool
  * Date: 2019/4/12
  * Time: 22:40
  * Software: GoLand
*/

package top_interview_quesitons

import (
	"testing"
)

func TestSingleNumber(t *testing.T)  {
	number := SingleNumber([]int{4, 1, 2, 1, 2})
	//00000100  4
	//00000000  0
	//00000100 10进制结果为4

	//00000001  1
	//00000100  4
	//00000101 10进制结果为5

	//00000101  5
	//00000010  2
	//00000111 10进制结果为7

	//00000111  7
	//00000001  1
	//00000110 10进制结果为6

	//00000110  6
	//00000010  2
	//00000100 10进制结果为4

	t.Log(number)
}