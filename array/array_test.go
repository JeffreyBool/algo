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
	if err := array.Set(0, 2); err != nil{
		t.Error(err)
	}

	//t.Logf("array cap: %d \n",array.GetCapacity())
	//t.Logf("array count: %d \n",array.Count())
	//t.Log(array)

	for i:= 0; i < 10; i++{
		if err  := array.Add(uint(i), i+1); err != nil{
			t.Error(err)
			break
		}
	}

	if err  := array.Add(10, 11); err != nil{
		t.Error(err)
	}
	t.Log(array)

	if err  := array.Add(10, 12); err != nil{
		t.Error(err)
	}

	t.Logf("test array cap: %d \n",array.GetCapacity())
	t.Logf("test array count: %d \n",array.Count())
	t.Log(array)
}
