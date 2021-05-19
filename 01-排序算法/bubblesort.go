/*
Author: Conk
Function: BubbleSort
CreateTime: 2021-05-18 00:10:00
UpdateTime: 2021-05-18 00:10:00
01排序算法-冒泡排序:
思想:
	比较相邻的元素，如果前一个比后一个大，就把它们两个调换位置
执行过程:
	5 4 3 2 1
	第一轮: 5
	5 > 4 : 4 5 3 2 1
	5 > 3 : 4 3 5 2 1
	5 > 2 : 4 3 2 5 1
	5 > 1 : 4 3 2 1 5
	第二轮: 4
	4 > 3 : 3 4 2 1 5
	4 > 2 : 3 2 4 1 5
	4 > 1 : 3 2 1 4 5
	以此类推....

平均时间复杂度: O(n^2)
空间复杂度: 1
稳定性: 稳定
*/

package main

import "fmt"

// swap: 交换函数
func Swap(arr []int, i, j int) {
	arr[j], arr[i] = arr[i], arr[j]
}

// BubbleSort: 冒泡排序
func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				Swap(arr, i, j)
			}
		}
	}
}


func main() {
	testData01 := []int{1, 2, 3, 4, 5}
	BubbleSort(testData01)
	fmt.Println(testData01)

	testData02 := []int{5, 4, 3, 2, 1}
	BubbleSort(testData02)
	fmt.Println(testData02)

	testData03 := []int{5, 1, 2, 2, 3, 4, 4}
	BubbleSort(testData03)
	fmt.Println(testData03)
}