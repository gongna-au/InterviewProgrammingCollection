package InterviewProgrammingCollection

/*
输入描述
第一行有一个正整数n(1<=n<=5000)，代表小明有几个球。

第二行有n个范围在[-1000000,1000000]内的整数，第i个代表编号为1到n的球上写的数。

第三行是一个长度为n的仅由`R`和`B`构成的字符串，第i个字母代表编号为i的球是红色(R)或蓝色(B)

第四行有一个正整数m(1<=m<=100000)，代表小明进行了几次操作。

第五行有m个递增的正整数，第i个代表小明进行的第i次操作时间点。每个时间点小明只会进行至多一次操作。时间点的范围在[1,1000000000]内。

第六行有m个整数，第i个代表小明进行的操作。0为询问当前时间点袋中球上的数字之和，正数x代表放入了编号为x的球，负数-x代表取出了编号为x的球。

最开始袋子是空的。 你可以认为球上的数字变化均发生在时间点之前，而每次操作均发生在时间点之后。输入保证操作合法。

输出描述
设小明进行了k次询问。你需要在一行中先输出k，然后输出k个数，第i个代表第i次询问的答案。

题目保证小明进行过至少一次询问。


样例输入
3
-5 4 6
RBR
9
1 2 3 4 5 6 8 13 14
1 3 0 2 -3 0 -1 0 -2
样例输出
3 4 2 -5
*/
