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
- **insert into**: select into 用于从源表查询数据集结果.

        SELECT table_column1, table_column2, table_column3...
        INTO new_table_name [IN another_database]
        FROM table_name;

- **limit**: 用于限定返回的行数. 当后面接两个参数如 limit n,m 时, 表示从第 n 条记录开始返回 m 条记录.
- **offset**: 用于指定跳过的记录条数, 单独使用可能不起作用



## inner join
利用 inner join , 可以在通用字段中匹配时从两个表合并记录

    select 
        a.feedid, a.flag as v0, a.duration as v0_duration, 
        b.flag as v1,b.duration as v1_duration 
    from 
        test a inner join test b on a.feedid = b.feedid 
    where a.flag = 'V0' and b.flag = 'V1';


## 参考
1. [CSDN - 数据库（第一范式，第二范式，第三范式](https://blog.csdn.net/Dream_angel_Z/article/details/45175621)
2. [SELECT INTO 敘述句](https://www.fooish.com/sql/select-into.html)