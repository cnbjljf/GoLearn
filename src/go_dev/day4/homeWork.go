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
	/* 冒泡算法不做解释，太简单了 */
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
		// 下面这个循环是说当前位置大于0说明开始循环到第二个数了，
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
		for leftFlag < rightFlag && k < intSlice[rightFlag] {
			// 开始交换,把比k小的数（array[right_flag] 放到左边）
			rightFlag--
		}

		intSlice[leftFlag] = intSlice[rightFlag]
		intSlice[rightFlag] = k

		// 左边的下标开始向右移动
		for leftFlag < rightFlag && k >= intSlice[leftFlag] {
			// 原理同上，left_flag +=1只是不断找比k大的数
			leftFlag++
		}

		// 开始交换,把比k大的数（array[right_flag] 放到右边）
		intSlice[rightFlag] = intSlice[leftFlag]
		intSlice[leftFlag] = k

	}

	kuaiShu(intSlice, start, leftFlag-1) //  对左边的数据排序，递归算法
	kuaiShu(intSlice, leftFlag+1, end)   // 对右边的数据排序,递归算法
	return intSlice
}

func xuanZhe(intSlice []int) []int {
	/*
		选择排序，排序思想如下：
		假设一个数组aa[4,3，6,1,23]
		1. 对比数组中第一个元素4和第二个元素3，显然3比4小，那么我们用一个变量k来记住3的位置（也就是下标）
		2. 接着第二次比较，第二次比较拿3与6比较，显然3比6小，那么k的值不变，继续下一轮，
		3. 上面的k值如果没有找到比第二个元素3的话，那么k值就不变，如果找到了比3小的话，那么k的值要变。
		4. 循环完成后，那么k值就是就是这个数组最小那个数的下标了。然后就进行判断，如果这个数的下标不是
		第一个元素的下标，那么就让第一个元素与下标为k的元素交换下，这么整个数组最小的数就到了数组第一位，
		同理可得找出第二个小的数，然后与第二个元素交换位置......
	*/
	lenSlice := len(intSlice)
	var k int
	for i := 0; i < lenSlice; i++ {
		for j := i + 1; j < lenSlice; j++ {
			if intSlice[j] > intSlice[j-1] {
				tmp := intSlice[j]
				intSlice[j] = intSlice[j-1]
				intSlice[j-1] = tmp
				k = j //  用一个变量k来记住当前两数最小值位置（也就是下标）
			} else {
				k = j // 如果当前的数小于等于前一位数，那么说明这个数是两个数的最小值，下标为j
			}
		}
		if intSlice[k] != intSlice[i] { // 把当前标记的最小值与第i个元素调换
			tmp := intSlice[i]
			intSlice[i] = intSlice[k]
			intSlice[k] = tmp
		}
	}
	return intSlice
}

func main() {
	ss := [5]int{3, 1, 56, 10, 25}
	fmt.Println("冒泡算法", maoPao(ss[:]))
	fmt.Println("插入算法", chaRu(ss[:]))
	fmt.Println("快速排序", kuaiShu(ss[:], 0, len(ss)-1))
	fmt.Println("选择排序", xuanZhe(ss[:]))
}
