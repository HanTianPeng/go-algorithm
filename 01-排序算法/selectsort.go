/*
Author: Conk
Function: select_sort
CreateTime: 2021-05-18 00:10:00
UpdateTime: 2021-05-18 00:10:00
01排序算法-选择排序:
在数组中找到最小值放到0号索引位置
在数组中找到第二小值放到1号索引位置
以此类推....

平均时间复杂度: O(n^2)
空间复杂度: 1
稳定性: 不稳定
*/

package main

import "fmt"

func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		// 默认最小值在0号索引位置
		maxVal := arr[i]
		maxIndex := i

		for j := i + 1; j < len(arr); j++ {
			if maxVal > arr[j] {
				maxVal = arr[j]
				maxIndex = j
			}
		}

		// 从0号索引开始锁定位置
		if maxIndex != i {
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
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

	testData03 := []int{5, 1, 2, 2, 3, 4, 4}
	SelectSort(testData03)
	fmt.Println(testData03)
}
