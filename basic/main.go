package main

import "fmt"

//什么是回调？
//先分层，有层才有回调一说
//一般：上层调下层
//回调：下层调上层
//上层：如一个本地客户端，假设有函数f1和f2
//下层：如一个第三方sdk，假设有函数f3
//假设代码执行流程为: f1->f3->f2（f1调f3，f3调f2）
//那么f2就被称为回调函数，因为它违反了一般上依赖下的原则，而是相反的被它的下层所依赖
//上下层如何连通呢？
//上层需要提供f2供下层f3调用，下层f3需要提供数据供上层f2使用，即下层把数据提供给上层，由上层决定如何利用数据
func main() {
	f1()
}

//////////////////////////////////////////////////////////////////////////////上层
func f1() {
	//do something high level needs
	f3()
	//do something else that high level needs
}

//f2为回调函数
func f2(name string) {
	fmt.Println("hello," + name) //上层拿到下层的数据自行处理
}

//////////////////////////////////////////////////////////////////////////////下层
//实际情况是下层并不知道上层的回调函数名叫f2，我将会在improve版本里进行优化
func f3() {
	//do something low level needs
	name := "guobin"
	f2(name) //下层给上层提供数据
	//do something else that low level needs
}
