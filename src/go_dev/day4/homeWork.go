// homeWork
package main

import (
	"fmt"
)

/*
实现一个冒泡排序
实现一个选择排序
实现一个插入排序
实现一个快速排序
*/

func maoPao(intSlice []int) []int {
	for i := 0; i < len(intSlice); i++ {
		for j := i + 1; j < len(intSlice); j++ {
			if intSlice[i] < intSlice[j] {
				tmp := intSlice[i]
				intSlice[i] = intSlice[j]
				intSlice[j] = tmp
			}
		}
	}
	return intSlice
}

func chaRu(intSlice []int) []int {
	/*
		插入排序思想基本是这样的：
		我们选取列表第2个数开始，把第一个数和第二个数对比，
		1 如果第一个数比第二个数大，那么调换下。
		2 如果第二个数比第一个数小，那么就不需要调换。
		3 依次类推，同理可得。。。
	*/
	for i := 1; i < len(intSlice); i++ {
		// 下面这个循环是说#当前位置大于0说明开始循环到第二个数了，
		// 而且当前列表元素的前一位(该元素左边第一位)大于当前的元素
		for i > 0 && intSlice[i-1] > intSlice[i] {
			currentNum := intSlice[i]
			intSlice[i] = intSlice[i-1]
			intSlice[i-1] = currentNum
			i = i - 1
		}

	}
	return intSlice
}

func kuaiShu(intSlice []int, start, end int) []int {
	/*
		快速排序：
		通过一趟排序将要排序的数据分割成独立的两部分，
		其中一部分的所有数据都比另外一部分的所有数据都要小，
		然后再按此方法对这两部分数据分别进行快速排序，
		整个排序过程可以递归进行，以此达到整个数据变成有序序列。
	*/

	if start >= end { // 意味着排序结束了
		return intSlice
	}

	k := intSlice[start] //  设K为中间数
	leftFlag := start    // 左侧数的下标，待会移动的时候就是通过下标移动
	rightFlag := end     // 右侧数的下标

	for leftFlag < rightFlag {
		if intSlice[leftFlag] > intSlice[rightFlag] {
		tmp:
			tmp := intSlice[rightFlag]
			intSlice[rightFlag] = intSlice[leftFlag]
			intSlice[leftFlag] = tmp
			leftFlag++
			rightFlag--
		}
	}
	return intSlice
}

func main() {
	ss := [5]int{3, 12, 56, 10, 25}
	// fmt.Println(maoPao(ss[:]))
	// fmt.Println(chaRu(ss[:]))
	fmt.Println(kuaiShu(ss[:], 0, 0))
}
