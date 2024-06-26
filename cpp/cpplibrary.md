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

## openssl
openssl 是一个加密及 SSL/TLS 工具集.
- OpenSSL EVP: evp 函数提供高水平 openssl 加密算法.
- AES_GCM_256 加密算法: GCM (即 Galois/Counter Mode)模式本质上是 AES 模式加上 GMAC(Galois Message Authentication Code, 伽罗华消息认证码)进行哈希计算的一种组合模式. GCM 模式可以提供对消息的加密和完整性校验.
- Key: 对称秘钥, 长度可以为 128、192、256 bits, 用来加密明文的密码.
- IV(Initialisation Vector): 初始向量, 它的选取必须随机. 通常以明文的形式和密文一起传送, 作用和 MD5 的加盐类似, 防止同样的明文块始终加密成同样的密文块.
- ADD(Additional Authenticated Data): 附加身份验证数据. ADD 数据不需要加密, 通常以明文形式与密文一起传递给接收者.
- Mac tag(MAC 标签): 将确保数据在传输和存储过程中不会被意外更改或恶意篡改. 标签在解密操作期间使用, 以确保密文和 AAD 未被篡改. 加密时, tag 由明文、密钥 Key、IV、ADD 共同产生.

### openssl/evp.h 
通过 openssl/evp.h 提供的加密函数可以进行 aes 等加密操作. 以下是 evp 提供的 api 函数:
- EVP_CIPHER_CTX_new: 创建加密的 context.
- EVP_EncryptInit_ex: 设置加密 context、加密类型、密钥等.
- EVP_CIPHER_CTX_ctrl: 设置可用的各类加密模式.
- EVP_EncryptUpdate: 将输入的字节加密并写到输出中. 函数可以被多次调用用于加密连续块状的数据.
- EVP_EncryptFinal_ex: 对最终数据进行加密. 它使用标准块填充.
- EVP_CIPHER_CTX_free: 清楚密码上下文中的所有信息并释放与其关联的任何已分配内存, 包含 ctx 本身.


## 参考
1. [facebook/folly](github.com/facebook/folly)
2. [folly学习心得](https://blog.csdn.net/thanklife/article/details/80117429)
3. [稀疏索引 sparsehash](https://github.com/sparsehash/sparsehash)
4. [Google Sparse Hash](https://goog-sparsehash.sourceforge.net/)
5. [EVP_aes_256_gcm](https://www.openssl.org/docs/manmaster/man3/EVP_aes_256_gcm.html)
6. [aesgcm](https://github.com/majek/openssl/blob/master/demos/evp/aesgcm.c)
7. [AES_GCM_256加密算法](https://www.cnblogs.com/Galesaur-wcy/p/16843564.html)