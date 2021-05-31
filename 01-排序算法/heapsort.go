/*
Author: Conk
Function: HeapSort
CreateTime: 2021/5/31 11:35 下午
UpdateTime: 2021/5/31 11:35 下午
堆排序:
什么是堆:
	堆是一个完全二叉树
	堆上的任意节点值都必须大于等于（大顶堆）或小于等于（小顶堆）其左右子节点值
大顶堆: 堆上的任意节点都大于等于子节点值<根节点是堆中最大的元素>
小顶堆: 堆上的任意节点都小于等于子节点值<根节点是堆中最小的元素>
堆其实可以用一个数组表示:
	给定一个节点的下标i，左子节点为2i+1，右子节点为2i+2
时间复杂度: O(nLog2n)
空间复杂度: O(1)
*/

package main

import (
	"fmt"
	"math/rand"
)

// Heap: 堆
type Heap struct {
	arr []int
	size int
}

// SwapV4: 交换
func SwapV4(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// Heapify: 堆化-小顶堆
func Heapify(heap Heap, parentNode int) {
	leftNode := parentNode * 2 + 1
	rightNode := parentNode * 2 + 2

	minNode := parentNode
	if leftNode < heap.size && heap.arr[leftNode] > heap.arr[minNode] {
		minNode = leftNode
	}
	if rightNode < heap.size && heap.arr[rightNode] > heap.arr[minNode] {
		minNode = rightNode
	}

	if minNode != parentNode {
		SwapV4(heap.arr, minNode, parentNode)
		Heapify(heap, minNode)
	}
}

// BuildHeap: 建堆
func BuildHeap(arr []int) (heap Heap) {
	heap.arr = arr
	heap.size = len(arr)

	for i := heap.size / 2; i >= 0; i-- {
		Heapify(heap, i)
	}
	return
}

// HeapSort: 堆排序
func HeapSort(arr []int) {
	heap := BuildHeap(arr)

	for heap.size > 0 {
		SwapV4(heap.arr, 0, heap.size - 1)
		heap.size--
		Heapify(heap, 0)
	}
}

// InitTestDataV2: 初始化测试用例数据
func InitTestDataV2(count int) []int {
	if count <= 0 {
		count = 10
	}
	result := make([]int, 0)
	// 随机生成0-10之间的数字
	for i := 0; i < count; i++ {
		result = append(result, rand.Intn(10))
	}
	return result
}

func main() {
	testData01 := []int{1, 2, 3, 4, 5}
	HeapSort(testData01)
	fmt.Println(testData01)

	testData02 := []int{5, 4, 3, 2, 1}
	HeapSort(testData02)
	fmt.Println(testData02)

	testData03 := []int{5, 1, 2, 2, 3, 4, 4, 8}
	HeapSort(testData03)
	fmt.Println(testData03)

	testData04 := []int{2, 4, 6, 5, 9, 8, 3, 1}
	HeapSort(testData04)
	fmt.Println(testData04)

	// 随机生成测试用例数据进行数据验证
	for i := 0; i < 20; i++ {
		data := InitTestDataV2(5)
		HeapSort(data)
		fmt.Println(data)
	}
}
