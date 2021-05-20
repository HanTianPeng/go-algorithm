/*
Author: Conk
Function: ShellSort
CreateTime: 2021-05-20 00:10:00
UpdateTime: 2021-05-20 00:10:00
01排序算法-希尔排序:
思想:
	步长 + 插入排序
执行过程:
9	6	11	3	5	12	8	7	10	15	14	4	1	13	2
步长为4:
9				5				10				1
开始插入排序:
1				5				9				10

	6				12				15				13
开始插入排序:
	6				12				13				15

		11				8				14				2
开始插入排序:
		2				8				11				14

			3				7				4
开始插入排序:
			3				4				7

最终:
1	6	2	3	5	12	8	4	9	13	11	7	10	15	14
步长-1，以此类推....

平均时间复杂度: O(n^1.3)
空间复杂度: 1
稳定性: 不稳定
特点: 快
*/

package main

import "fmt"

// InsertSortV2: 插入排序,类似打扑克牌
func InsertSortV2(arr []int, gap int) {
	// 从索引为gap号的位置开始, 将gap号位置作为第一个进行插入排序的元素
	for i := gap; i < len(arr); i++ {
		// 采用临时变量方式保存待插入元素
		insertVal := arr[i]
		insertIndex := i - gap
		// 依次与前面元素进行对比: 比前面元素小,则将前面元素往后挪动一个gap位置
		for insertIndex >= 0 && insertVal < arr[insertIndex] {
			arr[insertIndex + gap] = arr[insertIndex]
			insertIndex = insertIndex - gap
		}
		// 判断如果当前位置与最开始作为插入元素的索引值是否相等: 不相等则需要插入
		if insertIndex + gap != insertIndex {
			arr[insertIndex + gap] = insertVal
		}
	}
}

// ShellSort: 希尔排序
func ShellSort(arr []int) {
	for gap := len(arr) / 2; gap >= 1; gap = gap / 2 {
		// 插入排序
		InsertSortV2(arr, gap)
	}
}

// ShellSortV1: 希尔排序, knuth序列定义步长
func ShellSortV1(arr []int) {
	h := 1
	for h <= len(arr) / 3 {
		h = h * 3 + 1
	}

	for gap := h; gap >= 1; gap = (gap - 1) / 3 {
		// 插入排序
		InsertSortV2(arr, gap)
	}
}


func main() {
	testData01 := []int{1, 2, 3, 4, 5}
	ShellSort(testData01)
	fmt.Println(testData01)

	testData02 := []int{5, 4, 3, 2, 1}
	ShellSort(testData02)
	fmt.Println(testData02)

	testData03 := []int{5, 1, 2, 2, 3, 4, 4, 8}
	ShellSort(testData03)
	fmt.Println(testData03)

	testData04 := []int{2, 4, 6, 5, 10, 8, 3, 1}
	ShellSort(testData04)
	fmt.Println(testData04)

	testData05 := []int{1, 2, 3, 4, 5}
	ShellSortV1(testData05)
	fmt.Println(testData05)

	testData06 := []int{5, 4, 3, 2, 1}
	ShellSortV1(testData06)
	fmt.Println(testData06)

	testData07 := []int{5, 1, 2, 2, 3, 4, 4, 8}
	ShellSortV1(testData07)
	fmt.Println(testData07)

	testData08 := []int{2, 4, 6, 5, 10, 8, 3, 1}
	ShellSortV1(testData08)
	fmt.Println(testData08)
}
