# interface 简介
golang 不是一个基于 class 的语言. golang 提供 interface 类型可以支持一些面向对象的特性. 任何实现接口定义方法的类都可以实例化该接口, 接口和实现类之间没有依赖关系. 通过组合，可以实现面向对象的继承. 通过重新实现方法可以实现面向对象的覆盖. 至于多态，则可以通过任意实现接口的指针的实例化完成.

## type
golang 提供丰富的类型系统， interface 也是类型的一种 type ... interface 可以知道.

## empty interface
golang 允许不带任何方法的 interface, 这种 interface 叫 empty interface. Go 语言采用 runtime.iface 表示带方法的接口, runtime.eface 表示不带任何方法的空接口. 空接口 interface {} 类型的变量含两个指针, 一个指针指向值的类型, 另一个指针指向实际的值. 当两个指针均为 0 时, 空接口 interface {} 才为 nil.

## ref
1. https://sanyuesha.com/2017/07/22/how-to-understand-go-interface/