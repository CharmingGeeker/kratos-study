# kratos学习笔记六(kratos 服务发现 discovery)

除warden直连方式外，kratos有另一个服务发现sdk : [discovery](https://github.com/bilibili/discovery)

discovery 可以先简单理解为一个http服务、它最简单的发现过程可能是这样的:

> 1、service 向discovery 服务注册 appid
> 2、client 通过 appid 从discovery 查询 service 的addr

还包含了很多功能在里面的，例如**服务自发现**、**负载均衡**等



[blog地址](https://www.cnblogs.com/ailumiyana/p/12188958.html)
