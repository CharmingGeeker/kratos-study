# kratos学习笔记--构建单独的http或者grpc demo项目
## grpc
1. 新建grpc项目   
```shell
kratos new rpcdemo --grpc
```

2. run 项目  

   ```shell
   kratos run
   ```

   `warden` ：简单了解了下kratos的grpc框架 不是直接使用的google的grpc，类比http也是对grpc接口做了定制包装而成的。不改gRPC源码，基于接口进行包装集成trace、log、prom等组件打通自有服务注册发现系统discovery实现更平滑可靠的负载均衡算法。

Client.go

```go
package main

import (
   "fmt"

   pb "kratos-demo/api"

   "golang.org/x/net/context"
   "google.golang.org/grpc"
)

const (
   Address = "127.0.0.1:50052"
)

func main() {

   conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
   if err != nil {
      fmt.Println(err)
   }
   defer conn.Close()

   c := pb.NewDemoClient(conn)

   req := new(pb.HelloReq)
   req.Name = "kratos grpc"
   r, err := c.SayHelloURL(context.Background(), req)
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Println(r.Content)

}
```

## http

前面内容都是基于kratos 原本demo的使用， 现在我们自己随便定义一个liveroom.proto的bm服务, 看看需要改动哪些，跑起微服务。

1. 创建一个只生成bm代码的项目liveroom

```shell
kratos new liveroom -d C:\项目路径 --http
```

2. 但这只是前两章的demo项目，接着删掉api 路径下的pb.go和bm.go，定义自己的api.proto    

3. `go generate` 生成新的go接口。   

4. client.go 的newclient()接口也需要重定义。    

5. 接着service 层 业务逻辑层 重新定义我们的接口实现    

6. `go generate` 依赖注入层修改wire.go,重新生成静态分析文件wire_gen.go。   
7. `kratos run`
  
   




[blog地址](https://www.cnblogs.com/ailumiyana/p/12114703.html)