/*
Author: Conk
Function: CountSort
CreateTime: 2021/5/25 10:51 下午
UpdateTime: 2021/5/25 10:51 下午
思想: 计数排序
执行过程:
	固定范围的值,并将值映射到数组的索引，再统计个数
	5, 1, 2, 2, 3, 4, 4, 8
	[0, 0, 0, 0, 0, 0, 0, 0, 0]
	映射到索引进行统计
	[0, 1, 2, 1, 2, 1, 0, 0, 1]
	由计数数组转变成累加数组
	[0, 1, 3, 4, 6, 7, 7, 6, 7]
平均时间复杂度: O(n+k)
空间复杂度: O(n+k)
稳定性: 稳定
优点: 在已知值的范围进行排序
场景: 分数排名等
*/

package main

import (
	"fmt"
	"math/rand"
)

// CountSort: 计数排序
func CountSort(arr []int) []int {
	// 初始化一个计数器数组
	count := make([]int, 10)
	result := make([]int, 0)

	// 遍历原数组,按照计数器数组的索引进行统计
	for i := 0; i < len(arr); i++ {
		count[arr[i]] = count[arr[i]] + 1
	}

	// 遍历数组,重新赋值
	for j := 0; j < len(count); j++ {
		for count[j] > 0 {
			result = append(result, j)
			count[j] = count[j] - 1
		}
	}
	return result
}

// CountSortV1: 计数排序
func CountSortV1(arr []int) []int {
	// 初始化一个计数器数组
	count := make([]int, 10)
	result := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		count[arr[i]] = count[arr[i]] + 1
	}

	// 更新计数器数组为累加数组
	for j := 1; j < len(count); j++ {
		count[j] = count[j] + count[j - 1]
	}

	// 保持稳定性,进行遍历赋值
	for k := len(arr) - 1; k >= 0; k-- {
		result[count[arr[k]] - 1] = arr[k]
		// 赋值后需要将累加数组的值-1
		count[arr[k]] = count[arr[k]] - 1
	}
	return result
}

// InitTestData: 初始化测试用例数据
func InitTestData(count int) []int {
	if count <= 0 {
		count = 10
	}
	result := make([]int, 0)
	// 随机生成0-10之间的数字
	for i := 0; i < count; i++ {
		result = append(result, rand.Intn(10))
	}
	return result
}

func main() {
	testData01 := []int{1, 2, 3, 4, 5}
	testData01 = CountSortV1(testData01)
	fmt.Println(testData01)

	testData02 := []int{5, 4, 3, 2, 1}
	testData02 = CountSortV1(testData02)
	fmt.Println(testData02)

	testData03 := []int{5, 1, 2, 2, 3, 4, 4, 8}
	testData03 = CountSortV1(testData03)
	fmt.Println(testData03)

	testData04 := []int{2, 4, 6, 5, 9, 8, 3, 1}
	testData04 = CountSortV1(testData04)
	fmt.Println(testData04)

	// 随机生成测试用例数据进行数据验证
	for i := 0; i < 20; i++ {
		data := InitTestData(5)
		data = CountSortV1(data)
		fmt.Println(data)
	}
}