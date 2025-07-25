# 简介
c/c++ 的基础知识.

## memcpy
memcpy 是 c/c++ 中常用的内存拷贝函数。它将 src 内存内的值拷贝到 dest 中。最简单的实现方式是一个字节一个字节的拷贝，但是这种方式性能较低。另外还需要处理内存重叠的问题。
拷贝性能: 对于 32 位的机器，地址总线是 32 位，一次可以复制 4 字节，效率更高。
内存重叠: 如果 dest 和 src 的内存存在重叠情况，需要处理。若是 dest 起始部分与 src 存在重叠，则需要从末端开始拷贝，以避免覆盖情况。

## 对象生存周期和资源管理(RALL, Resource Acquisition is Initialization) 
C++ 没有自动垃圾回收。RALL 资源获取即初始化，它的核心是把资源和对象的生命周期绑定，对象创建获取资源，对象销毁释放资源。在 RALL 指导下，C++ 把底层的资源管理问题提升到了对象生命周期管理的层次。
现代 C++ 通过声明堆栈上的对象尽可能避免使用堆内存。对于需要较大资源的堆栈，它应尽可能归对象所有。对象初始化时，获取它拥有的资源，最后在析构函数中释放资源。
在 C++ 11 中，标准库提供了智能指针。智能指针处理拥有的内存的分配和删除，无需显示的编写析构函数。

## 文件
在 cpp 工程中，包含两类文件，.cpp 文件和 .h 文件，其中 .cpp 文件被称作 c++ 源文件，里面主要放 c++ 的源代码，而 .h 文件则被称作 c++ 头文件。

## 源文件 .cpp
.cpp 文件被称作 c++ 源文件，里面放的主要是 cpp 的源代码。

## 头文件 .h 
cpp 的头文件 .h 的作用主要是被其它的 .cpp 文件包含进去，在 cpp 文件中，可以通过 #include 来关联头文件。本身不参与编译，但是实际上它们的内容在多个 .cpp 文件中得到了编译。头文件中应该只放变量和函数的声明，而不能放他们的定义。但有三个例外:
1. 头文件中可以写 const 对象的定义。
2. 头文件中可以写内联函数(inline)的定义。
3. 头文件中可以写类(class)的定义。

在编译的时候，并不会去寻找 .h 文件的实现，而是在 link 的时候去寻找。在 cpp 文件中引入 #include 实际上只是引入申明，使得编译可以通过。程序并不关心实现在哪里。源文件编译后生成目标文件(.o 或 .obj 文件)，在目标文件中，函数和变量作为一个符号。在 Makefile 中需要说明链接哪个 .o 或 .obj 文件，此时连接器会去这个 .o 或 .obj 文件中找到 cpp 文件中实现的函数，再把它们 build 到 Makefile 指定的可执行文件中。

.h 文件中能包含的内容:
1. 类成员数据的声明，但不能赋值。
2. 类静态数据成员的定义和赋值，但不建议，只声明即可。
3. 类的成员函数声明
4. 非类的成员函数的声明
5. 常熟的定义: 如 constint a = 5
6. 静态函数的定义
7. 类的内联函数的定义

.h 文件中不能包含的内容: 
1. 所有非静态变量(不是类的数据成员) 的声明
2. 默认命名空间的声明不放在头文件，using namespace std; 应放在 .cpp 文件中, 在 .h 文件中可以使用 std::string 等。

## 多文件编译
cpp 的多文件编译方式主要是借助编译、链接过程

## 右值引用
C++ 11 引入了右值引用, 用 && 表示, 主要用来实现移动语义和完美转发, 提高程序效率. 右值引用可以绑定到右值(临时对象或将亡值), 从而避免不必要的拷贝, 实现高效的对象移动. 尤其对于 std::vector、std::string 这样的大型对象.

## 常见题目
1. 类的缺省函数？
-   构造函数、析构函数、拷贝构造函数、赋值函数
2. 拷贝构造函数被调用？
- 类对象初始化另一个对象、函数的形参是类对象、函数的返回值是类的对象。
3. 什么时候重写拷贝构造函数？
- 涉及到动态存储分配空间的时候，要自己写拷贝构造函数，且需要深拷贝。
4. 友元、继承、公有成员函数等的作用？
- 友元可以用于在类外访问类的非公有成员。
5. 对象间如何实现数据的共享？
- 通过类的静态成员变量实现对象的数据共享。
6. 虚函数如何实现的？
- 虚函数表
7. delete 与 delete[] 的区别？
- delete只会调用一次析构函数，而delete[]会调用每一个成员的析构函数。
8. 堆和栈的区别？
- 栈区市由编译器自动分配释放，存放函数的参考值，局部变量等信息。堆区则一般是由程序员分配释放，或在程序结束时由操作系统回收，也可能造成内存泄漏。
9. 虚拟函数与普通成员函数的区别？内联函数和构造函数能否为虚拟函数？
- 虚拟函数是 virtual 关键字区分，有虚拟指针和虚函数表，普通成员函数没有。内联函数和构造函数不能为虚拟函数。
10. c++ 的多态怎么实现的?
11. 构造函数可以是虚函数么?(不能)
12. 左值和右值?
13. 如何避免野指针?

## 语法
c++ 的一些语法。

### switch case 语法
switch case, 一个 switch 语句可以测试一个变量等于多个值的情况，每个值为一个 case。使用示例如下:

        switch(expression){
            case constant-expression:
                statement(s);
                break; // 可选的
            case constant-expression:
                statement(s);
                break; // 可选的
            default : // 可选的
                statement(s);
        }

switch 语句有以下规则需要遵循:
- switch 语句中的每个 expression 必须是一个整型或枚举型，或者是一个包含转换函数可转换为整型或枚举类型的 class 类型，expression 不能是 string 等类型，但可以是 char 等。
- 可以有任意数量的 case 语句。每个 case 语句后有一个比较值和一个冒号。
- case 的 constant-expression 与 expression 变量必须是相同数据类型，且必须是常量或字面量。
- 执行变量为等于 case 的变量，case 后的语句会被执行，直到遇到 break 语句。如果 case 语句没有 break，则会继续执行后面的 case。
- 当遇到 break 语句时，switch 终止。
- switch 可以包含一个 default 语句，在所有 case 不为真时执行 default 语句。

### 引用传值语法
c++ 可以通过 & 声明引用类型，在赋值后相当于给变量取别名。

### const 用法
const 是 constant 的缩写，在 c++ 中用来修饰内置类型变量，自定义对象，成员函数，返回值，函数参数。const 用于告诉编译器某值是保持不变的，编译器会强制使用这个约束。

const 普通变量: 被定义为常量的变量，不能被再次赋值。
const 指针变量: const 修饰指针变量有三种情况，1. 修饰指针指向的内容，则内容为不可斌量。 2. 修饰指针，则指针不可变。 3. 修饰指针及指针指向的内容，则都不可变。

    const int *p = 8; // 修饰指向的内容
    int* const p1 = &a; // 修饰指针
    const int* const p2 = &a; // 修饰指针及指向的内容

const 参数传递和函数返回值: 
+ 修饰参数传递有三种情况, 1. 修饰值传递 2. 修饰指针 3. const & 传参
+ 修饰函数有三种情况, 1. 修饰内置类型返回值 2. 修饰自定义类型的作为返回值(不能被修改和被赋值) 3. 修饰返回的指针或引用

const 修饰成员函数: 可以防止成员函数修改被调用对象的值。


### 指针
dangling pointer: 当对象从内存中取消分配，但没有修改指向对象的指针的值，则出现 dangling pointer。
野指针: 指向内存被释放的内存或者没有访问权限的内存的指针。
悬空指针: 指向已经被删除对象的指针称为悬空指针。
智能指针: 指向动态分配(堆)对象指针的类。shared_ptr 便是 c++ 标准库中的智能指针，是为多个所有者可能必须管理对象在内存中的生命周期的方案设计的。


### 命名空间 namespace
cpp 的命名空间定义可以使用关键字 namespace，后面跟命名空间的名称。 如: 

    // 定义
    namespace namespace_name {
        // 代码声明
    }
    // 使用
    name::code // 变量或函数

using 指令: 可以通过使用 using namespace 指令，这样使用命名空间时就可以不用在前面加上命名空间的名称。
不连续的命名空间: 一个命名空间的代码可以定义在不同的部分中，可以分散在多个文件中。
嵌套命名空间: 命名空间可以嵌套，可以在一个命名空间中定义另一个命名空间。在使用是通过 :: 嵌套访问。

### 面向对象语法
c++ 的面向对象的特性主要是封装、多态、继承。通过 class 可以封装成员函数与数据，并且可以通过 public、private、protected 设置作用域，其中 class 默认的作用域是 private。
虚函数: C++ 的多态是通过虚函数实现的。在派生类中，可以重新定义虚函数，当使用指针或者基类的引用来引用派生的类对象时，可以执行该函数的派生类版本。
纯虚函数: 定义纯虚函数是为了实现一个接口，起到规范的作用，继承这个基类的派生类必须实现这个纯虚函数。
继承: 创建类时候，可以不重新编写成员函数和数据成员，而是通过继承基类的方式。继承包括私有继承、保护继承、公有继承，通常使用 public 公有继承。

派生类中如果实现了和基类一样的成员函数，则派生类将覆盖基类的成员函数，这取决于指针或对象的类型，在编译时候便确定了，即使是基类指针指向派生类对象，也将调用基类的成员函数实现。而虚函数的作用就是当基类指针指向派生类对象时，则使用的是派生类对象的成员函数，这样便完成了多态的特性，在程序运行时才通过对象的类型确定实现内容。


以下是代码例子，关于虚函数与普通成员函数在继承中的不同表现，虚函数展现出了多态特性。

    class Base{
        public:
        void echo(){cout<<"base"<<endl;}
        void basesay(){cout<<"base say"<<endl;}
        virtual void say();
    };

    void Base::say(){
        cout<<"base print"<<endl;
    }

    class Spec : public Base{
        public:
        void echo(){cout<<"spec"<<endl;}
        void say(){cout<<"spec print"<<endl;}
    };

    void testObjectOriented(){
        Base *p;
        Base *b = new(Base);
        Spec *s = new(Spec);
        b->echo(); // 基类指针指向基类
        s->echo(); // 派生类指针指向派生类, 派生类成员函数覆盖了基类实现
        s->basesay(); // 派生类指针指向派生类, 继承基类的成员函数

        // 普通成员函数与虚函数测试
        p = b;
        p->echo(); // 基类指针指向基类, 调用基类普通成员函数
        p->say(); // 基类指针指向基类, 调用基类虚成员函数

        p = s;
        p->echo(); // 基类指针指向派生类, 普通成员函数调用了基类的实现
        p->say();  // 基类指针指向派生类，虚成员函数调用了派生类的实现
    }

## ref() 和 cref() 
c++11 中引入了 std::ref 用于取某个变量的引用，为了解决一些传参问题。主要是在函数式编程如 std::bind 的时候，对参数是直接拷贝而不是引用的，需要借助 std::ref 传递引用，std::cref 则是 const 类型引用。


## Linux 库
linux 是由 c/c++ 编写的操作系统，linux 开放了很多 c/c++ 调用的库。

### 共享内存
在 linux 系统开发中，可以通过 shm_open 函数创建或者打开共享内存文件，shm_open 操作的文件一定是位于 tmpfs 文件系统内的。常见的 linux 系统中 tmpfs 文件系统存放在 /dev/shm 内。通过 fruncate 函数会将参数 fd 指定的文件大小改为参数 length 指定的大小。

int shm_open(const char *name, int oflag, mode_t mode): 用于创建或者打开共享内存文件。
int ftruncate(int fd, off_t length): 会将参数 fd 指定的文件大小改为参数 length 指定的大小，参数 fd 为已打开的文件描述词。
void *mmap64( void *addr, size_t len, int protection, int flags, int fildes, off64_t off): 用于建立内存映射文件，将某个文件的内容映射到内存中，对该内存区域的读写也就是对文件的读写。

### lambda 表达式
在 c++ 11 及更高版本中, 提供了 Lambda 表达式，是一种在被调用的位置或作为参数传递给函数的位置匿名函数对象(闭包)的简便用法.

## 内存管理
c++ 将程序分为 5 个区, 分别是堆、栈、自由存储区、全局/静态存续区、常量存续区.
- 堆(Heap): 从低地址向高地址增长. 容量大于栈、程序中动态分配的内存在此区域. 
- 栈: 从高地址向低地址增长, 由编译器自动管理分配. 局部变量、函数参数值、返回变量存储在此.
- 自由存储区(Text Segment): 存放可执行程序的机器码.
- 全局/静态存续区(BSS): 存放未初始化的全局和静态变量.
- 常量存续区(Data Segment): 存放已初始化的全局、静态变量、常量数据.

内存泄漏的几种情况:
- 类的构造函数和析构函数中 new 和 delete 没有配套
- 释放对象数组时没有使用 delete[] 而使用了 delete
- 没有将基类的析构函数定义为虚函数
- 没有正确清除嵌套的对象指针


## bazel 库管理
bazel 是与 Make、Maven、Gradle 类似的开源构建和测试工具. 支持构建 c++、java、Android、iOS 项目. Bazel 基于工作区 workspace 概念.
- workspace: 用于指定当前文件夹就是一个 Bazel 的工作区. WORKSPACE 文件总是存在于项目的根目录下.
- BUILD: BUILD 文件用于告诉 Bazel 怎么构建项目的不同部分。
- TARGET: 一个 target 指向一系列的源文件和依赖, 一个 target 也可以指向别的 target.
- 规则: 包括 cc_binary、cc_import、cc_library、cc_proto_library、cc_shared_library、fdo_prefetch_hints、fdo_profile、propeller_optimize、cc_test、cc_toolchain、cc_toolchain_suite
- cc_binary: 主要属性如下
    + name: 强制属性, target 的唯一名称.
    + srcs: 可选属性, 表示源文件.
    + deps: 要链接到二进制目标的其他库的列表.
    + includes: 要添加到编译行中的定义列表.
    + visibility: 规则可见性, 有 private 表示包私有, public 对所有包可用, 默认为 private.
- cc_library: 主要属性如下
    + name: target 的唯一名称。
    + deps: 要链接到二进制目标的其他库的列表.
    + hdrs: 此预编译库发布的将由源文件直接添加到相关规则中的头文件列表.

## coroutine 
c++ 20 开始正式支持了协程, 提供无栈协程. 新增了关键词及类型.
- 协程新关键词
    + co_wait: 会调用一个等待体对象(awaiter). 根据对象体内部接口, co_wait 将根据接口决定进行什么操作.
    + co_yield: 用来暂停协程并且绑定的 promise 里面放入一个值.
    + co_return: 往绑定的 promise 里面放入一个值, 同时结束这个协程.
- 协程新类型
    + coroutine_handle: 用于指代暂停或执行的协程. coroutine_handle 的每个特化均为字面类型.
    + coroutine_traits: 从协程的返回类型与形参类型确定承诺类型.
    + suspend_always: 空类, 用于指示 await 表达式始终暂停并且不产生值. 成员函数包括 await_ready、await_suspend、await_resume.
    + suspend_never: 空类, 用于指示 await 表达式绝不暂停并且不生产值. 成员函数包括 await_ready、await_suspend、await_resume.

## coredump
程序 core 是指应用程序无法保持正常 running 状态而发生的崩溃行为, 程序 core 时会生成相关的 core-dump 文件, 是程序崩溃时程序状态的数据备份. core 文件中包含内存、处理器、寄存器、程序计数器、栈指针等状态信息.

coredump 产生的几种可能情况:
- 内存访问越界: 数组越界、字符串结束符不正常、strcpy 等字符串操作函数读写越界
- 多线程程序使用线程不安全函数
- 多线程读写的数据未加锁保护
- 非法指针: 空指针、野指针、悬挂指针
- 堆栈溢出: 超大的局部变量导致堆栈溢出
- 内存超限
- 线程超限: 系统线程数超过限制
- 机器故障: 如 SIGBUG、SIGEMT、SIGIOT 等

coredump 定位方法:
- 通过 GDB 定位 core 
- 打印日志
- 定位代码行: 去编译优化、程序计数器 + addr2line、函数栈修复、无规律 core 栈 AddressSanitizer
- 定位 core 原因: 信号量确认、异常汇编命令、异常变量、优化变量、异常函数地址

### gdb 工具
GDB 是由 GND 开源组织发布、UNIX/LINUX 操作系统下的、基于命令行的、功能强大的程序调试工具. 它支持断点、单步执行、打印变量、观察变量、查看寄存器、查看堆栈等调试手段.

## 火焰图



## 参考
1. [C++ 继承](https://www.runoob.com/cplusplus/cpp-inheritance.html)
2. [C++ 虚函数和纯虚函数的区别](https://www.runoob.com/w3cnote/cpp-virtual-functions.html)
3. [C++11 的 std::ref 用法](https://murphypei.github.io/blog/2019/04/cpp-std-ref)
4. [C++ Lambda表达式基本用法](https://lellansin.wordpress.com/2014/01/02/c-lambda表达式基本用法/)
5. [C++ 20 协程 Coroutine之剖析](https://www.51cto.com/article/718493.html)
6. [C++20 新特性 协程 Coroutines(1)](https://zhuanlan.zhihu.com/p/349210290)
7. [渡劫 C++ 协程](https://github.com/bennyhuo/cppcoroutines)
8. [如何快速定位程序Core？](https://developer.baidu.com/article/detail.html?id=293651)
9. [bazel入门](https://htl2018.github.io/2020/04/05/bazel入门/)
10. [Bazel C/C++ 规则 ](https://bazel.build/reference/be/c-cpp?hl=zh-cn)