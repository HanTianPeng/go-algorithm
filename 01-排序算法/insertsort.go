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
*/

package main

import "fmt"

func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1

		for insertIndex >= 0 && insertVal < arr[insertIndex] {
			arr[insertIndex + 1] = arr[insertIndex]
			insertIndex--
		}

		// 插入
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

	testData03 := []int{5, 1, 2, 2, 3, 4, 4}
	InsertSort(testData03)
	fmt.Println(testData03)
}