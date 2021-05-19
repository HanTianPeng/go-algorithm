/*
Author: Conk
Function: SelectSort
CreateTime: 2021-05-19 00:10:00
UpdateTime: 2021-05-19 00:10:00
01排序算法-归并排序:
思想:
	分而治之
执行过程:
		5	4				3	2	1
	[5			4]		[3			2	1]
	[5]			[4]		[3]		[2]			[1]
		[4 5]				[3]		 [1 2]
							    [1 2 3]
				[1 2 3 4 5]
	先分后治

平均时间复杂度: O(nlog2n)
空间复杂度: n
稳定性: 稳定
*/

package main

import "fmt"

// Merge: 合并两个有序数组
func Merge(left, right int, arr, arrLeft, arrRight []int) {
	// arrLeft, arrRight均从0号索引开始比较
	leftIndex := 0
	rightIndex := 0
	// 注意: 遍历是起始点即终止条件: 否则arr[k]会被重复赋值
	for k := left; k < right; k++ {
		// arrLeft已经读取完毕
		if leftIndex >= len(arrLeft) {
			arr[k] = arrRight[rightIndex]
			rightIndex++
		}else if rightIndex >= len(arrRight) {
			// arrRight已经读取完毕
			arr[k] = arrLeft[leftIndex]
			leftIndex++
		}else if arrLeft[leftIndex] <= arrRight[rightIndex] {
			// 左小,右移左加一
			arr[k] = arrLeft[leftIndex]
			leftIndex++
		}else {
			// 右小,右指针加一
			arr[k] = arrRight[rightIndex]
			rightIndex++
		}
	}
}

/*
MergeSort: 归并排序
left:
right: 一定是数组长度
 */
func MergeSort(arr []int, left, right int) {
	if right - left <= 1 {
		return
	}
	middle := (left + right) / 2
	MergeSort(arr, left, middle)
	MergeSort(arr, middle, right)

	arrLeft := make([]int, middle - left)
	arrRight := make([]int, right - middle)

	copy(arrLeft, arr[left: middle])
	copy(arrRight, arr[middle: right])

	// 治
	Merge(left, right, arr, arrLeft, arrRight)
}


func main() {
	testData01 := []int{1, 2, 3, 4, 5}
	MergeSort(testData01, 0, len(testData01))
	fmt.Println(testData01)

	testData02 := []int{5, 4, 3, 2, 1}
	MergeSort(testData02, 0, len(testData02))
	fmt.Println(testData02)

	testData03 := []int{5, 1, 2, 2, 3, 4, 4, 8}
	fmt.Println()
	MergeSort(testData03, 0, len(testData03))
	fmt.Println(testData03)

	testData04 := []int{2, 4, 6, 5, 10, 8, 3, 1}
	MergeSort(testData04, 0, len(testData04))
	fmt.Println(testData04)
}



