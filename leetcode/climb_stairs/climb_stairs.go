/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
*/
package main

import "fmt"

// //法一：递归 时间O^2n 空间On
// func climbStairs(n int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 2 {
// 		return 2
// 	}
// 	return climbStairs(n - 1) + climbStairs(n - 2)
// }

// //法二：记忆化递归 都是On
// func climbStairs(n int) int {
// 	memory := make(map[int]int)
// 	return climbStairsRecursion(n, memory)
// }
// func climbStairsRecursion(n int, memory map[int]int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 2 {
// 		return 2
// 	}
// 	if _, ok := memory[n]; ok {	//ok 为true 表示map中存在key为n的值
// 	return memory[n]
// 	}
// 	memory[n] = climbStairsRecursion(n-1, memory) + climbStairsRecursion(n-2, memory)
// 	return memory[n]
// }

// // 法三：动态规划 都是On
// func climbStairs(n int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	dp := make([]int,n+1)
// 	dp[1] = 1
// 	dp[2] = 2
// 	for i := 3; i <= n; i++ {
// 		dp[i] = dp[i - 1] + dp[i - 2]
// 	}
// 	return dp[n];
// }
//法四：滚动数组的动态规划 记录n-1和n-2两个状态 时间 On 空间O1
//滚动数组：https://blog.csdn.net/weixin_47750287/article/details/114021450?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522164934897416780271516104%2522%252C%2522scm%2522%253A%252220140713.130102334..%2522%257D&request_id=164934897416780271516104&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~top_positive~default-1-114021450.142^v7^pc_search_result_control_group,157^v4^control&utm_term=%E6%BB%9A%E5%8A%A8%E6%95%B0%E7%BB%84&spm=1018.2226.3001.4187
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	first := 1
	second := 2
	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}

func main() {
	fmt.Println(climbStairs(8))
}
