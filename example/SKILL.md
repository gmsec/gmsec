---
name: example-dev
description: gmsec example 示例微服务专用开发指南。当用户在 gmsec/example 目录下开发、调试、修改 proto 接口、自定义 REST 接口名与路由（@Router 注解、ginrpc）、执行 make gen / make orm / make run / make build / make master / make windows / make linux / make mac、编写业务逻辑或单元测试时，必须加载本 skill 并遵循其中的目录约定、命令流程与代码规范。example 既是官方示例项目，也是通过 make new 复制派生其他微服务的模板，修改它会影响所有新建服务，涉及模板改动需先向用户确认。
---

# gmsec example 示例服务开发指南（AI 助手必读）

## 项目定位（先读懂再动手）

`example/` 是 gmsec 框架的**官方示例微服务**，同时是**新服务的模板**（仓库根目录 `make new service=xxx` 会整体复制本目录并全局替换 `example` → `xxx`）。因此它有双重身份：

- **示例**：演示 gRPC + Gin 双协议接入的完整写法（hello 业务）。
- **模板**：本目录的任何结构性改动（目录布局、main.go 骨架、注册方式）都会被未来新建的服务继承——**改动需谨慎，涉及模板级调整先向用户确认**。

**本服务关键信息**：

| 项 | 值 |
|---|---|
| Go module | `example` |
| gRPC 服务注册名 | `gmsec.srv.example` |
| REST 路由前缀 | `/example/api/v1` |
| 端口 | `82`（conf/config.yml，gRPC 与 Gin 共用） |
| 示例业务 | `hello`（SayHello 方法） |
| 数据库 | mysql `localhost:3306/oauth_db`（tools/config.yml） |

**命令执行位置**：本 skill 所列 make 命令均在 **example 目录内**执行（即当前项目根目录）。仅 `make new service=xxx` 是仓库级命令，需在上级 gmsec 目录执行，与本目录无关。

---

## 1. 目录结构与职责

```
example/
├── main.go                  # 入口：CallBack() 完成服务初始化全链路；main() 判断 dev/托管模式
├── main_test.go             # 集成测试：TestServer 起服务 + TestClient 并发压测
├── gogenerate.go            # //go:generate 入口，make gen 的触发点
├── Makefile                 # 本项目全部命令（见下表）
├── go.mod / go.sum
├── conf/
│   ├── config.yml           # 服务配置（base/mysql/redis/etcd/consul/jaeger/port）
│   └── gen_router.data      # ginrpc 自动生成的路由数据（勿手改）
├── generate/
│   └── proto_makefile       # protoc 编译脚本：编译 ../apidoc/proto/example 与 common
├── rpc/example/             # proto 生成代码（hello.pb.go / hello.pb.gmsec.go），勿手改
├── docs/
│   ├── markdown/Hello.md    # 由注释自动导出的 markdown API 文档
│   └── swagger/swagger.json # swagger 文档
├── tools/
│   ├── gormt(.exe)          # ORM 代码生成工具
│   └── config.yml           # gormt 配置（指向 oauth_db，生成到 internal/model）
├── file/                    # 静态资源目录（/file 路由挂载）
└── internal/
    ├── api/
    │   ├── context.go       # ginrpc 上下文定义（api.NewAPIFunc）
    │   └── user.go          # 自定义 API 上下文扩展样例
    ├── config/
    │   ├── common.go        # CfgBase 基础配置 + consul 热更新逻辑
    │   └── MyIni.go         # Config 结构体（新增配置项在此扩展）
    ├── core/                # dao 封装、etcd 客户端示例
    ├── routers/
    │   ├── api_root.go      # ★ 业务挂载唯一入口：OnInitRoot（新业务必改此文件）
    │   └── gen_router.go    # ginrpc 自动生成的路由注册（勿手改）
    ├── service/
    │   └── hello/hello.go   # 业务实现样板（新业务照此写）
    └── model/               # gormt 生成产物（make orm 输出，勿手改）
```

**禁止手工编辑**：`rpc/**`、`apidoc/rpc/**`（上级目录）、`internal/model/**`、`docs/**`、`conf/gen_router.data`、`internal/routers/gen_router.go` —— 全部是生成产物，改动会被覆盖。

## 2. Makefile 命令速查（全部在 example 目录执行）

| 命令 | 作用 | 使用时机 |
|---|---|---|
| `make gen` | 经 `go generate` → `generate/proto_makefile` 编译 proto 生成接口代码 | **每次修改 `../apidoc/proto/example/*.proto` 后必跑** |
| `make orm` | `./tools/gormt -o internal/model` 生成 ORM 模型 | 数据库表结构变更后 |
| `make run` | `go build` + `./example debug` 前台运行（不清模型） | 日常开发自测首选 |
| `make build` | `make clear` → 重新 `make orm` → 编译 → 前台运行 | 提测/发布前完整校验 |
| `make windows` | 交叉编译 `example.exe`（windows/amd64） | Windows 环境部署 |
| `make linux` | 交叉编译 `example`（linux/amd64） | 生产服务器部署标准命令 |
| `make mac` | 交叉编译 `example`（darwin/amd64） | macOS 调试 |
| `make clear` | 清理：internal/model、err/、二进制、rpc 下所有 .go | ⚠️ 谨慎，执行后需 `make gen` 恢复 |
| `make install` / `make source_install` | 安装 proto 环境（调上级 proto_install.sh） | protoc 不可用时 |
| `make master` | 升级 public/ginrpc/micro/goplugins 依赖到 master | 同步框架最新版 |

## 3. 标准开发流程：新增一个接口（以完整动作为准）

以新增业务 `foo` 为例，严格按以下顺序：

**第 1 步：写 proto** —— 编辑 `../apidoc/proto/example/`（在 example 的上级 apidoc 目录，不是本目录）

```protobuf
syntax = "proto3";
package example;                        // 固定为 example
import "common/common_public.proto";    // 需要公共消息时
option go_package = "rpc/example";       // 固定

// FooBar 消息头注释：英文标识 + 中文说明
service Foo {
    // DoFoo 方法注释（会导出到 API 文档）
    rpc DoFoo(FooRequest) returns (FooReply) {}
}

// FooRequest 请求结构
message FooRequest {
    string name = 1; // 名字（每个字段必须行尾注释）
}

// FooReply 响应结构
message FooReply {
    common.ResultResp result = 1; // 通用结果
}
```

规则：`package`/`service` 不得改名（本服务固定 example）；新 proto 文件可按业务域拆分放同目录；公共消息只加在 `apidoc/proto/common/`，不在业务 proto 里重复定义；字段编号只增不改。

**第 2 步：生成代码** —— 回到 example 目录

```bash
make gen
```

产物：`rpc/example/foo.pb.go`（消息结构体）+ `rpc/example/foo.pb.gmsec.go`（gRPC 注册函数 `proto.RegisterFooServer`、客户端 `proto.GetFooClient()`、服务名 `proto.GetFooName()`），并同步一份到 `../apidoc/rpc/example/`。

**第 3 步：实现业务** —— 新建 `internal/service/foo/foo.go`，照抄 hello 样板：

```go
package foo

import (
    "context"
    proto "example/rpc/foo" // 注意：包路径是 rpc/example，按生成代码实际 import 为准
)

// Foo 业务对象（空 struct，勿加业务字段，依赖走内部封装）
type Foo struct{}

// DoFoo 签名必须与 proto 一一对应
func (f *Foo) DoFoo(ctx context.Context, req *proto.FooRequest) (*proto.FooReply, error) {
    // gRPC 侧可取 metadata：md, ok := metadata.FromIncomingContext(ctx)
    // Gin 侧可取原始上下文：c := api.Width_context(ctx)
    return &proto.FooReply{}, nil // 返回 error 即返回失败，响应壳由框架统一包装
}
```

**第 4 步：挂载注册** —— 修改 `internal/routers/api_root.go` 的 `OnInitRoot`，两行缺一不可：

```go
h2 := new(foo.Foo)
args = append(args, h2)          // → gin restful 自动注册（默认路径 /example/api/v1/foo.do_foo）
proto.RegisterFooServer(s, h2)    // → gRPC 服务端注册
```

**第 5 步：验证**

```bash
make run    # 前台启动，浏览器/swagger 验证 restful；或 go test 走 gRPC 验证
```

## 4. 业务实现规范（照 hello.go 执行）

- **方法签名固定**：`func (o *T) Method(ctx context.Context, req *proto.XxxRequest) (*proto.XxxReply, error)`，与 proto 的 rpc 定义一一对应；漏实现会导致注册不完整。
- **同一份实现双协议生效**：gRPC 与 Gin 共用此方法，业务代码不区分调用来源；REST 路由由 ginrpc 按方法名自动映射。
- **统一响应勿手工拼**：成功返回 `(reply, nil)`；失败返回 `(nil, err)`——`api_root.go` 的 `GinAfter` 中间件自动包装 `message.GetSuccessMsg()/GetErrorMsg()`，gRPC error 会经 `status.Convert` 透出 code/message。错误码使用 `github.com/xxjwxc/public/message` 常量体系，不自造。
- **日志**：统一 `mylog.Info/Warn/Error`（`github.com/xxjwxc/public/mylog`）；每次调用耗时由中间件统一打点，业务内不重复打耗时日志；`fmt.Println` 仅限临时调试。
- **链路追踪**：Jaeger 已由 `main.go` + `UseJager` 中间件接好，业务无需关心；需要跨服务透传时用 gin 上下文里的 `jager_ctx`。

### 4.1 自定义 REST 接口名（`@Router` 注解，ginrpc）

REST 路由由 [ginrpc](https://github.com/xxjwxc/ginrpc) 生成。**默认不写注解**，路径即「对象名.方法名」下划线小写，带 req 的方法为 POST：

```
Hello.SayHello  →  POST /example/api/v1/hello.say_hello     （见 docs/markdown/Hello.md 实际导出）
```

需要自定义（更语义化的 URL、GET/POST 多方法、内部接口不暴露 HTTP）时，在**业务方法上方**加一行注解：

```go
// SayHello 打个招呼（该行注释会作为接口描述写入 swagger/markdown）
// @Router /hello/say [post,get]
func (h *Hello) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
    return &proto.HelloReply{Message: "ok"}, nil
}
```

生效后：`POST /example/api/v1/hello/say`、`GET /example/api/v1/hello/say`。

**注解规则：**

| 写法 | 含义 |
|---|---|
| `// @Router /hello/say [post,get]` | 路径相对组前缀 `/example/api/v1`；方法列表逗号分隔，可写 post/get/put/delete 等 |
| `// @Router /hello/say` | 省略方括号时**默认 get** |
| `// @Router - [-]` | **忽略** HTTP 映射，该方法只保留 gRPC 能力（内部接口常用） |
| 多条 `@Router` | 同一方法可注册多个路径/方法组合 |
| `// @Router /x [POST][router aaa,bbb]` | 第三方网关标记，配合网关插件 |
| 注解下另起 `// SayHello 说明文字` | 作为接口描述导出文档（ginrpc 以方法名开头的注释为 note） |

**注意事项：**

1. 注解必须紧跟函数上方（Go doc 注释组内），写在函数体内无效。
2. 注册时自动生成 `internal/routers/gen_router.go` 与 `conf/gen_router.data`，**禁止手改**；改注解后重新 `make run`/编译即可生效（不涉及 proto，无需 `make gen`）。
3. `gen_router.go` 的 `init()` 依赖 `internal/routers` 包被 import —— example/main.go 已 import，无需改动。
4. 自定义路径不要与已有接口重复，否则注册冲突启动失败。
5. 只影响 REST：gRPC 方法名与客户端 stub 仍按 proto，服务间调用不变。
6. `ginrpc.WithBigCamel(true)` 可把默认路径切成 `Hello.SayHello` 大驼峰风格（`api_root.go` 中按需加选项，默认关闭）。

## 5. 数据库操作（make orm）

```bash
make orm   # = ./tools/gormt -o internal/model
```

- 配置文件：`tools/config.yml`（当前指向 mysql `localhost:3306/oauth_db`，type: 0=mysql / 1=sqlite / 2=mssql）。
- 流程：建表（表、列写好注释）→ 确认 `tools/config.yml` 连接信息 → `make orm` → 业务代码 import `example/internal/model` 使用。
- 关键开关：`is_out_func: true`（生成 QueryOne/QueryAll/Update/Add/Delete 快捷函数）、`is_foreign_key: true`（外键生成嵌套关联）、`is_null_to_point: false`（可空列生成指针类型，涉及 NULL 字段建议开启）、`table_prefix`（表前缀按需填）。
- `make build` 每次都会 clear 后重新 `make orm`，模型始终与库表同步；`internal/model` 生成物禁手改。

## 6. 单元测试

参照 `main_test.go` 的双进程模式：

```bash
# 终端 1：起服务（注册 hello + gin 路由 + plugin.Run 阻塞）
go test -run TestServer -v

# 终端 2：客户端用例（SetClientServiceName → GetExampleClient → workpool 20 并发压测）
go test -run TestClient -v
# 或全量：
go test ./... -v
go test -cover ./...   # 覆盖率
```

客户端调用三要素：`micro.SetClientServiceName(proto.GetExampleName(), "服务注册名")`（注册名须与服务端 `micro.WithName` 一致，example 默认 `gmsec.srv.example`）→ `proto.GetExampleClient()` 取 stub → 正常传 ctx 调用。新增接口后应仿照 TestClient 添加对应调用用例。

## 7. 配置管理

- 结构体：`internal/config/MyIni.go` 的 `Config` + `common.go` 的 `CfgBase`。**新增配置项三步**：Config struct 加字段（带 `yaml`+`consul` tag）→ 加 `GetXxx()` 函数 → `conf/config.yml` 补默认值。
- 热更新：`config.yml` 配置了 `consul_addr` 时，框架以 `consul_tag/service_name` 为前缀自动拉取 consul KV 并热更新（`consult.AutoLoadConfig`）——**生产密码等敏感信息走 consul，本地文件只放开发默认值**。
- `is_dev: true`：前台直跑 + ginrpc debug + swagger 输出；`is_dev: false`：`server.On(...).Start` 托管模式（可注册为系统服务），带 `debug` 参数可强制前台。
- 端口在 `conf/config.yml` 的 `port: 82`，改端口注意与其他服务避让。

## 8. 构建、提交与部署

```bash
make run      # 开发自测（快速）
make build    # 提测前完整校验（清模型→重生成→编译→前台跑）
make linux    # 生产编译（服务器标准）
make windows  # Windows 部署
make mac      # macOS 调试
```

部署产物 = 单个二进制 + `conf/config.yml`（+ 需要的静态资源）→ 上传服务器 → 重启 → 验证 `/health`（健康检查）、`/metrics`（grafana 监控）、swagger 页面抽查接口。

提交注意：生成产物（rpc/、docs/、internal/model/）是否入库遵循仓库现状；commit message 区分 proto 变更 / 业务变更 / 表结构变更。

## 9. 常见问题排查（example 专属）

| 现象 | 处理 |
|---|---|
| `make gen` 报 protoc 未找到 | 在 example 目录 `make install`（Windows 需在 git-bash 下执行上级 `proto_install_windows_git_bash.sh`） |
| `make clear` 后编译不过 | clear 会删除 `rpc/example/*.go`，重新 `make gen` 即可 |
| 新接口 404 / gRPC 调不通 | 检查是否漏改 `api_root.go`：`args = append` 与 `proto.RegisterXxxServer` 两行都要 |
| 服务起不来且日志打印生成默认配置 | `conf/config.yml` 缺失或损坏，框架会自动生成默认配置并退出（os.Exit(0)）——补配置重启 |
| `tools/gormt` 不存在 | 到 https://github.com/xxjwxc/gormt 下载对应平台版本放入 `tools/` |
| 前端拿不到零值字段 | 生成流程已去 `,omitempty`；若仍缺失说明生成代码被手工改过，重新 `make gen` |
| 端口占用 | 改 `config.yml` 的 `port` |
| gRPC 客户端调用失败 | `SetClientServiceName` 的注册名与服务端不一致，或服务端未启动 |
| 想切换 etcd 注册 | `main.go` 中有注释掉的 etcdv3 示例，替换 `registry.NewDNSNamingRegistry()` 并配置 `etcd_info` |
| `@Router` 注解不生效 | 注解须紧贴函数上方；改后需重新编译启动；确认 `internal/routers` 包被 import（main.go 已 import） |
| 启动报路由重复注册 | 两个方法的 `@Router` 路径+方法完全相同，改用不同路径或方法 |

## 10. 核心依赖库（团队重要库）

example 的能力全部来自以下自有库，排障时优先查它们的 README/源码，不要绕开自行实现：

| 库 | 在 example 中的用途 |
|---|---|
| [xxjwxc/public](https://github.com/xxjwxc/public) | `mylog` 日志、`message` 统一错误码/响应、`tools`、`dev`、`mydoc/myswagger` 文档、`consult` 配置热更新 |
| [xxjwxc/ginrpc](https://github.com/xxjwxc/ginrpc) | gRPC↔Gin 自动绑定、`@Router` 自定义路由、参数绑定校验、文档导出、`WithBeforeAfter` 中间件 |
| [gmsec/micro](https://github.com/gmsec/micro) | `micro.NewService`、服务注册发现（DNS/etcd）、客户端寻址、`tracer` 链路追踪 |
| [gmsec/goplugins](https://github.com/gmsec/goplugins) | `plugin.Run` 同端口拉起 gRPC+Gin、`api` 上下文、注册中心插件 |
| [xxjwxc/gormt](https://github.com/xxjwxc/gormt) | `make orm` 的 ORM 代码生成工具 |

升级（在 example 目录执行）：

```bash
make master   # 等同于下面四条
go get -u github.com/xxjwxc/public@master
go get -u github.com/xxjwxc/ginrpc@master
go get -u github.com/gmsec/micro@master
go get -u github.com/gmsec/goplugins@master
```

升级后 `go mod tidy` + `make build` 全量回归。`go.mod` 中已备好注释掉的 `replace` 模板（`../../public`、`../../ginrpc`、`../../micro`、`../../goplugins`），改库源码联调时打开，**提交前务必注释回去**。

## AI 助手行为约束（重要）

1. example 是模板项目：**结构性改动（目录、骨架文件、注册方式）先向用户确认**，纯业务演示性质的补充可直接做。
2. 改 proto 必提醒执行 `make gen`；改表结构必提醒执行 `make orm`；生成产物一律不手工编辑。
3. 新增接口严格走第 3 节五步流程，新业务实现照 `hello.go` 样板写，不引入框架外的 DI/日志/错误码体系。
4. `make clear` 等破坏性命令先确认再执行。
5. proto 文件实际位置在上级 `../apidoc/proto/example/`，编辑 proto 时注意目录，不要在 example 目录里新建 proto。
6. 需要自定义 REST 接口名/路径时，只在业务方法上加 `@Router` 注解（见 4.1），不改 `gen_router.go`/`gen_router.data`，不为改路径而改 proto 方法名。
7. 四个核心库为团队自维护库，不在业务代码里做规避式补丁；需要改库行为先说明并确认是否走上游。
