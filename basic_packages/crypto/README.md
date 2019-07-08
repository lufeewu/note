# 简介
crypto 标准库提供常用的密码（算法）及常量. 如 AES、DES、RSA、SHA、HMAC、MD5 等，实现了椭圆曲线数字签名等算法，提供 rand 可以用于加解密的随机数生成器.

## 源码
代码量总计 60000 多行，测试代码 36000 多行. 
+ type PublicKey interface 
+ type PrivateKey interface 
+ type Hash 
    - func (h Hash) Available() bool 
    - func (h Hash) Size() int 
    - func (h Hash) New() hash.Hash 
+ func RegisterHash(h Hash, f func() hash.Hash) 
+ type Signer interface 
+ aes 
+ cipher 
+ des 
+ dsa 
+ ecdsa 
+ ed25519 
+ elliptic 
+ hmac 
+ internal 
+ md5 
+ rand 
    - func Int(rand io.Reader, max *big.Int) (n *big.Int, err error) 
    - 返回一个在 [0, max] 区间服从均匀分布的随机值，如果 max <= 0，则会 panic
    - func Prime(rand io.Reader, bits int) (p *big.Int, err error)
    - 返回一个具有指定字位数的数字，该数字具有很高可能性是质数 
+ rc4 
+ rsa 
+ sha1 
+ sha256 
+ sha512 
+ subtle 
+ tls 
+ x509 

## ref
1. [Understanding Cryptography:A Textbook for Students and Practitioners](http://swarm.cs.pub.ro/~mbarbulescu/cripto/Understanding%20Cryptography%20by%20Christof%20Paar%20.pdf)
2. 译书: 深入浅出密码学