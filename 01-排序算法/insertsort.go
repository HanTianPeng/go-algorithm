/*
Author: Conk
Function: InsertSort
CreateTime: 2021-05-18 00:10:00
UpdateTime: 2021-05-18 00:10:00
01排序算法-插入排序:
从第一个元素开始，该元素可以认为已经被排序
取出下一个元素，在已经排序的元素序列中从后向前扫描
如果该元素（已排序）小于新元素，将该元素移到下一位置
5 4 3 2 1
第一轮: 4
4 < 5 : 4 5 3 2 1
第二轮: 3
3 < 5 : 4 3 5 2 1
3 < 4 : 3 4 5 2 1
以此类推....

平均时间复杂度: O(n^2)
空间复杂度: 1
稳定性: 稳定
特点: 样本小且基本有序的时候效率较高
*/

package main

import "fmt"

func SwapV1(arr []int, i, j int) {
	arr[j], arr[i] = arr[i], arr[j]
}

// InsertSortV1: 插入排序,通过交换函数实现
func InsertSortV1(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			SwapV1(arr, j, j - 1)
		}
	}
}

// InsertSort: 插入排序, 类似打扑克牌
func InsertSort(arr []int) {
	// 从索引为1号位置开始,将1号位置作为第一个进行插入排序的元素
	for i := 1; i < len(arr); i++ {
		// 优化点:临时变量保存
		insertVal := arr[i]
		insertIndex := i - 1
		// 依次与前面元素进行对比: 比前面元素小,则将前面元素往后挪动一个位置
		for insertIndex >= 0 && insertVal < arr[insertIndex] {
			arr[insertIndex + 1] = arr[insertIndex]
			insertIndex--
		}

		// 判断如果当前位置与最开始作为插入元素的索引值是否相等: 不相等则需要插入
		if insertIndex + 1 != i {
			arr[insertIndex + 1] = insertVal
		}
	}
}


func main() {
	testData01 := []int{1, 2, 3, 4, 5}
	InsertSort(testData01)
	fmt.Println(testData01)

	testData02 := []int{5, 4, 3, 2, 1}
	InsertSort(testData02)
	fmt.Println(testData02)

	testData03 := []int{5, 1, 2, 2, 3, 4, 4, 8}
	InsertSort(testData03)
	fmt.Println(testData03)

	testData04 := []int{2, 4, 6, 5, 10, 8, 3, 1}
	InsertSort(testData04)
	fmt.Println(testData04)

	testData05 := []int{1, 2, 3, 4, 5}
	InsertSortV1(testData05)
	fmt.Println(testData05)

	testData06 := []int{5, 4, 3, 2, 1}
	InsertSortV1(testData06)
	fmt.Println(testData06)

	testData07 := []int{5, 1, 2, 2, 3, 4, 4, 8}
	InsertSortV1(testData07)
	fmt.Println(testData07)

	testData08 := []int{2, 4, 6, 5, 10, 8, 3, 1}
	InsertSortV1(testData08)
	fmt.Println(testData08)
}