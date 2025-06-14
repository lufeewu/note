# 简介
数据结构包括 slice、array、map、channel、interface 等

## slice

    a := make([]int, 3)
    b := []int{1, 2, 3}
    cap(a) = len(a)
    cap(b) = len(b)
    copy(b, a) // 拷贝两个切片长度较小的值


### 面试题

	for i := 0; i < 10; i++ {
		go func(){
			fmt.Print(i, " ")
		}
	}


	for k := 0; k < 10; k++ {
		go func() {
			j := k
			fmt.Print(j, " ")
		}()
	}

## map
map 的并发安全
map 底层的 bucket 拉链法解决冲突，哈希函数高低位方式定位
