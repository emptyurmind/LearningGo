package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 9, 8, 3, 10, 1, 5, 7, 2, 6}
	fmt.Println(bubbleSort(arr))
	fmt.Println(selectionSort(arr))
	fmt.Println(insertionSort(arr))
	fmt.Println(shellSort(arr))
	fmt.Println(mergeSort(arr))
	fmt.Println(quickSort(arr))
}

// 冒泡排序
func bubbleSort(arr []int) []int {
	length := len(arr)
	// 外层for控制交换多少次，因为每次都会把最大值浮到表面
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func selectionSort(arr []int) []int {
	// 选择排序是假定前面已经排好了，需要找到这个位置该有的数
	// 首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置。
	//再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
	length := len(arr)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length-1; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func insertionSort(arr []int) []int {
	// 工作原理:通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
	for i := range arr {
		// 将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。
		//从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。
		pre := i - 1 // 起初只和前一个比
		for pre >= 0 && arr[pre] > arr[i] {
			arr[pre+1] = arr[pre]
			pre--
		}
		arr[pre+1] = arr[i]
	}
	return arr
}

func shellSort(arr []int) []int {
	// 希尔排序，也称递减增量排序算法，是插入排序的一种更高效的改进版本。但希尔排序是非稳定排序算法。
	// 基本思想：先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，待整个序列中的记录"基本有序"时，再对全体记录进行依次直接插入排序。
	length := len(arr)
	gap := 1
	for gap < length/3 {
		gap = gap*3 + 1 // //动态定义间隔序列
	}
	for gap > 0 {
		// 中间就是插入排序，只不过是从后往前
		for i := gap; i < length; i++ {
			pre := i - gap
			for pre >= 0 && arr[pre] > arr[i] {
				arr[pre+gap] = arr[pre]
				pre -= gap
			}
			arr[pre+gap] = arr[i]
		}
		gap = gap / 3
	}
	return arr
}

func mergeSort(arr []int) []int {
	// 分治
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[:middle]
	right := arr[middle:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	var res []int
	for len(left) != 0 && len(right) != 0 {
		// 合并两个有序数组
		if left[0] < right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	// 处理剩余的
	for len(left) != 0 {
		res = append(res, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		res = append(res, right[0])
		right = right[1:]
	}
	return res
}

func quickSort(arr []int) []int {
	// 快速排序使用分治法（Divide and conquer）策略来把一个串行（list）分为两个子串行（sub-lists）
	// 快速排序的最坏运行情况是 O(n²)，比如说顺序数列的快排。但它的平摊期望时间是 O(nlogn)，
	// 且 O(nlogn) 记号中隐含的常数因子很小，比复杂度稳定等于 O(nlogn) 的归并排序要小很多。所以，对绝大多数顺序性较弱的随机数列而言，快速排序总是优于归并排序。
	return quick(arr, 0, len(arr)-1)
}

func quick(arr []int, left int, right int) []int {
	if left < right {
		partition := partition(arr, left, right)
		quick(arr, left, partition-1)
		quick(arr, partition+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	// 选定一个pivot，然后以他为中心轴，移动元素到之前或之后
	pivot := left
	idx := pivot + 1
	for i := idx; i < right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[idx] = arr[idx], arr[i]
			idx++
		}
	}
	arr[pivot], arr[idx-1] = arr[idx-1], arr[pivot]
	return idx - 1
}
