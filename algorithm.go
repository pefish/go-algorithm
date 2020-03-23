package go_algorithm

import (
	"math/rand"
	"time"
)

// 洗牌算法，不改变原值
func Shuffle(values []int) []int {
	result := make([]int, len(values))
	temp := append([]int{}, values...) // 深拷贝一份
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(temp) > 0 {
		n := len(temp)
		randIndex := r.Intn(n)
		temp[n-1], temp[randIndex] = temp[randIndex], temp[n-1]
		result[n-1] = temp[n-1]
		temp = temp[:n-1]
	}
	return result
}
