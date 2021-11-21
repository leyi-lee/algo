package dp

import "fmt"

func MaxProfit(price []int) int {
	n := len(price)
	prices := make([]int,  1)
	prices[0] = 0
	prices = append(prices, price...)
	fmt.Println(prices)


	f := make([][]int, n + 1)
	for i := range f {
		f[i] = make([]int, 2)
		f[i][0] = -100000
		f[i][1] = -100000
	}

	// 持仓 0 1
	f[0][0] = 0

	for i := 1; i <= n; i++ {

		if i == 4 {
			fmt.Println(f[i - 1], f[i][1])
		}
		f[i][1] = max(f[i][1], 0-prices[i])
		f[i][0] = max(f[i][0], f[i - 1][1] + prices[i])


		f[i][1] = max(f[i][1], f[i - 1][1])
		f[i][0] = max(f[i][0], f[i - 1][0])


		//for j := 0; j < 2; j++ {
		//	f[i][j] = max(f[i][j], f[i - 1][j])
		//}
	}

	fmt.Println(f)

	ans := f[n][0]
	return ans
}


func max(x int, y int) int {
	if x > y  {
		return x
	}
	return y
}

func coinChange() int {
	coins := []int{1,2,5}
	amount := 11

	opt := make([]int, amount + 1)

	opt[0] = 0

	for i := 0; i < amount;i++ {
		opt[i + 1] = -1000000
		for _,v := range coins {
			if i + v <= amount {
				opt[i + 1] = max(opt[i + 1], opt[i + v] + 1)
			}
		}
	}

	if opt[amount] == -1000000 {
		return -1
	}
	return opt[amount]
}

func min(x int, y int) int {
	if x > y  {
		return y
	}
	return x
}