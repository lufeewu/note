# 简介
golang 提供了 container 标准库，提供数据类型堆、双向链表及环形链表

## 模块
+ heap
+ list
+ ring

## heap 堆结构
提供了对任意类型（实现 heap.interface接口）的堆操作

+ type Interface
+ Init(h Interface)
+ Push(h Interface,x interface{})
+ Pop(h Interface)interface{}
+ Remove(h Interface,i int) interface{}
+ Fix(h Interface, i int)

## list 双向链表结构
container/list 标准库中实现了双向链表, 它是一个包含了 root 头结点信息的链表. root 是一个用于记录头结点、尾结点信息的结点。链表中的第一个结点的 pre 指针和最后一个结点的 next 指针指向 root，而 root 的 next 指针和 pre 指针则分别指向第一个结点和最后一个结点。root 结点不记录具体数据，也不是链表中的真实结点，它作为虚拟结点用于提升查询效率。
+ type Element struct
    - Next() *Element
    - Prev() *Element
+ type List struct
    - Init() *List
    - Len() int
    - Front() *Element
    - Back() *Element
    - PushFront(v interface{}) *Element
    - PushFrontList(other *List)
    - PushBack(v interface{}) *Element
    - PushBackList(other *List)
    - InsertBefore(v interface{}, mark *Element) *Element
    - InsertAfter(v interface{}, mark *Element) *Element
    - MoveToFront(e *Element)
    - MoveToBack(e *Element)
    - MoveBefore(e, mark *Element)
    - MoveAfter(e, mark *Element)
    - Remove(e *Element) interface{}

## ring 环形链表结构
+ type ring struct
    - New(n int) *Ring
    - Len() int
    - Next() *Ring
    - Prev() *Ring
    - Move(n int) *Ring
    - Link(s *Ring) *Ring
    - Unlink(n int) *Ring
    - Do(f func(interface{}))
