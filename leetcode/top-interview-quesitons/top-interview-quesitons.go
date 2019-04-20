/**
  * Author: JeffreyBool
  * Date: 2019/4/12
  * Time: 22:39
  * Software: GoLand
*/

package top_interview_quesitons

import (
	"fmt"
)

func SingleNumber(nums []int) int {
	ans := 0
	for	_,num := range nums {
		fmt.Printf("ans: %d \n",ans)
		ans ^= num
		fmt.Printf("%d  ^ %d 的值为: %d \n",ans,num,ans)
	}

	return ans
}