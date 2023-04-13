# 简介
sql 语句学习, 普通数据库操作、大数据分析.

## 语法
- join: 通过 join 操作可以将两个或多个表的行结合. join 语句包括 left join、right join、inner join、full join 等.
- 关系型数据库范式: 范式是设计数据结构过程要遵循的规则和指导方法. 目前有 8 种范式, 1NF、2NF、3NF、BCNF、4NF、5NF、DKNF、6NF. 常用的是第一、二、三范式.
    + 第一范式: 列的原子性.
    + 第二范式: 每个表必须有主键, 其它数据元素与主关键字一一对应. 每一个非主属性完全依赖于主键.
    + 第三范式: 在满足第二范式的基础上, 非主键列必须直接依赖于主键, 不能存在传递依赖.

## 实践
sql 操作实践.

- **insert into**: select into 用于从源表查询数据集结果


## 参考
1. [CSDN - 数据库（第一范式，第二范式，第三范式](https://blog.csdn.net/Dream_angel_Z/article/details/45175621)