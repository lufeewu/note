# 简介

## 源码
bufio 总计 3552 行代码，除去测试代码仅 1100 多行.
+ bufio.go
    - 实现了带缓存的 io 读写，实现了 io.Reader、io.Writer 接口对象. 
    - 同时提供了缓冲和一些文本 I/O 的帮助函数对象.
    - var ErrInvalidUnreadByte、ErrInvalidUnreadRune、ErrBufferFull、ErrNegativeCount
    - type Reader struct
    - type Writer struct
    - type ReadWriter struct
+ scan.go
    - 提供方便读取数据的接口，如从有换行符分隔的文本里读取一行.
    - type Scanner struct
       - func NewScanner(r io.Reader) *Scanner
       - func (s *Scanner) Split(split SplitFunc)
       - func (s *Scanner) Scan() bool
       - func (s *Scanner) Bytes() []byte
       - func (s *Scanner) Text() string
       - func (s *Scanner) Err() error

## 分析
bufio 实现了带缓存的 I/O 操作，它也是 golang 基础库之一. 它实现了 io 库的 Reader、Writer、ReaderWriter 接口, 其中方法 Read(p []byte)(n int,err error) 的实现如下

    func (b *Reader) Read(p []byte) (n int, err error) {
        n = len(p)
        if n == 0 {
            if b.Buffered() > 0 {
                return 0, nil
            }
            return 0, b.readErr()
        }
        if b.r == b.w {
            if b.err != nil {
                return 0, b.readErr()
            }
            if len(p) >= len(b.buf) {
                // Large read, empty buffer.
                // Read directly into p to avoid copy.
                n, b.err = b.rd.Read(p)
                if n < 0 {
                    panic(errNegativeRead)
                }
                if n > 0 {
                    b.lastByte = int(p[n-1])
                    b.lastRuneSize = -1
                }
                return n, b.readErr()
            }
            // One read.
            // Do not use b.fill, which will loop.
            b.r = 0
            b.w = 0
            n, b.err = b.rd.Read(b.buf)
            if n < 0 {
                panic(errNegativeRead)
            }
            if n == 0 {
                return 0, b.readErr()
            }
            b.w += n
        }

        // copy as much as we can
        n = copy(p, b.buf[b.r:b.w])
        b.r += n
        b.lastByte = int(b.buf[b.r-1])
        b.lastRuneSize = -1
        return n, nil
    }

该 Read 将从写入 buf 但还为读出的字节缓存读出到 byte 数组 p 中，并在 buf 都已经读取出时候，归零缓存空间（即 b.r = 0 、 b.w = 0), 并写入新数据到缓存 bufio 中(b.rd.Read(b.buf). 对于缓存 bufio 的写入，可以参看 Write 方法的实现.