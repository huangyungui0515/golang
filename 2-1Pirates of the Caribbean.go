package main

import "fmt"

/* 加勒比海盗船，最优装载问题：
	加勒比海盗抢到一批破古董，每个古董有一定的重量。但船也有一定的承载能力。
	比如有10个古董，但船只能承载 30千克的重量。
	那么最多能装几个古董
   解决思路：
	a：输入相应的古董重量，及古董名称，及船支的承载重量
	b：然后将其按重量排序
	c：逐个汇总，再与船只承载重量比较。就得到最多能装几个古董
 */

type gudong struct {
	weight int
	name   string
}

func sort(arr []gudong) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j+1].weight < arr[j].weight {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}
}
func main() {
	var arr []gudong
	var gdinfo gudong
	var num, weight int
	sum := 0
	icount := 0
	sname := ""
	fmt.Println("请输入每个古董的重量，用空格分开，以输入-1结束")
	for {
		fmt.Scan(&num)
		gdinfo.weight = num
		fmt.Scan(&sname)
		gdinfo.name = sname
		if num == -1 {
			break
		} else {
			arr = append(arr, gdinfo)
		}
	}
	fmt.Println("请输入船所能随的重量:")
	fmt.Scan(&weight)
	//将整个切片排序
	sort(arr)
	//fmt.Printf("%v", arr)
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i].weight
		if (sum <= weight) {
			icount = icount + 1
		}
	}
	fmt.Printf("计算后总共可以装%d个古董,分别是 \n", icount)
	for i := 0; i < icount; i++ {
		fmt.Println(i, ":", arr[i].name)
	}
	//fmt.Printf("计算后总共可以装%d 物品", icount)
}
