---
name: gmsec-dev
description: gmsec 微服务开发规范与流程指南。当用户在本项目中新建微服务、编写/修改 proto 接口、自定义 REST 接口名与路由（@Router 注解、ginrpc）、执行 make gen / make orm / make new / make master / make windows / make linux / make mac、开发业务逻辑、接入数据库、排查 protoc 或 gormt 相关问题时，必须加载本 skill 并严格遵循其中的命名规范、目录结构与命令流程。
---

# gmsec 微服务开发指南（AI 助手必读）

gmsec 是基于 **gRPC + Gin** 的 Go 微服务集成框架：一套 proto 定义同时生成 gRPC 接口和 RESTful 接口（ginrpc 自动绑定参数），grpc 与 gin 共用端口，内置 gormt 数据库代码生成、服务注册发现（默认 DNS/mdns，可切换 etcd）、Jaeger 链路追踪、Prometheus 监控、Swagger/Markdown 文档自动导出。

**仓库根目录**：`<repo>/`（下文所有相对路径以此为基准）
- `Makefile` —— 仓库级命令（新建服务、proto 环境安装）
- `apidoc/` —— proto 源文件与生成的 rpc 代码（跨服务共享）
- `example/` —— 服务模板项目，新服务均由它复制而来

**执行原则**：所有 make 命令必须在正确目录下执行——`make new service=xxx` 在**仓库根目录**执行；`make gen`、`make orm`、`make windows/linux/mac` 在**对应服务目录**（如 `example/`、`xxx/`）下执行。

---

## 1. 新建微服务

### 命令

```bash
# 在仓库根目录执行
make new service=[微服务名字]
# 例：
make new service=user
```

### 命令做了什么

1. 复制 `./example/` 整个模板到 `./[service]/`（若目录已存在则跳过，不会覆盖）。
2. 全量字符串替换 `example` → `[service]`，涉及：
   - `./[service]/go.mod`（module 名 → `[service]`）
   - `./[service]/Makefile`（`NAME := [service]`）
   - `./[service]/generate/proto_makefile`（proto 编译目标）
   - `./[service]/conf/config.yml`（`service_name` 等）
   - 所有 `*.go` 文件的 import 路径（`example/rpc/example` → `[service]/rpc/[service]`）
3. 复制 `./apidoc/proto/example/` → `./apidoc/proto/[service]/`，并将其中 proto 的 `package`、`service` 名替换为 `[service]`。

### 命名规范

- 服务名：**全小写、单个英文单词**（如 `user`、`order`、`pay`），不用驼峰、不用下划线。
- 服务名会同时成为：Go module 名、proto 的 `package`/`service` 名、生成代码包名、服务注册名 `gmsec.srv.[service]`、REST 路由前缀 `/[service]/api/v1`。一旦确定不要随意改动。
- gRPC 服务注册名固定格式：`gmsec.srv.[service]`（见 `main.go` 中 `micro.WithName`）。

### 新服务目录结构（继承自 example 模板）

```
[service]/
├── main.go                    # 程序入口：服务初始化、路由挂载、插件启动
├── main_test.go               # 集成测试模板（TestServer/TestClient）
├── gogenerate.go              # go:generate 入口，触发 proto 编译
├── Makefile                   # 服务级命令：gen/orm/build/windows/linux/mac/clear
├── go.mod
├── conf/
│   └── config.yml             # 服务配置（yaml，可被 consul 覆盖）
├── generate/
│   └── proto_makefile         # protoc 编译脚本（make gen 的实际执行体）
├── rpc/                       # proto 生成的代码（linux/mac 下为指向 ../rpc 的软链）
├── docs/                      # 自动生成的 API 文档（markdown/ + swagger/）
├── tools/
│   ├── gormt[.exe]            # ORM 代码生成工具
│   └── config.yml             # gormt 配置（数据库连接等）
└── internal/
    ├── api/                   # ginrpc 上下文定义（api.NewAPIFunc 等）
    ├── config/                # 配置结构体与加载逻辑
    ├── core/                  # dao/连接等核心封装
    ├── routers/
    │   ├── api_root.go        # 服务注册与路由挂载入口（OnInitRoot）
    │   └── gen_router.go      # 自动生成的路由注册
    ├── service/               # 业务逻辑实现（按服务名分目录）
    └── model/                 # gormt 生成的 ORM 模型（make orm 输出）
```

### 新建服务后的必做事项

1. `cd [service]` 进入服务目录，先执行 `make gen` 重新生成本服务的 proto 代码（模板里带的是 example 的产物，替换后需重编译确认无误）。
2. 检查 `conf/config.yml`，修改端口（`port`）等配置，避免与已有服务端口冲突。
3. 执行 `make run`（或 `go build`）验证服务能正常启动。

---

## 2. 编写 API 接口（proto 文件）

### 文件位置

proto 源文件统一放在 **`apidoc/proto/[service]/`** 目录，不要放到服务目录里。公共消息（如通用响应结构）放在 `apidoc/proto/common/`。

### 编写规范

以 `apidoc/proto/example/hello.proto` 为标准模板：

```protobuf
syntax = "proto3"; // 指定proto版本（固定 proto3）

package example;  // 包名 = 服务名（make new 已自动替换）

import "common/common_public.proto"; // 引用公共消息（按需）

option go_package = "rpc/example";   // 固定格式：rpc/[服务名]

// 定义example服务（注释必写，service 上方一行）
service example {
    // 定义SayHello方法（注释必写，会导出到 API 文档）
    rpc SayHello(HelloRequest) returns (HelloReply) {}
}

// HelloRequest 请求结构（注释格式：`结构体英文名 + 空格 + 中文说明`）
message HelloRequest {
    string name = 1; // 名字（每个字段必须有行尾注释，说明字段用途）
}

// HelloReply 响应结构
message HelloReply {
    string message = 1; // 消息
    common.ResultResp result = 2; // 多proto文件引用样例
}
```

**硬性规则：**

1. `package` 与 `service` 名必须等于服务名（框架按此约定生成包名和注册函数）。
2. `option go_package = "rpc/[服务名]"`，不要写成别的路径。
3. **注释规范**（注释会被提取生成 swagger/markdown 文档）：
   - service 定义上方：`// 定义XX服务` 或直接写服务用途说明。
   - rpc 方法上方：一行中文方法说明。
   - message 上方：`// [消息名] [中文说明]`，消息名与实际英文标识一致。
   - 每个 message 字段行尾必须有 `//` 注释，写明字段含义。
4. 跨 proto 引用公共消息走 `import "common/common_public.proto"`，公共消息一律定义在 `apidoc/proto/common/`，**不要在各服务的 proto 里重复定义通用结构**。
5. 字段编号：新增字段只增不改不复用；已上线接口不允许修改已有字段编号或类型（protobuf 向后兼容原则）。
6. 一个服务 proto 可拆多个文件（`generate/proto_makefile` 会遍历目录下所有 `.proto`），建议按业务域拆分，文件名小写。

---

## 3. 生成 API 代码

### 命令

```bash
# 先进入服务目录
cd example        # 或 cd [service]
make gen
```

`make gen` 实际执行 `go generate` → `make -s -f generate/proto_makefile`，后者完成：

1. （Linux/mac）创建 `rpc` 软链接指向 `../rpc`，保证各服务共享 `apidoc/rpc` 下的公共代码。
2. 用 `protoc --gmsec_out=plugins=gmsec:./` 编译 `apidoc/proto/[service]/*.proto`。
3. 自动去掉生成代码中的 json `,omitempty` 标记（便于前端拿到完整字段）。
4. 将生成产物同步拷贝到 `../apidoc/rpc/[service]/`（供仓库级共享）。
5. 同时编译 `apidoc/proto/common/*.proto` 到 `apidoc/rpc/common/`。

### 生成产物及用途

| 产物 | 位置 | 用途 |
|---|---|---|
| `[name].pb.go` | `rpc/[service]/` | protobuf 标准生成代码：消息结构体（请求/响应）、序列化 |
| `[name].pb.gmsec.go` | `rpc/[service]/` | 框架扩展代码：gRPC 服务端注册函数（如 `proto.RegisterExampleServer`）、客户端获取函数（`proto.GetExampleClient()`）、服务名函数（`proto.GetExampleName()`） |
| `apidoc/rpc/common/*.pb.go` | `apidoc/rpc/common/` | 公共消息结构，各服务共享 |
| `docs/markdown/*.md`、`docs/swagger/swagger.json` | 服务目录 `docs/` | 由代码注释自动导出的 API 文档 |

**使用时机**：每次新增/修改 proto 文件后必须执行；生成文件**禁止手工修改**（会被下次生成覆盖）。

---

## 4. 开发业务逻辑

### 4.1 实现接口

业务实现放在 `internal/service/[业务名]/[业务名].go`。写法以 `internal/service/hello/hello.go` 为模板：

```go
package hello

import (
    "context"
    proto "[service]/rpc/[service]" // 引用生成的代码

    "github.com/xxjwxc/public/tools"
    "google.golang.org/grpc/metadata"
)

type Hello struct{} // 业务对象，方法集需覆盖 proto 中 service 定义的所有 rpc

// SayHello 方法签名必须与 proto 定义一一对应：
// func (o *对象) 方法名(ctx context.Context, req *proto.XxxRequest) (*proto.XxxReply, error)
func (h *Hello) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
    md, ok := metadata.FromIncomingContext(ctx) // 可从 ctx 取 gRPC metadata
    ...
    return &proto.HelloReply{
        Message: tools.GetRandomString(8),
    }, nil // 返回 error 时框架自动包装统一响应（见 api_root.go 的 GinAfter）
}
```

**规则：**
- 一个 proto `service` 对应一个业务对象（空 struct），对象里的**每个 rpc 方法都要有同名方法实现**，否则注册不完整。
- 方法签名固定 `(ctx context.Context, req *proto.XxxRequest) (*proto.XxxReply, error)`，参数返回类型与 proto 消息一一对应。
- 同一对象同时被注册为 gRPC 服务端和 gin restful 接口（ginrpc 自动按「对象名.方法名」映射路由，如 `/[service]/api/v1/hello.say_hello`，详见 4.5），**业务代码无需感知调用来源**。
- 如需取 gin 原始上下文：`c := api.Width_context(ctx)`（见 hello.go 注释示例）。

### 4.2 挂载注册到框架

新对象必须在 `internal/routers/api_root.go` 的 `OnInitRoot` 中注册（该函数在 `main.go` 启动时被调用）：

```go
// OnInitRoot 初始化
func OnInitRoot(s server.Server, router gin.IRoutes, objs ...interface{}) {
    var args []interface{}
    h := new(hello.Hello)
    args = append(args, h)                        // 1. 加入 gin 注册列表（restful 自动绑定）
    proto.RegisterExampleServer(s, h)            // 2. 注册为 gRPC 服务端（Register + 服务名大驼峰 + Server）
    args = append(args, objs...)
    OnInitRouter(router, args...)
}
```

新增业务对象的两步：`args = append(args, new(xxx.Xxx))` + `proto.RegisterXxxServer(s, 实例)`。

### 4.3 服务初始化流程（main.go，一般无需改动）

`main.go` 的 `CallBack()` 完成整条初始化链，理解顺序有助于排查启动问题：

1. swagger host/basePath 设置；
2. `tracer.WithTracer(...)` Jaeger 链路追踪初始化；
3. `registry.NewDNSNamingRegistry()` 服务注册方式（默认 DNS/mdns；etcd 见注释示例）；
4. `micro.NewService(micro.WithName("gmsec.srv.[service]"), ...)` 创建 gRPC 服务；
5. `gin.Default()` + 中间件（Cors、UseJager）+ 路由组 `/[service]/api/v1`；
6. `routers.OnInitRoot(service.Server(), v1)` —— **挂载业务对象的唯一入口**；
7. `plugin.Run(plugin.WithMicro(service), plugin.WithGin(router), plugin.WithAddr(":"+port))` 同时启动 gRPC 与 gin（共用端口），阻塞等待。

启动模式：`is_dev: true` 或 `./[service] debug` 直接前台运行；否则以服务方式托管（`server.On(...).Start`，支持安装为系统服务）。

### 4.4 依赖注入说明

框架采用**显式构造 + 统一注册**，不用额外 DI 框架：业务对象在 `OnInitRoot` 中 `new` 出来统一挂载；对象自身需要的外部依赖（数据库、redis 等）在包内 `init()` 或首次使用时通过 `internal/core` 封装懒加载获取（连接配置来自 `internal/config`）。

### 4.5 自定义 REST 接口名与路由（@Router 注解）

REST 路径由 [ginrpc](https://github.com/xxjwxc/ginrpc) 生成，**默认规则无需任何注解**，需要自定义时在 Go 业务方法上加 `@Router` 注释即可。gRPC 接口名不受影响（始终等于 proto 定义）。

#### 默认规则（不写注解时）

| 项 | 规则 | 示例 |
|---|---|---|
| 路径 | `对象名.方法名`，均转为**下划线小写**（`mybigcamel.UnMarshal`） | `Hello.SayHello` → `/hello.say_hello` |
| 方法 | 有 req 参数（ctx+req 两参）→ `POST`；只有 ctx（一参）→ `ANY` | SayHello 两参 → `POST` |
| 前缀 | 由 `main.go` 的 `router.Group("/[service]/api/v1")` 决定 | 完整路径 `/example/api/v1/hello.say_hello` |
| 大驼峰模式 | `ginrpc.WithBigCamel(true)` 时路径改为 `Hello.SayHello`（默认 false 为 web 风格） | — |

#### 注解语法

```go
// SayHello 打个招呼（这行注释会作为接口描述导出到 swagger/markdown 文档）
// @Router /hello/say [post,get]
func (h *Hello) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
    ...
}
```

- 格式：`// @Router <路径> [<方法1>,<方法2>]`，路径**相对路由组前缀**，最终为 `/example/api/v1/hello/say`。
- 方法列表可写 `post,get,put,delete` 等；**省略 `[...]` 时默认 `get`**。
- **忽略 HTTP 映射**：`// @Router - [-]` —— 该方法只保留 gRPC 能力，不暴露 REST（内部接口常用）。
- 一个方法可写**多条** `@Router` 注释注册多个路径/方法组合。
- 第三方网关标记：`// @Router /x [POST][router aaa,bbb,ccc]`（`[名字 数据]` 形式，配合网关插件使用）。
- 注解必须`//`紧贴函数上方（`f.Doc` 注释组内），方法体内部注释无效。

#### 生效机制与注意事项

1. 注册时 ginrpc 用 AST 解析源码 → 生成 `internal/routers/gen_router.go` 与 `conf/gen_router.data`（`api_root.go` 已配置 `ginrpc.WithOutPath("internal/routers")`）。
2. 入口需 import（或 `_ ` 匿名 import）`[service]/internal/routers` 包，`gen_router.go` 的 `init()` 才会执行；example/main.go 已直接 import，无需额外处理。
3. **改了 `@Router` 注解后重新编译启动即可**（不走 make gen，proto 未变）；生成的两个文件不要手改，会被覆盖。
4. 路径不要与已有接口重复，否则注册冲突导致启动失败。
5. 只有 REST 侧命名变化：gRPC 方法名/客户端 stub 仍按 proto 定义，服务间调用不受影响。

---

## 5. 数据库操作（make orm）

### 命令

```bash
# 在服务目录执行
make orm        # 实际执行：./tools/gormt -o internal/model
```

[gormt](https://github.com/xxjwxc/gormt) 读取数据库表结构，自动生成 gorm 模型与 CRUD 快捷函数到 `internal/model/`。

### 配置方式

编辑 `tools/config.yml`（gormt 配置，与 `conf/config.yml` 是两份独立配置）：

```yaml
base:
    is_dev : false        # 开发模式
out_dir : ./model         # 输出目录（make orm 已指定 -o internal/model，可覆盖）
url_tag : json            # web url tag
db_tag : gorm             # 数据库标签类型
simple : true             # 简单输出
is_out_sql : false        # 是否输出 sql 原信息
is_out_func : true        # 是否生成快捷 CRUD 函数（QueryOne/QueryAll/Update/Add/Delete 等）
is_web_tag : true         # 是否打 json 标记
is_web_tag_pk_hidden : true # json 标记是否隐藏主键
is_foreign_key : true     # 是否导出外键关联（生成关联结构）
is_table_name : true      # 直接生成表名、列名（不转驼峰）
is_null_to_point : false  # 允许 NULL 的列是否生成指针类型字段
table_prefix : ""         # 表前缀（有则填，如 "t_"）
db_info:                  # 数据库连接
    host : localhost
    port : 3306
    database : oauth_db
    username : root
    password : 123456
    type: 0               # 0:mysql, 1:sqlite, 2:mssql
```

**使用流程**：先在数据库建好表 → 修改 `tools/config.yml` 的 `db_info` 指向目标库 → `make orm` → 在业务代码中 import `internal/model` 使用。

### 表到结构体的映射规则与注意事项

- 一张表 → 一个 struct（表名 `snake_case` → 默认不转驼峰，`is_table_name: true` 时字段直用列名）。
- 列注释会写入 struct 字段注释；**表/列写好注释，生成代码才可读**。
- `is_null_to_point: true` 时 `DEFAULT NULL` 的列生成指针类型字段，避免零值歧义（涉及可空字段建议开启）。
- 有外键的表会生成嵌套关联结构（`is_foreign_key: true`），可直接带关联查询。
- 生成的代码在 `internal/model/`，**禁止手工修改**；表结构变更后重新 `make orm` 覆盖。
- `make clear` 会清空 `internal/model/`，`make build` 每次也会先 `make clear` + 重新 `make orm` 再编译——**保证模型与库表始终一致**。

---

## 6. 构建可执行程序

在服务目录执行，均为 `CGO_ENABLED=0` 静态交叉编译，产物为单文件可执行程序：

| 命令 | 产物 | 适用场景 |
|---|---|---|
| `make windows` | `[service].exe`（amd64） | Windows 环境（开发机自测、Windows 服务器部署） |
| `make linux` | `[service]`（amd64） | **生产部署标准命令**（服务器多为 Linux） |
| `make mac` | `[service]`（darwin amd64） | macOS 本地调试/演示 |

补充命令：

- `make run`：本地快速构建 + 前台启动（开发自测首选，不会先 clear/orm）。
- `make build`：清理 → `make orm` 重新生成模型 → 编译 → `./[service] debug` 前台运行（**提测/发布前用它做完整校验**）。
- `make clear`：清理产物（会删除 `internal/model/`、`err/`、二进制、以及 `./rpc` 下所有 .go —— 谨慎操作）。

部署产物：编译出的单个二进制 + `conf/config.yml`（+ 需要的静态资源目录）拷贝到服务器即可运行；`is_dev: false` 时以服务方式托管，`./[service] debug` 可强制前台调试。

---

## 7. 其他开发环节规范

### 7.1 配置管理

- 配置文件：`conf/config.yml`（yaml），结构体定义在 `internal/config/MyIni.go`（`Config`），基础字段在 `common.go`（`CfgBase`）。
- 新增配置项：先在 `MyIni.go` 的 `Config` struct 加字段（带 `yaml` + `consul` tag 及注释），再加对应 `GetXxx()` 访问函数，最后在 `config.yml` 补默认值。
- **conf 文件不入库敏感信息**：生产密码等通过 consul 覆盖——配置了 `consul_addr` 后框架自动以 `consul_tag/service_name` 为前缀拉取 KV 并热更新（`consult.AutoLoadConfig`），本地文件只放开发默认值。
- `is_dev: true` 为开发模式（前台直跑 + ginrpc debug + swagger 文档输出）。

### 7.2 日志规范

统一使用 `github.com/xxjwxc/public/mylog`，禁止直接 `fmt.Println` 输出业务日志（仅限临时调试）：

```go
mylog.Info(...)   // 常规信息
mylog.Error(...)  // 错误，务必记录
mylog.Warn(...)   // 告警
```

中间件已统一打点：`api_root.go` 的 `GinAfter` 会记录每次调用的方法名与耗时（`[middleware] call[func] [耗时]`）。业务内打印遵循：请求入口打印关键参数（注意脱敏），异常路径必须 `mylog.Error`。

### 7.3 错误码与统一响应

框架统一响应封装在 `github.com/xxjwxc/public/message`，由 `GinAfter` 中间件自动包装——**业务代码只需返回 `(*proto.XxxReply, error)`，不要手工拼响应壳**：

- 成功：`message.GetSuccessMsg()`，data 填入 resp。
- 失败：返回 error 即可；gRPC error 会被 `status.Convert` 解出 code/message；自定义错误码用 `message.GetErrorMsg(message.XxxCode)` 体系。
- 错误码统一定义在 `message` 包的常量体系中，不要散落各服务自行造码。

### 7.4 单元测试

参照 `main_test.go` 的模式（TestServer + TestClient）：

- `TestServer`：构造 `micro.NewService` → `new(hello.Hello)` → `proto.RegisterXxxServer` + `ginrpc` 注册 → `plugin.Run` 起本地服务。
- `TestClient`：`micro.SetClientServiceName(proto.GetExampleName(), "服务注册名")` → `proto.GetExampleClient()` 获取客户端 → 并发调用（示例用 `workpool.New(20)` 开 20 并发验证）。

运行方式：

```bash
# 在服务目录执行
go test -run TestServer -v        # 起服务端（阻塞）
go test ./... -v                  # 跑全部测试
go test -cover ./...              # 带覆盖率
```

测试是**双进程模式**：一个终端跑 `TestServer`，另一个终端跑客户端用例（或直接用 TestClient 压测）。

### 7.5 代码提交与部署流程

标准流程：

1. 修改 proto → `make gen`；改表结构 → `make orm`。
2. `make build` 本地完整校验（自动重生成模型 + 编译 + 前台运行）。
3. 自测通过后提交代码；commit message 说明改动点（proto 变更 / 业务变更 / 表结构变更分类型描述）。
4. 发布：`make linux` 交叉编译 → 二进制 + `conf/config.yml` 上传服务器 → 重启（`server.On` 托管模式支持 `restart_*.sh` 类脚本）。
5. 上线后验证：`GET /health`（健康检查）、`GET /metrics`（grafana 监控）、swagger 文档页面抽查接口。

### 7.6 环境安装与依赖升级（仓库级命令，仓库根目录执行）

```bash
make install          # 安装 protoc 环境（mac 脚本；linux 用 proto_install_linux.sh，
                      #  windows git-bash 用 proto_install_windows_git_bash.sh）
make source_install   # 源码方式安装
# 服务目录内：
make master           # 升级 public/ginrpc/micro/goplugins 四大依赖到 master
```

### 7.7 常见问题排查

| 现象 | 原因与处理 |
|---|---|
| `make gen` 报 `protoc: command not found` | 未安装 proto 环境 → 仓库根目录 `make install`（Windows 用 git-bash 执行 windows 脚本） |
| Windows 下 `rpc` 目录不是软链 | 正常：软链仅在 Linux/mac 创建（`make gen` 中的 `ln -s` 分支），Windows 直接用拷贝同步 |
| 业务对象没生效 / 404 | 未在 `api_root.go` 的 `OnInitRoot` 中 `args = append` 或漏了 `proto.RegisterXxxServer` |
| 服务启动即退出、日志为空 | `conf/config.yml` 缺失时框架会自动生成默认配置并 `os.Exit(0)` —— 补好配置后重启 |
| 端口冲突 | 修改 `config.yml` 的 `port`（grpc 与 gin 共用该端口） |
| 前端拿不到空字段 | 生成流程已自动去除 `,omitempty`；若字段缺失检查是否手工改过 `rpc/` 下生成代码 |
| 找不到 gormt 工具 | `tools/gormt` 不存在时到 https://github.com/xxjwxc/gormt 下载对应平台版本放入 `tools/` |
| `make clear` 后编译失败 | clear 会删除 `rpc/` 下所有 .go，需重新 `make gen` |
| 服务间调用失败 | 客户端需先 `micro.SetClientServiceName(proto.GetXxxName(), "服务注册名")`，且注册名与服务端 `micro.WithName` 一致 |
| etcd/consul/jaeger 连不上 | 检查 `config.yml` 对应 `etcd_info`/`consul_addr`/`jaeger_addr`；默认 DNS 注册模式无需 etcd |

---

## 8. 核心依赖库（团队重要库）

本框架的能力几乎全部来自以下几个自有库，遇到问题时**优先查这些库的 README / 源码**，不要绕过它们自行实现同等功能：

| 库 | 角色 | 本框架中的体现 |
|---|---|---|
| [xxjwxc/public](https://github.com/xxjwxc/public) | 基础库 | `mylog` 日志、`message` 统一错误码/响应、`tools` 工具函数、`dev` 开发模式、`mydoc/myswagger` 文档导出、`consult` 配置热更新 |
| [xxjwxc/ginrpc](https://github.com/xxjwxc/ginrpc) | gRPC ↔ Gin 自动绑定 | 业务对象注册即生成 REST 路由、`@Router` 注解自定义路径/方法、参数自动绑定与校验、swagger/markdown 文档生成、`WithBeforeAfter` 中间件 |
| [gmsec/micro](https://github.com/gmsec/micro) | 微服务内核 | `micro.NewService`、服务注册发现（DNS/mdns、etcd/nacos）、`SetClientServiceName` 客户端寻址、`tracer` 链路追踪 |
| [gmsec/goplugins](https://github.com/gmsec/goplugins) | 服务运行插件 | `plugin.Run` 同时拉起 gRPC + Gin 并共用端口、`api` 自定义上下文、注册中心插件 |
| [xxjwxc/gormt](https://github.com/xxjwxc/gormt) | ORM 代码生成 | `make orm` / `tools/gormt`，表 → gorm 模型 + CRUD 函数 |

### 升级依赖

```bash
# 在服务目录执行：一次性把四个核心库升到 master 最新
make master
# 等价于：
go get -u github.com/xxjwxc/public@master
go get -u github.com/xxjwxc/ginrpc@master
go get -u github.com/gmsec/micro@master
go get -u github.com/gmsec/goplugins@master
```

**时机**：需要用到框架新特性、修复了上游 bug 时；升级后必须 `go mod tidy` + `make build` 全量回归。

### 本地联调（改库源码时）

`go.mod` 中保留着注释掉的 replace 模板，调试库代码时打开即可，**提交前务必注释回去**（避免把本地路径带进版本库）：

```go
// replace github.com/xxjwxc/public => ../../public
// replace github.com/xxjwxc/ginrpc => ../../ginrpc
// replace github.com/gmsec/goplugins => ../../goplugins
// replace github.com/gmsec/micro => ../../micro
```

**约定**：不随意 fork 或替换这四个库；确需改动先提 issue/PR 走上游，避免各服务版本分叉。

---

## AI 助手行为约束（重要）

1. 修改任何 proto 后，必须提醒/协助执行 `make gen`；生成文件（`rpc/`、`apidoc/rpc/`、`internal/model/`、`docs/`）一律不手工编辑。
2. 新增接口的完整动作序列：写 proto（带完整注释）→ `make gen` → 在 `internal/service/` 实现 → 在 `api_root.go` 注册 → `make run` 验证。
3. 命令执行位置要判断准确：仓库根目录命令（`make new`、`make install`）与服务目录命令（`make gen/orm/build/windows/linux/mac/run`）不可混用。
4. 涉及 `make clear`、删除文件等破坏性操作，先向用户确认再执行。
5. 遵循本框架的既有模式（空 struct 业务对象 + OnInitRoot 显式注册 + message 统一响应），不要引入框架外的 DI、日志、错误码体系。
6. 需要自定义 REST 接口名/路径时，**只在 Go 方法上加 `@Router` 注解**（见 4.5），不得手改 `gen_router.go`/`gen_router.data`，也不得为了改路径去改 proto 方法名。
7. 四个核心库（public / ginrpc / micro / goplugins）属于团队自维护库：改其行为应先说明原因并确认是否走上游 PR，不在业务服务里打补丁式规避。
