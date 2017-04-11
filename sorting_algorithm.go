package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	num      = 100000
	rangeNum = 100000
)

func main() {
	// fmt.Println(rand.NewSource(time.Now().Unix() + time.Now().UnixNano()))
	randSeed := rand.New(rand.NewSource(time.Now().Unix() + time.Now().UnixNano()))
	var buf []int
	for i := 0; i < num; i++ {
		buf = append(buf, randSeed.Intn(rangeNum))
	}
	// fmt.Println(buf)
	t := time.Now()
	// maopao1(buf) //18s
	// xuanze2(buf) //20s
	// charu3(buf) //3.8s
	// xier4(buf) //27ms
	// kuaisu5(buf) //12ms
	// guibing6(buf) //15ms
	// duipai7(buf) //20ms

	fmt.Println(time.Since(t))
}

//堆排序
func duipai7(buf []int) {
	temp, n := 0, len(buf)
	for i := (n - 1) / 2; i >= 0; i-- {
		MixHeapFixdown(buf, i, n)
	}
	for i := n - 1; i > 0; i-- {
		temp = buf[0]
		buf[0] = buf[i]
		buf[i] = temp
		MixHeapFixdown(buf, 0, i)
	}
	// fmt.Println(buf)
}

func MixHeapFixdown(a []int, i, n int) {
	j := 2*i + 1
	for j < n {
		if j+1 < n && a[j+1] < a[j] {
			j++
		}
		if a[i] <= a[j] {
			break
		}
		a[i], a[j] = a[j], a[i]

		i = j
		j = 2*i + 1
	}
}

//归并排序
func guibing6(buf []int) {
	tmp := make([]int, len(buf))
	mergeSort(buf, 0, len(buf)-1, tmp)
	// fmt.Println(buf)
}
func mergeSort(a []int, first, last int, tmp []int) {
	if first < last {
		middle := (first + last) / 2
		mergeSort(a, first, middle, tmp)
		mergeSort(a, middle+1, last, tmp)
		mergeArray(a, first, middle, last, tmp)
	}
}
func mergeArray(a []int, first, middle, end int, tmp []int) {
	i, m, j, n, k := first, middle, middle+1, end, 0
	for i <= m && j <= n {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			k++
			i++
		} else {
			tmp[k] = a[j]
			k++
			j++
		}
	}
	for i <= m {
		tmp[k] = a[i]
		k++
		i++

	}
	for j <= n {
		tmp[k] = a[j]
		k++
		j++
	}
	for ii := 0; ii < k; ii++ {
		a[first+ii] = tmp[ii]
	}
}

//快速排序
func kuaisu5(buf []int) {
	kuai(buf, 0, len(buf)-1)
	// fmt.Println("kuaisu times: ", times)
	// fmt.Println(buf)
}
func kuai(a []int, l, r int) {
	if l >= r {
		return
	}
	i, j, key := l, r, a[l]
	for i < j {
		for i < j && a[j] > key {
			j--
		}
		if i < j {
			a[i] = a[j]
			i++
		}
		for i < j && a[i] < key {
			i++
		}
		if i < j {
			a[j] = a[i]
			j--
		}
	}
	//i==j
	a[i] = key
	kuai(a, l, i-1)
	kuai(a, i+1, r)

}

//希尔排序
func xier4(buf []int) {
	times := 0
	// tmp := 0
	length := len(buf)
	incre := length
	for {
		incre /= 2
		for k := 0; k < incre; k++ {
			for i := k + incre; i < length; i += incre {
				for j := i; j > k; j -= incre {
					times++
					if buf[j] < buf[j-incre] {
						buf[j-incre], buf[j] = buf[j], buf[j-incre]
					} else {
						break
					}
				}
			}
		}
		if incre == 1 {
			break
		}
	}
	// fmt.Println("xier times: ", times)
	// fmt.Println(buf)
}

//插入排序
func charu3(buf []int) {
	times := 0
	for i := 1; i < len(buf)-1; i++ {
		for j := i; j > 0; j-- {
			if buf[j] < buf[j-1] {
				times++
				buf[j-1], buf[j] = buf[j], buf[j-1]
			} else {
				break
			}
		}
	}
	// fmt.Println("charu times: ", times)
	// fmt.Println(buf)
}

//选择排序
func xuanze2(buf []int) {
	times := 0
	for i := 0; i < len(buf)-1; i++ {
		min := i
		for j := i; j < len(buf); j++ {
			times++
			if buf[min] > buf[j] {
				min = j
			}
		}
		if min != i {
			buf[i], buf[min] = buf[min], buf[i]
		}
	}
	fmt.Println("xuanze times: ", times)
	fmt.Println(buf)
}

// 冒泡排序
func maopao1(buf []int) {
	times := 0
	for i := 0; i < len(buf)-1; i++ {
		flag := false
		for j := 1; j < len(buf)-i; j++ {
			if buf[j-1] > buf[j] {
				times++
				buf[j-1], buf[j] = buf[j], buf[j-1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	fmt.Println("maopao times:", times)
	fmt.Println(buf)
}
