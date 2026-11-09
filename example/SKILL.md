---
name: gmsec-svc-dev
description: gmsec 服务目录内开发指南（通用，适用于 example 及一切由 make new 派生的服务）。当用户在某服务目录下开发、调试、修改接口、自定义 REST 路由（@Router 注解、ginrpc）、执行 make gen / make orm / make run / make build / make windows / make linux / make mac、编写业务逻辑或单元测试、管理配置、排查服务内问题时，必须加载本 skill 并遵循其中的目录约定、命令流程与代码规范。文档中 [service] 指当前服务名；proto 编写规范与 make gen 底层机制见仓库根目录 SKILL.md。
---

# gmsec 服务开发指南 · 服务目录内适用（AI 助手必读）

**适用范围**：本文档适用于本服务目录（`example` 或任何 `make new` 派生的服务）。由于 `make new` 复制 example 时会连同本文档带入新服务，因此写成通用形式，`[service]` 即当前服务名。仓库级事项（新建服务、proto 设计规范、make gen 机制、核心依赖库、环境安装）见**仓库根目录的 `SKILL.md`**。

---

## 1. 服务目录结构与职责

```
[service]/
├── SKILL.md                  # 本文档
├── main.go                    # 入口：CallBack() 完成服务初始化全链路；main() 判断 dev/托管模式
├── main_test.go               # 集成测试：TestServer 起服务 + TestClient 并发压测
├── gogenerate.go              # //go:generate 入口，make gen 的触发点
├── Makefile                   # 本项目全部命令（见第 2 节）
├── conf/
│   ├── config.yml             # 服务配置（base/mysql/redis/etcd/consul/jaeger/port）
│   └── gen_router.data        # ginrpc 自动生成的路由数据（勿手改）
├── generate/
│   └── proto_makefile         # protoc 编译脚本：编译 ../apidoc/proto/[service] 与 common
├── rpc/[service]/             # proto 生成代码（*.pb.go / *.pb.gmsec.go），勿手改
├── docs/
│   ├── markdown/              # 由注释自动导出的 markdown API 文档
│   └── swagger/swagger.json   # swagger 文档
├── tools/
│   ├── gormt(.exe)            # ORM 代码生成工具
│   └── config.yml             # gormt 配置（数据库连接，与 conf/config.yml 相互独立）
├── file/                      # 静态资源目录（/file 路由挂载）
└── internal/
    ├── api/
    │   └── context.go         # ginrpc 上下文定义（api.NewAPIFunc）
    ├── config/
    │   ├── common.go          # CfgBase 基础配置 + consul 热更新逻辑
    │   └── MyIni.go           # Config 结构体（新增配置项在此扩展）
    ├── core/                  # dao/连接等核心封装
    ├── routers/
    │   ├── api_root.go        # ★ 业务挂载唯一入口：OnInitRoot（新业务必改此文件）
    │   └── gen_router.go      # ginrpc 自动生成的路由注册（勿手改）
    ├── service/               # 业务逻辑实现（按业务名分目录）
    └── model/                 # gormt 生成产物（make orm 输出，勿手改）
```

**禁止手工编辑**：`rpc/**`、`apidoc/rpc/**`（上级共享目录）、`internal/model/**`、`docs/**`、`conf/gen_router.data`、`internal/routers/gen_router.go` —— 全部是生成产物，改动会被覆盖。

## 2. Makefile 命令速查（全部在本服务目录执行）

| 命令 | 作用 | 使用时机 |
|---|---|---|
| `make gen` | 经 `go generate` → `generate/proto_makefile` 编译 proto 生成接口代码 | **每次修改 `../apidoc/proto/[service]/*.proto` 后必跑** |
| `make orm` | `./tools/gormt -o internal/model` 生成 ORM 模型 | 数据库表结构变更后 |
| `make run` | `go build` + `./[service] debug` 前台运行（不清模型） | 日常开发自测首选 |
| `make build` | `make clear` → 重新 `make orm` → 编译 → 前台运行 | 提测/发布前完整校验 |
| `make windows` | 交叉编译 `[service].exe`（windows/amd64） | Windows 环境部署 |
| `make linux` | 交叉编译 `[service]`（linux/amd64） | **生产服务器部署标准命令** |
| `make mac` | 交叉编译 `[service]`（darwin/amd64） | macOS 调试 |
| `make clear` | 清理：internal/model、err/、二进制、`./rpc` 下所有 .go | ⚠️ 谨慎，执行后需 `make gen` 恢复 |
| `make install` / `make source_install` | 安装 proto 环境（调仓库级 proto_install*.sh） | protoc 不可用时 |
| `make master` | 升级 public/ginrpc/micro/goplugins 四个核心库到 master | 同步框架最新版 |

## 3. 标准开发流程：新增一个接口

以新增业务 `foo` 为例，严格按以下五步（proto 编写细节见仓库根 SKILL.md 第 2 节）：

**第 1 步：写 proto** —— 在 `../apidoc/proto/[service]/` 下新建（proto 在 apidoc，不在本目录）

```protobuf
syntax = "proto3";
package [service];                          // 固定等于服务名
import "common/common_public.proto";         // 需要公共消息时
option go_package = "rpc/[service]";         // 固定格式

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

**第 2 步：生成代码**

```bash
make gen    # 产物 rpc/[service]/foo.pb.go + foo.pb.gmsec.go
```

**第 3 步：实现业务** —— 新建 `internal/service/foo/foo.go`，照 hello 样板：

```go
package foo

import (
    "context"
    proto "[service]/rpc/[service]" // 引用生成的代码
)

// Foo 业务对象（空 struct，勿加业务字段，依赖走内部封装）
type Foo struct{}

// DoFoo 签名必须与 proto 一一对应
func (f *Foo) DoFoo(ctx context.Context, req *proto.FooRequest) (*proto.FooReply, error) {
    // gRPC 侧可取 metadata：md, ok := metadata.FromIncomingContext(ctx)
    // Gin 侧可取原始上下文：c := api.WidthContext(ctx)
    return &proto.FooReply{}, nil // 返回 error 即返回失败，响应壳由框架统一包装
}
```

**第 4 步：挂载注册** —— 修改 `internal/routers/api_root.go` 的 `OnInitRoot`，两行缺一不可：

```go
h2 := new(foo.Foo)
args = append(args, h2)          // → gin restful 自动注册（默认路径 /[service]/api/v1/foo.do_foo）
proto.RegisterFooServer(s, h2)   // → gRPC 服务端注册
```

**第 5 步：验证**

```bash
make run    # 前台启动，swagger/浏览器验证 restful；或 go test 走 gRPC 验证
```

## 4. 业务实现规范

- **方法签名固定**：`func (o *T) Method(ctx context.Context, req *proto.XxxRequest) (*proto.XxxReply, error)`，与 proto 的 rpc 定义一一对应；漏实现会导致注册不完整。
- **同一份实现双协议生效**：gRPC 与 Gin 共用此方法，业务代码不区分调用来源；REST 路由由 ginrpc 按方法名自动映射。
- **统一响应勿手工拼**：成功返回 `(reply, nil)`；失败返回 `(nil, err)`——`api_root.go` 的 `GinAfter` 中间件自动包装 `message.GetSuccessMsg()/GetErrorMsg()`，gRPC error 会经 `status.Convert` 透出 code/message。错误码使用 `github.com/xxjwxc/public/message` 常量体系，不自造。
- **日志**：统一 `mylog.Info/Warn/Error`（`github.com/xxjwxc/public/mylog`）；每次调用耗时由中间件统一打点，业务内不重复打耗时日志；`fmt.Println` 仅限临时调试。
- **链路追踪**：Jaeger 已由 `main.go` + `UseJager` 中间件接好（服务端 span 会从请求头 Extract 上游 span，并回写 `Request.Context()`）；REST handler 的 `ctx` 参数已注入 span，跨服务透传统一用 `api.TraceCtx(ginCtx)` 取上下文，不要用裸 `context.Background()`。
- **服务初始化流程**（`main.go` 的 `CallBack()`，一般无需改动）：swagger 设置 → Jaeger tracer → 注册方式（默认 DNS/mdns，etcd 见注释示例）→ `micro.NewService(micro.WithName("gmsec.srv.[service]"))` → gin 路由组 `/[service]/api/v1` → `routers.OnInitRoot`（挂载业务唯一入口）→ `plugin.Run` 同端口拉起 gRPC+Gin。启动模式：`is_dev: true` 或 `./[service] debug` 前台直跑；否则 `server.On(...).Start` 托管模式。
- **依赖注入**：显式构造 + 统一注册，不用额外 DI 框架；业务对象在 `OnInitRoot` 中 `new` 挂载，外部依赖经 `internal/core` 封装懒加载。

### 4.1 自定义 REST 接口名与路由（`@Router` 注解，ginrpc）

REST 路径由 [ginrpc](https://github.com/xxjwxc/ginrpc) 生成，**默认不写注解**，需要自定义时在业务方法上方加一行注解；gRPC 接口名不受影响（始终等于 proto 定义）。

**默认规则（不写注解时）**

| 项 | 规则 | 示例 |
|---|---|---|
| 路径 | `对象名.方法名`，均转为**下划线小写** | `Hello.SayHello` → `/[service]/api/v1/hello.say_hello` |
| HTTP 方法 | 有 req 参数（ctx+req 两参）→ `POST`；只有 ctx（一参）→ `ANY` | SayHello 两参 → `POST` |
| 大驼峰模式 | `ginrpc.WithBigCamel(true)` 时路径改为 `Hello.SayHello`（默认 false 为 web 风格） | — |

**注解写法**

```go
// SayHello 打个招呼（以方法名开头的注释会作为接口描述导出到 swagger/markdown）
// @Router /hello/say [post,get]
func (h *Hello) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
    return &proto.HelloReply{Message: "ok"}, nil
}
```

生效后：`POST /[service]/api/v1/hello/say`、`GET /[service]/api/v1/hello/say`。

| 写法 | 含义 |
|---|---|
| `// @Router /hello/say [post,get]` | 路径相对组前缀 `/[service]/api/v1`；方法列表逗号分隔，可写 post/get/put/delete 等 |
| `// @Router /hello/say` | 省略方括号时**默认 get** |
| `// @Router - [-]` | **忽略** HTTP 映射，该方法只保留 gRPC 能力（内部接口常用） |
| 多条 `@Router` | 同一方法可注册多个路径/方法组合 |
| `// @Router /x [POST][router aaa,bbb]` | 第三方网关标记，配合网关插件 |
| 注解下另起 `// 方法名 说明文字` | 作为接口描述导出文档 |

**注意事项：**

1. 注解必须紧跟函数上方（Go doc 注释组内），写在函数体内无效。
2. 注册时自动生成 `internal/routers/gen_router.go` 与 `conf/gen_router.data`，**禁止手改**；改注解后重新 `make run`/编译即可生效（不涉及 proto，无需 `make gen`）。
3. `gen_router.go` 的 `init()` 依赖 `internal/routers` 包被 import——main.go 已直接 import，无需改动。
4. 自定义路径不要与已有接口重复，否则注册冲突启动失败。
5. 只影响 REST：gRPC 方法名与客户端 stub 仍按 proto，服务间调用不变。

## 5. 数据库操作（make orm）

```bash
make orm   # = ./tools/gormt -o internal/model
```

[gormt](https://github.com/xxjwxc/gormt) 读取数据库表结构，生成 gorm 模型与 CRUD 快捷函数到 `internal/model/`。

- 配置文件：`tools/config.yml`（与 `conf/config.yml` 相互独立），核心项：`db_info`（host/port/database/username/password/type：0=mysql / 1=sqlite / 2=mssql）、`is_out_func: true`（生成 QueryOne/QueryAll/Update/Add/Delete 快捷函数）、`is_foreign_key: true`（外键生成嵌套关联）、`is_null_to_point: false`（可空列生成指针类型，涉及 NULL 字段建议开启）、`table_prefix`（表前缀按需填）、`is_table_name: true`（直用表名列名不转驼峰）。
- **流程**：建表（表、列写好注释，列注释会写入 struct 注释）→ 确认 `tools/config.yml` 连接信息 → `make orm` → 业务代码 import `internal/model` 使用。
- **映射规则**：一张表 → 一个 struct；`DEFAULT NULL` 列按 `is_null_to_point` 决定是否指针；外键表生成嵌套关联结构可直接带关联查询。
- **注意**：`internal/model` 生成物禁手改，表结构变更后重新 `make orm` 覆盖；`make build` 每次先 `make clear` + 重新 `make orm` 再编译——模型与库表始终一致。`tools/gormt` 不存在时到 https://github.com/xxjwxc/gormt 下载对应平台版本放入 `tools/`。

## 6. 配置管理

- 配置文件：`conf/config.yml`（yaml），结构体在 `internal/config/MyIni.go`（`Config`）+ `common.go`（`CfgBase`）。**新增配置项三步**：Config struct 加字段（带 `yaml`+`consul` tag 及注释）→ 加 `GetXxx()` 函数 → `config.yml` 补默认值。
- **热更新**：配置了 `consul_addr` 时，框架以 `consul_tag/service_name` 为前缀自动拉取 consul KV 并热更新（`consult.AutoLoadConfig`）——**生产密码等敏感信息走 consul，本地文件只放开发默认值**。
- `is_dev: true`：前台直跑 + ginrpc debug + swagger 文档输出；`is_dev: false`：`server.On(...).Start` 托管模式（可注册为系统服务），带 `debug` 参数强制前台。
- 端口：`conf/config.yml` 的 `port`，gRPC 与 Gin 共用；多服务部署时各端口不得冲突。

## 7. 单元测试

参照 `main_test.go` 的双进程模式：

```bash
# 终端 1：起服务（注册业务对象 + gin 路由 + plugin.Run 阻塞）
go test -run TestServer -v

# 终端 2：客户端用例
go test -run TestClient -v
# 或全量：
go test ./... -v
go test -cover ./...   # 覆盖率
```

- `TestServer`：构造 `micro.NewService` → `new(xxx.Xxx)` → `proto.RegisterXxxServer` + `ginrpc` 注册 → `plugin.Run` 起本地服务。
- `TestClient`：`micro.SetClientServiceName(proto.GetXxxName(), "服务注册名")`（注册名须与服务端 `micro.WithName` 一致）→ `proto.GetXxxClient()` 取客户端 stub → 并发调用（模板用 `workpool.New(20)` 开 20 并发压测）。
- 新增接口后应仿照 TestClient 添加对应调用用例。

## 8. 构建、提交与部署

```bash
make run      # 开发自测（快速）
make build    # 提测前完整校验（清模型→重生成→编译→前台跑）
make linux    # 生产编译（服务器标准）
make windows  # Windows 部署
make mac      # macOS 调试
```

- **提交前**：`make build` 本地完整校验通过；commit message 说明改动点（proto 变更 / 业务变更 / 表结构变更分类型描述）。
- **部署产物**：单个二进制 + `conf/config.yml`（+ 需要的静态资源）→ 上传服务器 → 重启（托管模式支持 `restart_*.sh` 类脚本）。
- **上线后验证**：`GET /health`（健康检查）、`GET /metrics`（grafana 监控）、swagger 文档页面抽查接口。

## 9. 常见问题排查（服务内）

| 现象 | 处理 |
|---|---|
| 新接口 404 / gRPC 调不通 | 检查是否漏改 `api_root.go`：`args = append` 与 `proto.RegisterXxxServer` 两行都要 |
| `@Router` 注解不生效 | 注解须紧贴函数上方；改后需重新编译启动；确认 `internal/routers` 包已被 import |
| 启动报路由重复注册 | 两个方法的 `@Router` 路径+方法完全相同，改用不同路径或方法 |
| 服务起不来且日志打印生成默认配置 | `conf/config.yml` 缺失或损坏，框架会自动生成默认配置并退出（os.Exit(0)）——补配置重启 |
| `make clear` 后编译不过 | clear 会删除 `rpc/[service]/*.go`，重新 `make gen` 即可 |
| `make gen` 报 protoc 未找到 | 本目录 `make install`（Windows 需在 git-bash 下执行仓库级 `proto_install_windows_git_bash.sh`） |
| `tools/gormt` 不存在 | 到 https://github.com/xxjwxc/gormt 下载对应平台版本放入 `tools/` |
| 前端拿不到零值字段 | 生成流程已去 `,omitempty`；若仍缺失说明生成代码被手工改过，重新 `make gen` |
| 端口占用 | 改 `config.yml` 的 `port` |
| gRPC 客户端调用失败 | `SetClientServiceName` 的注册名与服务端不一致，或服务端未启动 |
| 想切换 etcd 注册 | `main.go` 中有注释掉的 etcdv3 示例，替换 `registry.NewDNSNamingRegistry()` 并配置 `etcd_info` |
| etcd/consul/jaeger 连不上 | 检查 `config.yml` 对应 `etcd_info`/`consul_addr`/`jaeger_addr`；默认 DNS 注册模式无需 etcd |

## AI 助手行为约束（重要）

1. 本文档是通用模板：修改它等于修改所有 `make new` 派生服务的模板——**结构性改动先向用户确认**，业务示例性补充可直接做。
2. 改 proto 必提醒执行 `make gen`；改表结构必提醒执行 `make orm`；生成产物一律不手工编辑。
3. 新增接口严格走第 3 节五步流程，新业务实现照 `internal/service/hello/hello.go` 样板写，不引入框架外的 DI/日志/错误码体系。
4. 需要自定义 REST 接口名/路径时，只在业务方法上加 `@Router` 注解（见 4.1），不改 `gen_router.go`/`gen_router.data`，不为改路径而改 proto 方法名。
5. `make clear` 等破坏性命令先确认再执行。
6. proto 文件位置在上级 `../apidoc/proto/[service]/`，编辑 proto 时注意目录，不要在服务目录里新建 proto。
7. 仓库级事项（新建服务、公共消息、依赖库升级、多服务协作）遵循并引导用户查阅仓库根目录的 `SKILL.md`。
