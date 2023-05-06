# 简介
一些 cpp 的库.

## folly
folly(Facebook Opensource Library) 是 facebook 基于 c++ 14 开源的 c++ 库, 作为对 C++ 标准库的补充. 下面是一些 folly 提供的库:
- Arena: 内存管理库.
- AtomicHashMap: 高性能哈希表, 使用几乎无锁实现.
- Baton: 常用作线程同步、等待、通知的标识符号.
- Benchmark: 提供了简单的框架用于写入和执行基准测试.
- Bits: 位处理的组件, 针对速度进行了优化.
- ConcurrentSkipList: 实现了多线程环境下的 skiplist. 
- Conv: 数据转换例程, 针对速度和安全进行了优化.
- DiscriminatedPtr: 类似 boost:variant, 局限于指针
- Dynamic: 为 c++ 提供运行时动态类型.
- FBString: 对标准库的 string 性能优化的版本.
- FBVector: 对标准库的 vector 性能优化的版本.
- File: 用于文件对象.
- Fingerprint: 用于计算 Rabin 指纹.
- Function: 是一个多态函数封装, 类似于 std::function 但多了一些特性.
- Histogram: histogrm.h 定义了一个直方图类, 用于追踪大规模流数据.
- json: 用于序列化和反序列化 json 数据.
- Likely: 当比编译器更了解分支情况时.
- Malloc: 用于内存分配, 提供比 jemalloc 更智能的函数.
- Memory: 与 Malloc 一起用于内存分配.
- MPMCQueue: 是一个高性能有界并发队列, 支持多生产者、多消费者, 并支持阻塞.
- PackedSyncPtr: 高度专业化的数据结构, 包含一个指针、1 bit 自旋锁和 15 bit 填充位.
- Poly: 是一个类模板, 可以使定义类型擦除多态对象包装器相对容易。
- Preprocessor: 预处理设施.
- Random: 用于使用时间和 pid 产生随机数种子.
- Range: 类似 boost 的随机访问数据包装类.
- TimeoutQueue: 定时器队列.
- ThreadLocal: 改进的本地存储线程, 用于存储非内置类型.
- ThreadCachedInt: 使用线程缓存的高性能原子增量.
- TokenBucket: 线程安全的令牌桶实现.

## spasehash
google 开源的[稀疏索引 sparsehash](https://github.com/sparsehash/sparsehash). 是一个非常节省内存的 hash_map 实现, 每个 entry 只有 2 bit 的开销.

## 参考
1. [facebook/folly](github.com/facebook/folly)
2. [folly学习心得](https://blog.csdn.net/thanklife/article/details/80117429)
3. [稀疏索引 sparsehash](https://github.com/sparsehash/sparsehash)
4. [Google Sparse Hash](https://goog-sparsehash.sourceforge.net/)