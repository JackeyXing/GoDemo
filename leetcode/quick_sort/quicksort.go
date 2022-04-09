/*
基本思想
1.选定Pivot中心轴
2.将大于Pivot的数字放在Pivot的右边
3.将小于Pivot的数字放在Pivot的左边
4.分别对左右子序列重复前第三步操作
*/
package main

import "fmt"

//选择第一个为Pivot
func quickSort(arry []int, L int, R int) {
	if L >= R {
		return
	}
	left := L
	right := R
	pivot := arry[left]
	for left < right {
		for left < right && arry[right] >= pivot {
			right--
		}
		if left < right {
			arry[left] = arry[right]
		}
		for left < right && arry[left] <= pivot {
			left++
		}
		if left < right {
			arry[right] = arry[left]
		}
		if left >= right {
			arry[right] = pivot
		}
	}
	quickSort(arry, L, right-1)
	quickSort(arry, right+1, R)
}
func main() {
	arry := []int{2, 4, 1, 3, 1, 5}
	quickSort(arry, 0, len(arry)-1)
	fmt.Println(arry)
}
