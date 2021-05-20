/*
Author: Conk
Function: QuickSort
CreateTime: 2021-05-19 00:10:00
UpdateTime: 2021-05-19 00:10:00
01排序算法-快速排序:
思想:
	基准 + 分而治之
执行过程:
	2, 4, 6, 5, 10, 8, 3, 1
	基准: 5
	左边都比5小,右边都比5大:
		2, 4, 1, 3, 10, 8, 5, 6
		l: 4
		r: 3
		左边: 2, 4, 1, 3 left=0 right=3
		右边: 10, 8, 5, 6 left = 4 right = 7
	以此类推....

平均时间复杂度: O(nlog2n)
空间复杂度: log2n
稳定性: 稳定
*/

package main

import "fmt"

// SwapV2: 交换函数
func SwapV2(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// QuickSort: 快速排序，从大到小排序
func QuickSort(arr []int, left, right int) {
	l := left
	r := right
	// 基准
	pivot := arr[(l + r) / 2]
	// 一次快排
	for l < r {
		for arr[l] < pivot {
			l++
		}
		for arr[r] > pivot {
			fmt.Println("ssss", r, l)
			r--
		}
		// 交换
		if l < r {
			SwapV2(arr, l, r)
			l++
			r--
		}
	}
	// 边界处理
	if l == r {
		l++
		r--
	}
	// 分治
	if left < r {
		QuickSort(arr, left, r)
	}
	if right > l {
		QuickSort(arr, l, right)
	}
}


func main() {
	testData01 := []int{1, 2, 3, 4, 5}
	QuickSort(testData01, 0, len(testData01) - 1)
	fmt.Println(testData01)

	testData02 := []int{5, 4, 3, 2, 1}
	QuickSort(testData02, 0, len(testData02) - 1)
	fmt.Println(testData02)

	testData03 := []int{5, 1, 2, 2, 3, 4, 4, 8}
	fmt.Println()
	QuickSort(testData03, 0, len(testData03) - 1)
	fmt.Println(testData03)

	testData04 := []int{2, 4, 6, 5, 10, 8, 3, 1}
	QuickSort(testData04, 0, len(testData04) - 1)
	fmt.Println(testData04)
}