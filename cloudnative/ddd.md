# 简介
领域驱动设计 (domain-driven design, 缩写 DDD) 是软件代码的结构及语言(类别名称、类方法、类变量)需符合业务领域中的习惯方法.

## 领域模型
领域模型是 DDD 的核心概念, 它是对业务领域的抽象和建模, 描述了业务领域的实体、值对象、聚合根、领域服务等元素, 以及它们之间的关系和行为.
- 充血模型: 是一种更接近领域模型的视线方法, 它强调在领域对象中包含业务逻辑和规则. 充血模型中的领域对象具有丰富的行为和状态, 能够更好地反映现实世界的复杂性和动态性. 通过将业务逻辑放在逻辑层, 可以提高代码的可维护性和可读性, 并减少层与层之间的耦合度.
- 实体: 具有唯一标识符, 并具有生命周期的业务对象.
- 值对象: 没有唯一标识符, 通过属性值定义.
- 聚合: 它是领域模型中相对独立的、可持久化的实体集合，用于表示业务中的整体.
- 聚合根: 聚合根是聚合的入口点和主要控制者, 是聚合的一部分, 负责维护聚合内的一致性和完整性.
- 聚合分析: 聚合分析不是 DDD 的正式术语. 可以理解为对聚合设计和实现进行评估的过程. 这个分析旨在确保聚合结构的合理、业务规则得到正确应用, 系统能够有效维护数据一致性和完整性.

## AI 

## 参考
1. [Eric Evans 提倡在领域驱动设计中实验大语言模型](https://www.infoq.cn/article/miepyu9zscchoyzfqe2h)
2. [DDD领域驱动设计理论](https://tech.dewu.com/article?id=113)
3. [DDD 领域驱动设计高级知识点](https://blog.csdn.net/samsung_samsung/article/details/135216734)
4. [DDD领域驱动最全详解(图文全面总结) -- 知识铺](https://index.zshipu.com/geek001/post/20240627/DDD领域驱动最全详解图文全面总结--知识铺/)
5. [7种内聚与7种耦合的类型及强弱关系](https://blog.csdn.net/Marion158/article/details/115892451)
6. [阿里一面：谈一下你对DDD的理解？2W字，帮你实现DDD自由](https://www.cnblogs.com/crazymakercircle/p/17130939.html)
7. [一文搞懂DDD的12个核心概念与2大建模方法](https://new.qq.com/rain/a/20240523A04JCJ00?suid=&media_id=)
8. [DDD领域驱动设计落地实践（十分钟看完，半小时落地）](https://www.cnblogs.com/dennyzhangdd/p/14376904.html#_label2_0)
9. [领域模型、贫血模型与充血模型：概念与实践](https://cloud.baidu.com/article/3167078)
10. [字节面试：请说一下DDD的流程，用电商系统为场景](https://www.cnblogs.com/crazymakercircle/p/17827728.html)
11. [07 | DDD 分层架构：有效降低层与层之间的依赖](https://zq99299.github.io/note-book2/ddd/02/02.html)
