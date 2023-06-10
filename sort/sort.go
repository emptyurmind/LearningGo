package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{4, 9, 8, 3, 10, 1, 5, 7, 2, 6}
	fmt.Println(bubbleSort(arr))
	fmt.Println(selectionSort(arr))
	fmt.Println(insertionSort(arr))
	fmt.Println(shellSort(arr))
	fmt.Println(mergeSort(arr))
	fmt.Println(quickSort(arr))
	fmt.Println(heapSort(arr))
	fmt.Println(countingSort(arr))
	fmt.Println(bucketSort(arr))
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

func heapSort(arr []int) []int {
	// 一般用数组来表示堆，下标为 i 的结点的父结点下标为(i-1)/2；其左右子结点分别为 (2i + 1)、(2i + 2)
	// ① 最大堆调整（Max_Heapify）：将堆的末端子节点作调整，使得子节点永远小于父节点
	// ② 创建最大堆（Build_Max_Heap）：将堆所有数据重新排序
	// ③ 堆排序（HeapSort）：移除位在第一个数据的根节点，并做最大堆调整的递归运算

	// ① 先将初始的R[0…n-1]建立成最大堆，此时是无序堆，而堆顶是最大元素。
	// ② 再将堆顶R[0]和无序区的最后一个记录R[n-1]交换，由此得到新的无序区R[0…n-2]和有序区R[n-1]，且满足R[0…n-2].keys ≤ R[n-1].key
	// ③ 由于交换后新的根R[1]可能违反堆性质，故应将当前无序区R[1..n-1]调整为堆。然后再次将R[1..n-1]中关键字最大的记录R[1]和该区间的最后一个记录R[n-1]交换，由此得到新的无序区R[1..n-2]和有序区R[n-1..n]，且仍满足关系R[1..n-2].keys≤R[n-1..n].keys，同样要将R[1..n-2]调整为堆。
	// ④ 直到无序区只有一个元素为止。

	arrLen := len(arr)
	buildMaxHeap(arr, arrLen)
	fmt.Printf("arr's content is %d\n", arr)
	for i := arrLen - 1; i >= 0; i-- {
		// 对堆化数组排序，移除根节点，遍历长度-1
		swap(arr, 0, i)
		arrLen -= 1
		heapify(arr, 0, arrLen)
	}
	return arr
}

func buildMaxHeap(arr []int, arrLen int) {
	// 将数组堆化
	// 注意：这里只是堆化，堆化不代表整个堆已经是有序的了，只是符合大根堆/小根堆的性质
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
}

func heapify(arr []int, i, arrLen int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		swap(arr, i, largest)
		heapify(arr, largest, arrLen)
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func countingSort(arr []int) []int {
	// （1）找出待排序的数组中最大和最小的元素
	// （2）统计数组中每个值为i的元素出现的次数，存入数组C的第i项
	// （3）对所有的计数累加（从C中的第一个元素开始，每一项和前一项相加）
	// （4）反向填充目标数组：将每个元素i放在新数组的第C(i)项，每放一个元素就将C(i)减去1
	max := math.MinInt
	for _, value := range arr {
		if max < value {
			fmt.Printf("i's value is %d\n", value)
			max = value
		}
	}
	fmt.Printf("max's value is %d\n", max)
	bucketLen := max + 1
	bucket := make([]int, bucketLen) // 初始为0的数组

	sortedIndex := 0

	for i := range arr {
		bucket[arr[i]] += 1
	}

	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			arr[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}

	return arr
}

func bucketSort(arr []int) []int {
	// 桶排序是计数排序的升级版。它利用了函数的映射关系，高效与否的关键就在于这个映射函数的确定
	buckets := make([][]int, 5)
	res := make([]int, 0)
	fmt.Printf("buckets's size is %d\n", len(buckets))
	for i := range buckets {
		buckets[i] = make([]int, 0)
	}
	min := math.MaxInt
	max := math.MinInt
	for _, val := range arr {
		if val < min {
			min = val
		} else if val > max {
			max = val
		}
	}
	for _, val := range arr {
		idx := int(math.Floor(float64(val-min)) / float64(len(buckets)))
		buckets[idx] = append(buckets[idx], val)
	}
	for _, bucket := range buckets {
		insertionSort(bucket)
		for _, val := range bucket {
			res = append(res, val)
		}
	}
	return res
}
