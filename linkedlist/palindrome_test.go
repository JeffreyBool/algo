/**
  * Author: JeffreyBool
  * Date: 2019/4/21
  * Time: 22:00
  * Software: GoLand
*/

package linkedlist

import (
	"testing"
)

func TestIsPalindrome1(t *testing.T) {
	strs := []string{"heooeh", "hello", "heoeh", "a", ""}
	for _, str1 := range strs {
		linked := NewLinkedList()
		for _, v := range str1 {
			linked.AddLast(string(v))
		}

		linked.Print()
		t.Log(IsPalindrome1(linked))
	}
}
