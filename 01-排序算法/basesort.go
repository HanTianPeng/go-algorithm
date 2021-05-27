/*
Author: Conk
Function: BaseSort
CreateTime: 2021/5/27 11:50 下午
UpdateTime: 2021/5/27 11:50 下午
思想: 基于计数排序进行实现
执行过程:
平均时间复杂度: O(n*k)
空间复杂度: O(n+k)
稳定性: 稳定
*/

package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func GetEleLenMax(arr []int) int {
	count := 0
	for _, value := range arr {
		lenValue := len(strconv.Itoa(value))
		if lenValue > count {
			count = lenValue
		}
	}
	return count
}

// BaseSort: 基数排序
func BaseSort(arr []int) []int {
	// 获取数组中元素的最大位数值
	maxLen := GetEleLenMax(arr)

	count := make([]int, 10)
	result := make([]int, len(arr))

	// 遍历
	for i := 0; i < maxLen; i++ {
		division := int(math.Pow10(i))

		// 计数数组
		for j := 0; j < len(arr); j++ {
			index := arr[j] / division % 10
			count[index] = count[index] + 1
		}

		// 计数数组转变成累加数组
		for k := 1; k < len(count); k++ {
			count[k] = count[k] + count[k - 1]
		}

		//fmt.Println("ppp--->", count)
		// 累加数组保持稳定性,进行遍历赋值
		for l := len(arr) - 1; l >= 0; l-- {
			index := arr[l] / division % 10
			//fmt.Println("ooo==>", index)
			result[count[index] - 1] = arr[l]
			count[index] = count[index] - 1
		}
		count = make([]int, 10)
		//fmt.Println("result===>", result)
	}

	return result
}

// GetTestDataV2: 根据count,随机获取指定数量的数组
func GetTestDataV2(count int) []int {
	var result []int

	// 循环遍历随机获取值
	for i := 0; i < count; i++ {
		result = append(result, rand.Intn(10000))
	}
	return result
}

func GetRandomData() int {
	return rand.Intn(10000)
}

func main() {
	for i := 0; i < 10; i++ {
		testData := GetTestDataV2(4)
		fmt.Println(i, "--->", testData)
		newTestData := BaseSort(testData)
		fmt.Println(i, "===>", newTestData)
	}

	testData03 := []int{5, 1, 2, 2, 3, 4, 4, 8}
	testData03 = BaseSort(testData03)
	fmt.Println(testData03)

	testData04 := []int{2, 4, 6, 5, 9, 8, 3, 1}
	testData04 = BaseSort(testData04)
	fmt.Println(testData04)
}
