package InterviewProgrammingCollection

import "fmt"

/*
实验室需要配制一种溶液。现在，研究员面前有n种该物质的溶液，每一种有无限多瓶，第i种的溶液体积为xi，里面含有yi单位的该物质。研究员每次可以选择一瓶溶液，将其倒入另外一瓶（假设瓶子的容量无限），即可以看作将两个瓶子内的溶液合并。此时合并的溶液体积和物质含量都等于之前两个瓶子内的之和。

特别地，如果瓶子A与B的溶液体积相同，那么A与B合并之后该物质的含量会产生化学反应，使得该物质含量增加X单位。

研究员的任务是配制溶液体积恰好等于C的，且尽量浓的溶液（即物质含量尽量多）。研究员想要知道物质含量最多是多少。



输入描述
第一行三个正整数n,X,C；

第二行n个正整数x1,x2,...,xn，中间用空格隔开；

第三行n个正整数y1,y2,...,yn，中间用空格隔开。

对于所有数据，1≤n,X,C,yi≤1000,1≤xi≤C

数据保证至少存在一种方案能够配制溶液体积恰好等于C的溶液。

输出描述
输出一行一个整数，表示答案。


样例输入
3 4 16
5 3 4
2 4 1


样例输出
29
*/

func Input1() {
	var n int
	var X int
	var C int
	fmt.Scan(&n, &X, &C)
	data := make([]water, n)
	for i := 0; i < n; i++ {
		data[i] = water{}
	}
	fmt.Println(data)

	for i := 0; i < n; i++ {
		tempx := 0
		tempy := 0
		fmt.Scan(&tempx, &tempy)
		data[i].x = tempx
		data[i].y = tempy

	}
	fmt.Println(SloveWater(data, C, X))
}

func SloveWater(data []water, C int, X int) int {
	var dp = make([]int, C+1)
	// 体积为j时对应的最大的容量
	for _, v := range data {
		dp[v.x] = max(dp[v.x], v.y)
	}
	for j := 0; j < C+1; j++ {
		for i := 0; i < len(data); i++ {
			if j-data[i].x >= 0 {
				if j-data[i].x == data[i].x {
					dp[j] = max(dp[j], dp[j-data[i].x]+data[i].y+X)
				} else {
					dp[j] = max(dp[j], dp[j-data[i].x]+data[i].y)
				}
			}
		}

	}
	fmt.Print(dp)
	return dp[C]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

type water struct {
	// 体积
	x int
	// 物质的含量
	y int
}
