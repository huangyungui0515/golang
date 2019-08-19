package main

import (
	"fmt"
	//"strconv"
)

/* 阿里巴巴与四十大盗：背包问题
	阿里巴巴发现在强盗留下来的财宝，但他骑的毛驴运载能力有限。
	假设山洞中有 n 种宝物，每种宝物的重量，及价值都清楚
	比如 2 公斤黄金，价值是100块。2公斤白银 价格是80块.....
	阿里巴巴如何 选择，才能在承载能力下，得到价值最多的宝物
	解决思路：
	a：输入相应的财宝重量，及总价格，及毛驴的承载重量
	b：然后再求出每种财宝单位重量的 价值
	c：逐个汇总，再与船只承载重量比较
	d：如果大于承载重量，可敲碎宝物，再得到小碎片的财宝。直到财宝重量与 承载重量相等。
 */
type treasure = struct {
	weight   int     //财宝的重量
	worth    int     //财宝的价值
	perworth float32 //每个财宝的重量对应的价值
}

func sort(arr []treasure) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j+1].perworth > arr[j].perworth {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}
}

func main() {
	var treasureinfo []treasure
	var weight, worth int
	var temtreasure treasure
	var totalwieght int
	//var perworth float32
	fmt.Println("请输入相应的财宝重量 及价值, 用-1结束")
	for {
		fmt.Scan(&weight)
		if weight == -1 {
			break
		}
		fmt.Scan(&worth)
		temtreasure.weight = weight
		temtreasure.worth = worth
		//perworth, _ =float32(weight/ worth)// strconv.ParseFloat(fmt.Sprintf("%.000f", weight/worth), 32)
		temtreasure.perworth = float32(worth) / float32(weight)
		//将其加入到 treasureinfo 中
		treasureinfo = append(treasureinfo, temtreasure)
	}

	fmt.Println("请输入船支的承载重量")
	fmt.Scan(&totalwieght)
	sort(treasureinfo)
	fmt.Println(treasureinfo)
	sumweight := 0
	var sumprice float32 = 0.0
	for i := 0; i < len(treasureinfo); i++ {
		//如果当前重量 加上 下一财宝重量 小于 船只承载重量
		if (sumweight + treasureinfo[i].weight) <= totalwieght {
			sumweight = sumweight + treasureinfo[i].weight
			sumprice = sumprice + float32(treasureinfo[i].worth)
		} else {
			//如果大于说明超重，超重后的重量，可以将财宝敲碎取得最大的 价值
			sumprice = sumprice + float32(float32(totalwieght-sumweight)*treasureinfo[i].perworth)
			break
		}
	}
	//转称字符 %0.2f 将小数转换成2位小数
	fmt.Printf("可以装载价值为%.2f 的财宝", sumprice)
}
