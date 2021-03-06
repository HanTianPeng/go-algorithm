/*
Author: Conk
Function: SelectSort
CreateTime: 2021-05-18 00:10:00
UpdateTime: 2021-05-18 00:10:00
01排序算法-选择排序:
思想:
	在数组中找到最小值放到0号索引位置
	在数组中找到第二小值放到1号索引位置
执行过程:
	5 4 3 2 1
	第一轮: 最小5, 最小索引0
	5 > 4 : 最小4，最小索引1
	4 > 3 : 最小3，最小索引2
	3 > 2 : 最小2, 最小索引3
	2 > 1 : 最小1, 最小索引4
	0 != 4 : 1 4 3 2 5
	第二轮: 第二小4, 第二小索引1
	4 > 3 : 最小3，最小索引2
	3 > 2 : 最小2, 最小索引3
	2 < 5 : 最小2, 最小索引3
	3 != 1 : 1 2 3 4 5
	以此类推....
优化点:
	一次循环找出最小值，最大值

平均时间复杂度: O(n^2)
空间复杂度: 1
稳定性: 不稳定
特点: 不稳定
*/

package main

import "fmt"

// swap: 交换函数
func swap(arr []int, i, j int) {
	arr[j], arr[i] = arr[i], arr[j]
}

// SelectSort: 选择排序
func SelectSort(arr []int) {
	// 细节优化点1: len(arr) - 1; 外层少循环一次
	for i := 0; i < len(arr) - 1; i++ {
		// 默认最小值在0号索引位置
		minIndex := i

		// 遍历找出最小值和最小值对应的索引位置
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		// 从0号索引开始锁定位置
		if minIndex != i {
			swap(arr, i, minIndex)
		}
	}
}


func main() {
	testData01 := []int{1, 2, 3, 4, 5}
	SelectSort(testData01)
	fmt.Println(testData01)

	testData02 := []int{5, 4, 3, 2, 1}
	SelectSort(testData02)
	fmt.Println(testData02)

	testData03 := []int{5, 1, 2, 2, 3, 4, 4, 8}
	SelectSort(testData03)
	fmt.Println(testData03)

	testData04 := []int{2, 4, 6, 5, 10, 8, 3, 1}
	SelectSort(testData04)
	fmt.Println(testData04)
}