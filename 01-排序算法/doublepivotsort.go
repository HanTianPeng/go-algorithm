/*
Author: Conk
Function: DoublePivotSort
CreateTime: 2021/5/24 11:19 下午
UpdateTime: 2021/5/24 11:19 下午
01排序算法-双轴排序:
思想: 快速排序
执行过程:
	快速排序基于单轴进行排序，采用左右双指针夹逼方法实现
	双轴排序是基于快速排序实现,先确定两轴，再基于左右双指针夹逼方法 + 单指针右移动进行实现
平均时间复杂度: O(nlog2n)
空间复杂度: O(log2n)
稳定性: 不稳定
*/

package main

import (
	"fmt"
	"math/rand"
)

func SwapV3(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// DoublePivotSort: 双轴排序
func DoublePivotSort(arr []int, leftBound, rightBound int) {
	if leftBound >= rightBound {
		return
	}
	if arr[leftBound] > arr[rightBound] {
		SwapV3(arr, leftBound, rightBound)
	}
	// 创建两个轴
	pivotLeft := arr[leftBound]
	pivotRight := arr[rightBound]

	// 创建左右双指针
	left := leftBound
	right := rightBound

	// 左右指针夹逼: 比左轴小的,数据保留在左区域
	for left + 1 < rightBound && arr[left + 1] < pivotLeft {
		left++
	}
	// 左右指针夹逼: 比右轴大的,数据保留在右区域
	for right - 1 > leftBound && arr[right - 1] > pivotRight {
		right--
	}

	// 创建一个单指针
	position := left + 1
	// 单指针向右移动
	for position < right {
		// 当前值小于左轴
		if arr[position] < pivotLeft {
			left++
			SwapV3(arr, left, position)
			position++
		}else if arr[position] >= pivotLeft && arr[position] <= pivotRight {
			// 当前值在左轴与右轴之间
			position++
		}else if arr[position] > pivotRight {
			// 当前值大于右轴
			right--
			SwapV3(arr, position, right)
		}
	}
	// 将左右轴值切换到目标位置
	SwapV3(arr, leftBound, left)
	SwapV3(arr, right, rightBound)

	// 分治递归进行双轴排序: 左 中 右 三个区间
	DoublePivotSort(arr, leftBound, left - 1)
	DoublePivotSort(arr, left + 1 , right - 1)
	DoublePivotSort(arr, right + 1, rightBound)
}

// InitData: 用例测试
func InitData(count int) []int {
	if count <= 0 {
		count = 1
	}
	var data []int
	for i := 0; i < count; i++ {
		data = append(data, rand.Intn(100))
	}
	return data
}

func main() {
	testData01 := []int{1, 2, 3, 4, 5}
	DoublePivotSort(testData01, 0, len(testData01) - 1)
	fmt.Println(testData01)

	testData02 := []int{5, 4, 3, 2, 1}
	DoublePivotSort(testData02, 0, len(testData02) - 1)
	fmt.Println(testData02)

	testData03 := []int{5, 1, 2, 2, 3, 4, 4, 8}
	DoublePivotSort(testData03, 0, len(testData03) - 1)
	fmt.Println(testData03)

	testData04 := []int{2, 4, 6, 5, 10, 8, 3, 1}
	DoublePivotSort(testData04, 0, len(testData04) - 1)
	fmt.Println(testData04)

	// 随机生成测试用例数据进行数据验证
	for i := 0; i < 20; i++ {
		data := InitData(5)
		DoublePivotSort(data, 0, len(data) - 1)
		fmt.Println(data)
	}
}
