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
	"fmt"
)

func TestNewArray(t *testing.T) {
	array := array.NewArray(2)
	t.Log(array)


	mySlice := make([]int, 5, 10)

	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
}
