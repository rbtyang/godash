package sortdash

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/sync/errgroup"
)

//冒泡排序：
//  依次两两交换，大的往前冒，不断重复遍历，直到有序
func Bubble[T constraints.Ordered](sli []T) {
	leng := len(sli)
	for true {
		isSwap := false
		for i := 0; i < leng-1; i++ {
			if sli[i] > sli[i+1] {
				sli[i], sli[i+1] = sli[i+1], sli[i]
				isSwap = true
			}
		}
		if !isSwap {
			break
		}
	}
}

//冒泡排序：
func Bubble2[T constraints.Ordered](sli []T) {
	leng := len(sli)
	//该层循环控制 需要冒泡的轮数
	for i := 1; i < leng; i++ {
		//该层循环用来控制每轮 冒出一个数 需要比较的次数
		for j := 0; j < leng-1; j++ {
			if sli[i] < sli[j] {
				sli[i], sli[j] = sli[j], sli[i]
			}
		}
	}
}

//插入排序：
//  将 b区 中的 val 依次插到 a区中 最合适的位置
func Insertion[T constraints.Ordered](sli []T) {
	leng := len(sli)
	for b := 0; b < leng; b++ { //b区
		tmp := sli[b]
		for a := 0; a <= b; a++ { //a区
			if a == b { //位移到最后一个元素
				sli[a] = tmp //此时tmp存的是a区的最大值了
			} else if tmp < sli[a] { //找到比它大的了
				tmp, sli[a] = sli[a], tmp //插入并逐步位移后续元素
			}
		}
	}
}

//选择排序：
func Insertion2[T constraints.Ordered](sli []T) {
	leng := len(sli)
	for i := 0; i < leng; i++ {
		tmp := sli[i]
		//内层循环控制，比较并插入
		for j := i - 1; j >= 0; j-- {
			if tmp < sli[j] {
				//发现插入的元素要小，交换位置，将后边的元素与前面的元素互换
				sli[j+1], sli[j] = sli[j], tmp
			} else {
				//如果碰到不需要移动的元素，则前面的就不需要再次比较了。
				break
			}
		}
	}
}

//选择排序：
//  b区 中最小的 val 排到 a区 最后
func Selection[T constraints.Ordered](sli []T) {
	leng := len(sli)
	//双重循环完成，外层控制轮数，内层控制比较次数
	for pos := 0; pos < leng-1; pos++ {
		//先假设最小的值的位置
		minK := pos
		for j := pos + 1; j < leng; j++ {
			//sli[k] 是当前已知的最小值
			if sli[j] < sli[minK] {
				//比较，发现更小的,记录下最小值的位置；并且在下次比较时采用已知的最小值进行比较。
				minK = j
			}
		}
		//已经确定了当前的最小值的位置，保存到 k 中。如果发现最小值的位置与当前假设的位置 i 不同，则位置互换即可。
		if minK != pos {
			sli[pos], sli[minK] = sli[minK], sli[pos]
		}
	}
}

//快速排序：
//  1、先从数列中 取出一个数 作为基准数 val
//  2、分区过程，将比这个数大的数 全放到它的右边，小于或等于它的数 全放到它的左边 (右找小, 左找大)
//  3、再对左右区间 重复第二步，直到各区间 只有一个数
func Quick[T constraints.Ordered](sli []T) []T {
	leng := len(sli)
	//先判断是否需要继续进行
	if leng <= 1 {
		return sli
	}

	//选择第一个元素作为基准
	kVal := sli[0]
	//初始化左右两个切片
	lfSli, rtSli := []T{}, []T{}

	//遍历除了标尺外的所有元素，按照大小关系放入左右两个切片内
	for i := 1; i < leng; i++ {
		if sli[i] < kVal {
			//放入左边切片
			lfSli = append(lfSli, sli[i])
		} else {
			//放入右边切片
			rtSli = append(rtSli, sli[i])
		}
	}

	//再分别对左边和右边的切片进行相同的排序处理方式递归调用这个函数
	lfSort := Quick(lfSli)
	rtSort := Quick(rtSli)

	//合并
	return append(append(lfSort, kVal), rtSort...)
}

//快速排序：
//  1、先从数列中 取出一个数 作为基准数 val
//  2、分区过程，将比这个数大的数 全放到它的右边，小于或等于它的数 全放到它的左边 (右找小, 左找大)
//  3、再对左右区间 重复第二步，直到各区间 只有一个数
func QuickParallel[T constraints.Ordered](sli []T) []T {

	leng := len(sli)
	//先判断是否需要继续进行
	if leng <= 1 {
		return sli
	}

	//选择第一个元素作为基准
	kVal := sli[0]
	//初始化左右两个切片
	lfSli, rtSli := []T{}, []T{}

	//遍历除了标尺外的所有元素，按照大小关系放入左右两个切片内
	for i := 1; i < leng; i++ {
		if sli[i] < kVal {
			//放入左边切片
			lfSli = append(lfSli, sli[i])
		} else {
			//放入右边切片
			rtSli = append(rtSli, sli[i])
		}
	}

	//再分别对左边和右边的切片进行相同的排序处理方式递归调用这个函数
	lfSort, rtSort := []T{}, []T{}
	var eg errgroup.Group
	eg.Go(func() error {
		lfSort = Quick(lfSli)
		return nil
	})
	eg.Go(func() error {
		rtSort = Quick(rtSli)
		return nil
	})
	_ = eg.Wait()

	//合并
	return append(append(lfSort, kVal), rtSort...)
}
