函数和延迟
====

函数
----

在Go语言中，使用关键词`func`来声明一个函数。一个基本的函数结构如下：

```golang
func (h *Handler) FuncName(param1 int, param2 string) (rt1 int) {
    // Func body
}
```

`(h *Handler)`部分为函数持有者，
函数可以没有持有者，在声明中直接省略这部分即可。

```golang
func FuncName(param1 int, param2 string) (rt1 int) {
    // Func body
}
```

`FuncName`为函数名，在Go语言中，函数作为头等对象，
函数名和参数一样可以在代码中进行传递。
特别的，函数还可以以匿名形式被定义和调用。
```golang
func MyFunc() {
    var func1 := func() {
        fmt.Println("here is func1")
    }

    func1() // here is func 1

    func() {
        fmt.Println("here is a clojure func")
    }() // here is a clojure func
}
```

`(param1 int, param2 string)` 是参数列表。
参数列表可以为空，但是要保留括号。
当连续多个参数类型相同时，前面的参数类型声明可以省略，只保留最后一个参数类型即可。

```golang
func FuncName(param1, param2 int, param3 string) (rt1 int) {
    // Func body
}
```

特别的，参数列表可以支持可变长参数形式。

```golang
func FuncName(params ...int) (rt1 int) {
    // params is []int
}
// call this func like:
// FuncName(param1, param2, param3)
```

`(rt1 int)`部分为函数返回值列表。
当函数没有返回值的时候，这部分内容可以省略。

```golang
func FuncName(param1 int, param2 string) {}
```

函数返回值可以只声明类型。

```golang
func FuncName(param1 int, param2 string) (int, error) {}

当返回时只有一项的时候，可以省略括号。

```golang
func FuncName(param1 int, param2 string) error {}
```

延迟
----

在实际编写代码的实践中，我们通常建议出错即结束。
为了方便用户在这种编程模式下还能够进行统一的函数结束处理，
Go语言中提供了延迟机制，通过`defer`关键词实现。

```golang
func Example_func() {
	fmt.Println("step0")
	defer func() {
		fmt.Println("defer1")
	}()
	fmt.Println("step1")
	defer func() {
		fmt.Println("defer2")
	}()
	fmt.Println("step2")
	if true {
		return
	}
	defer func() {
		fmt.Println("defer3")
	}()
	fmt.Println("step3")

	// Output:
	// step0
	// step1
	// step2
	// defer2
	// defer1
}
```

通过上面的例子可以看出，defer是在函数执行的时候以栈的形式被注册，
并且在函数结束后按出栈顺序执行。

特别的，如果函数不是正常return，而是发生了panic，
那么所有在函数中被注册的defer都不会被执行。
