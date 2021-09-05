# gorm 简介
gorm 是一个 golang ORM 库

## 知识点
+ 数据库连接
+ 自动迁移
    - 保持更新到最新
    - 仅创建表，缺少列和索引
    - gorm.Model
    - SingularTable 表名复数模式
    - DefaultTableNameHandler 默认表名更改
    - DeleteAt 字段，软删除
+ crud
    - NewRecord()
    - Create()
    - First()
    - Last()
    - Find()
    - Where()
    - Not()
    - Or()
    - Set()
    - Assign()
    - FirstOrCreate()
    - Attrs()
    - Select()
    - Order()
    - Limit()
    - Offset()
    - Count()
    - Group()
    - Join()
    - Pluck()
    - Scan()
    - Scopes()
    - Table()
    - Preload()
    - Update()
    - Updates()
    - UpdateColumn()
    - UpdateColumns()