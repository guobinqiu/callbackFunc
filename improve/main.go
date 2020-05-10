package main

import (
	"fmt"
)

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

	f3(f2) //显式声明回调函数f2。注意f2后面没有括号，因为我们需要传递的是一个函数而不是它的执行结果

	//或使用匿名函数形式
	f3(func(name string) {
		fmt.Println("hello," + name) //上层拿到下层的数据自行处理
	})

	//do something else that high level needs
}

//f2为回调函数
func f2(name string) {
	fmt.Println("hello," + name) //上层拿到下层的数据自行处理
}

//////////////////////////////////////////////////////////////////////////////下层
//实际情况是下层并不知道上层的回调函数名叫f2，我们可以对f3提供一个函数类型的参数让上层去实现它
func f3(f2 func(name string)) {
	//do something low level needs
	name := "guobin"
	f2(name) //下层给上层提供数据
	//do something else that low level needs
}

//////////////////////////////////////////////////////////////////////////////
//和我们熟悉的jquery经典回调做个对比吧
//printPageContent函数就相当于上面的f1
//success函数就相当于上面的f2
//ajax函数就相当于上面的f3
//function printPageContent() {
//	$.ajax({
//		url: "http://www.baidu.com",
//		success: function(data) {
//			console.log(data);
//		},
//	})
//}
