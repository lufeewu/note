# 简介
golang 标准库之 compress 是几个常见类型解压缩格式的集合. 包括 bzip2、flate、gzip、lzw、zlib.

## 源码
5 个类型的包合计约 10000 行代码，除去测试代码约 5500 多行.
+ bzip2
    - 总计 1100 多行代码，除去测试代码约 870 行
    - 主要实现了哈夫曼(huffman) 和 MTF(move-to-front)
    - NewReader(r io.Reader) io.Reader
        - type reader struct
        - reader 实现了 Reader 接口. 
        - read(buf []byte)(int,error) 方法中使用 MTF 及 huffman
+ flate
    - 总计 5500 多行代码，除去测试代码约 2300 多行
    - flate 包实现了 deflate 无损压缩数据格式, 它使用 LZ77 算法与哈夫曼编码
    - [RFC 1951](http://tools.ietf.org/html/rfc1951)
    - dict_decoder.go 实现 LZ77，
    - gzip 和 zlib 实现了对基于 deflate 文件格式的访问
    - type Reader
        - NewReader(r io.Reader) io.ReadCloser
        - NewReaderDict(r io.Reader, dict []byte) io.ReaderCloser
    - type Writer
        - NewWriter(w io.Writer, lever int)(*Writer, error)
        - NewWriterDict(w io.Writer, level int, dict []byte) (*Writer, error)
        - Write(data []byte) (n int, err error)
        - Reset(dst io.Writer)
        - Flush() error
        - Close() error
+ gzip
    - 总计 1500 多行代码，除去测试代码约 980 多行
    - gzip 实现了 gizp 压缩, 参见 [RFC 1952](http://tools.ietf.org/html/rfc1952)
    - gzip 的基础是 deflate
    - type Header
    - type Reader
        - NewReader(r io.Reader) (*Reader, error)
    - type Writer
        - NewWriter(w io.Writer) *Writer
        - NewWriterLevel(w io.Writer, level int) (*Writer, error)
+ lzw
    - 总计 940 多行代码，除去测试代码约 530 多行
    - 实现了 Lempel-Ziv-Welch 数据压缩格式, 本包实现了用于 GIT、TIFF、PDF 的 lzw 压缩格式
    - “A Technique for High-Performance Data Compression” by T. A. Welch
+ zlib
    - 总计 810 多行代码，除去测试代码约 370 多行
    - 实现了读取时解压和写入是压缩, 使用了 compress/flate 

## ref
1. [golang pkg doc](https://studygolang.com/pkgdoc)