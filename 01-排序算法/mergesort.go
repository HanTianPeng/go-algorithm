/*
Author: Conk
Function: MergeSort
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

func MergeV2(arr []int, left, right, rightBound int) {
	// 创建三个指针
	leftIndex := left
	rightIndex := right
	k := 0
	temp := make([]int, rightBound - left + 1)

	// 遍历: 左右数组各依次获取从leftIndex,rightIndex元素进行对比,填充temp
	for leftIndex < right || rightIndex <= rightBound {
		// 左边已经读取完毕
		if leftIndex >= right {
			temp[k] = arr[rightIndex]
			rightIndex++
		}else if rightIndex > rightBound {
			// 右边已经读取完毕
			temp[k] = arr[leftIndex]
			leftIndex++
		}else if arr[leftIndex] <= arr[rightIndex] {
			// 左边小
			temp[k] = arr[leftIndex]
			leftIndex++
		}else {
			// 右边小
			temp[k] = arr[rightIndex]
			rightIndex++
		}
		k++
	}
	// 数据拷贝
	for j := 0; j < len(temp); j++ {
		arr[j + left] = temp[j]
	}
}

func MergeV1(arr []int, left, right,  rightBound int) {
	// 创建三个指针
	leftIndex := left
	rightIndex := right
	k := 0
	temp := make([]int, rightBound - left + 1)

	// 遍历: 左右数组各依次获取从leftIndex,rightIndex元素进行对比,填充temp
	for leftIndex < right && rightIndex <= rightBound {
		if arr[leftIndex] <= arr[rightIndex] {
			temp[k] = arr[leftIndex]
			leftIndex++
		}else {
			temp[k] = arr[rightIndex]
			rightIndex++
		}
		k++
	}
	// 右读取完毕
	for leftIndex < right {
		temp[k] = arr[leftIndex]
		k++
		leftIndex++
	}
	// 左读取完毕
	for rightIndex <= rightBound {
		temp[k] = arr[rightIndex]
		k++
		rightIndex++
	}

	for j := 0; j < len(temp); j++ {
		arr[left + j] = temp[j]
	}
}

// MergeSortV1: 归并排序
func MergeSortV1(arr []int, left, right int) {
	if left == right {
		return
	}
	middle := (left + right) / 2
	MergeSortV1(arr, left, middle)
	MergeSortV1(arr, middle+1, right)
	// 治
	MergeV1(arr, left, middle+1, right)
}


func main() {
	testData01 := []int{1, 2, 3, 4, 5}
	MergeSort(testData01, 0, len(testData01))
	fmt.Println(testData01)

	testData02 := []int{5, 4, 3, 2, 1}
	MergeSort(testData02, 0, len(testData02))
	fmt.Println(testData02)

	testData03 := []int{5, 1, 2, 2, 3, 4, 4, 8}
	MergeSort(testData03, 0, len(testData03))
	fmt.Println(testData03)

	testData04 := []int{2, 4, 6, 5, 10, 8, 3, 1}
	MergeSort(testData04, 0, len(testData04))
	fmt.Println(testData04)

	testData05 := []int{1, 2, 3, 4, 5}
	MergeSortV1(testData05, 0, len(testData05) - 1)
	fmt.Println(testData05)

	testData06 := []int{5, 4, 3, 2, 1}
	MergeSortV1(testData06, 0, len(testData06) - 1)
	fmt.Println(testData06)

	testData07 := []int{5, 1, 2, 2, 3, 4, 4, 8}
	MergeSortV1(testData07, 0, len(testData07) - 1)
	fmt.Println(testData07)

	testData08 := []int{2, 4, 6, 5, 10, 8, 3, 1}
	MergeSortV1(testData08, 0, len(testData08) - 1)
	fmt.Println(testData08)

	testData09 := []int{2, 4, 3, 1, 8, 7, 9, 5, 6, 10}
	MergeSortV1(testData09, 0, len(testData09) - 1)
	fmt.Println(testData09)
}



