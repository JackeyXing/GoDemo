package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	list := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		list = append(list, i+1)
	}

	rand.Seed(time.Now().UnixNano()) //设置随机种子
	for i := 0; i < 20; i++ {
		v := rand.Intn(1000000 - 1) + 1 //生成1-1000000的随机数
		fmt.Printf("针对 %d 进行二分查找：\n", v)
		idex := binarySearch(list, v)
		fmt.Printf("%d 的索引位置是:[%d]\n",v, idex)
		fmt.Println("-------------------------------------")
	}
}

func binarySearch(list []int, target int) int {
	low := 0
	high := len(list) - 1

	step := 0 //记录查找次数

	for {
		step++
		if low <= high {
			mid := (low + high) / 2
			guess := list[mid]
			if guess == target {
				fmt.Printf("共查找了 %d 次\n", step)
				return mid
			}
			if guess < target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
}
