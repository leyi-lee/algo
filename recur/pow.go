package recur

import "math"

func MyPow(x float64, n int) float64 {
	if n == 0 { // 判断n的条件
		return 1.0
	}

	if n < 0 {
		return 1.0/myPow(x, -n)
	}

	temp := myPow(x, int(math.Ceil(float64(n/2))))
	ans := temp * temp
	if n % 2 == 1 {
		ans *= x
	}

	return ans
}
