package codepractice

import "fmt"

/*
121. 买卖股票的最佳时机
*/

func maxProfit(prices []int) int {
	length := len(prices)
	if length == 0 || length == 1 {
		return 0
	}

	// 从第一天开始计算起
	minPrice := prices[0]
	// 最大利润起始为0，动态更新
	maxProfit := 0
	for end := 1; end < length; end++ {
		// 最大的利润就是：当前记录的最大利润，卖当天的价格和买入时候的差价，两者中的最大值
		maxProfit = max(maxProfit, prices[end]-minPrice)
		// 更新完最大利润之后，需要更新最低的买入价格
		minPrice = min(minPrice, prices[end])
	}

	return maxProfit
}

func maxProfitOld(prices []int) int {
	length := len(prices)
	if length == 0 || length == 1 {
		return 0
	}

	maxPrice := 0
	for start := 0; start < length-1; start++ {
		for end := start + 1; end < length; end++ {
			money := prices[end] - prices[start]
			if money > maxPrice {
				maxPrice = money
			}
		}
	}

	return maxPrice
}

func HandlePratice09Entry() {
	p := []int{7, 1, 5, 3, 6, 4}
	// p = []int{7, 6, 4, 3, 1}
	// p = []int{0}
	fmt.Println(maxProfit(p))
}
