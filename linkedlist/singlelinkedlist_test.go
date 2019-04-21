/**
  * Author: JeffreyBool
  * Date: 2019/4/21
  * Time: 18:04
  * Software: GoLand
*/

package linkedlist_test

import (
	"testing"
	"algo/linkedlist"
)

func TestNewListNode(t *testing.T) {
	node := linkedlist.NewListNode(1)
	t.Log(node.GetNext())
	t.Log(node.GetValue())
}

func TestLinkedList_AddFirst(t *testing.T) {
	linked := linkedlist.NewLinkedList()
	linked.AddFirst("first")
	t.Logf("cound: %d \n",linked.Count())
	t.Logf("%+v \n",linked.GetHead())
}