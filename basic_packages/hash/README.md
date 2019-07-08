# 简介
标准库 hash 提供了 hash 函数的接口

## 源码
总计 2890 多行，除去测试代码 1100 多行
+ type Hash interface 
+ type Hash32 interface 
+ type Hash64 interface
+ hash/adler32 
    - Adler-32校验和算法
    - func Checksum(data []byte) uint32
    - func New() hash.Hash32
+ hash/crc32
    - 32 bits 循环冗余校验算法
    - func Checksum(data []byte, tab *Table) uint32
    - func ChecksumIEEE(data []byte) uint32
+ hash/crc64
    - func Checksum(data []byte, tab *Table) uint64
+ hash/fnv
    - 实现了 FNV-1 和 FNV-1a 非加密 hash 函数
    - func New32() hash.Hash32
    - func New32a() hash.Hash32
    - func New64() hash.Hash64
    - func New64a() hash.Hash64

## ref
1. [Adler-32校验和算法, RFC 1950](https://tools.ietf.org/html/rfc1950)
2. [CRC-32 32位循环冗余校验](http://en.wikipedia.org/wiki/Cyclic_redundancy_check)
3. [CRC-64 64位循环冗余校验](http://en.wikipedia.org/wiki/Cyclic_redundancy_check)
4. [FNV-1、FNV-1a（非加密hash函数）](http://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function)