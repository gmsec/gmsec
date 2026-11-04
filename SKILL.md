---
name: gmsec-dev
description: gmsec 仓库级微服务开发规范：新建微服务（make new）、proto 接口设计与注释规范（apidoc）、make gen 代码生成机制与产物、核心依赖库管理（public/ginrpc/micro/goplugins/gormt）、proto 环境安装、跨服务命名约定与仓库级问题排查。当用户在 gmsec 仓库进行上述仓库级事项，或涉及多服务协作、apidoc 维护、服务创建与全局规范时，必须加载本 skill 并严格遵循；单个服务目录内的日常开发规范见各服务目录下的 SKILL.md（以 example/SKILL.md 为模板，make new 会自动继承到新服务）。
---

# gmsec 微服务开发指南 · 仓库级（AI 助手必读）

gmsec 是基于 **gRPC + Gin** 的 Go 微服务集成框架：一套 proto 定义同时生成 gRPC 接口和 RESTful 接口（ginrpc 自动绑定参数），grpc 与 gin 共用端口，内置 gormt 数据库代码生成、服务注册发现（默认 DNS/mdns，可切换 etcd）、Jaeger 链路追踪、Prometheus 监控、Swagger/Markdown 文档自动导出。

## 两级 skill 分工（各司其职）

| 文档 | 职责范围 | 覆盖内容 |
|---|---|---|
| **本文档**（仓库根 `SKILL.md`） | 仓库级 | 新建微服务、proto 接口设计规范、`make gen` 生成机制、核心依赖库、proto 环境安装、跨服务命名约定、仓库级排查 |
| **`[service]/SKILL.md`**（服务目录内） | 服务目录内 | 业务实现与挂载注册、`@Router` 自定义路由、`make orm`、配置/日志/错误码、单元测试、构建部署、服务内排查 |

`example/SKILL.md` 是服务级文档的模板：`make new` 复制 example 时会把它一并带入新服务，因此**服务级文档写在 example 里即自动被所有新服务继承**。本文档中 `[服务名]`/`[service]` 为占位符，适用于仓库内任意微服务乃至新仓库；example 仅作示例引用。

**仓库根目录结构**：
- `Makefile` —— 仓库级命令（新建服务、proto 环境安装）
- `apidoc/` —— proto 源文件与生成的 rpc 代码（**跨服务共享**）
- `example/` —— 服务模板项目，新服务均由它复制而来

**执行原则**：所有 make 命令必须在正确目录下执行——`make new service=xxx` 在**仓库根目录**执行；`make gen`、`make orm`、`make windows/linux/mac` 等**先 `cd [service]`** 再执行。

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

1. 复制 `./example/` 整个模板到 `./[service]/`（若目录已存在则跳过，不会覆盖）。**example/SKILL.md 等文档也会一并带入**。
2. 全量字符串替换 `example` → `[service]`，涉及：
   - `./[service]/go.mod`（module 名 → `[service]`）
   - `./[service]/Makefile`（`NAME := [service]`）
   - `./[service]/generate/proto_makefile`（proto 编译目标）
   - `./[service]/conf/config.yml`（`service_name` 等）
   - 所有 `*.go` 文件的 import 路径（`example/rpc/example` → `[service]/rpc/[service]`）
3. 复制 `./apidoc/proto/example/` → `./apidoc/proto/[service]/`，并将其中 proto 的 `package`、`service` 名替换为 `[service]`。

### 命名规范（跨服务全局约定）

- 服务名：**全小写、单个英文单词**（如 `user`、`order`、`pay`），不用驼峰、不用下划线。
- 服务名会同时成为：Go module 名、proto 的 `package`/`service` 名、生成代码包名、服务注册名 `gmsec.srv.[service]`、REST 路由前缀 `/[service]/api/v1`。**一旦确定不要随意改动**。
- gRPC 服务注册名固定格式：`gmsec.srv.[service]`（见 `main.go` 中 `micro.WithName`）。

### 新建服务后的必做事项

1. `cd [service]` 进入服务目录，先执行 `make gen` 重新生成本服务的 proto 代码（模板里带的是 example 的产物，替换后需重编译确认无误）。
2. 检查 `conf/config.yml`，修改端口（`port`）等配置，避免与已有服务端口冲突。
3. 执行 `make run`（或 `go build`）验证服务能正常启动。
4. 服务内后续开发遵循该服务目录下的 `SKILL.md`（由 example 继承而来）。

### 服务目录结构（make new 派生后各服务一致）

```
[service]/
├── SKILL.md                  # 服务级开发指南（从 example 继承）
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

---

## 2. 编写 API 接口（proto 文件）

### 文件位置（仓库级约定）

proto 源文件统一放在 **`apidoc/proto/[service]/`** 目录，不要放到服务目录里。公共消息（如通用响应结构）放在 `apidoc/proto/common/`。

### 编写规范

以 `apidoc/proto/example/hello.proto` 为标准模板：

```protobuf
syntax = "proto3"; // 指定proto版本（固定 proto3）

package [service];  // 包名 = 服务名（make new 已自动替换）

import "common/common_public.proto"; // 引用公共消息（按需）

option go_package = "rpc/[service]";   // 固定格式：rpc/[服务名]

// 定义[service]服务（注释必写，service 上方一行）
service [service] {
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
4. 跨 proto 引用公共消息走 `import "common/common_public.proto"`，公共消息一律定义在 `apidoc/proto/common/`，**不要在各服务的 proto 里重复定义通用结构**；公共消息变更会影响所有服务，修改前需评估影响面。
5. 字段编号：新增字段只增不改不复用；已上线接口不允许修改已有字段编号或类型（protobuf 向后兼容原则）。
6. 一个服务 proto 可拆多个文件（`generate/proto_makefile` 会遍历目录下所有 `.proto`），建议按业务域拆分，文件名小写。

---

## 3. 生成 API 代码（make gen 机制）

### 命令

```bash
# 先进入服务目录
cd [service]
make gen
```

`make gen` 实际执行 `go generate` → `make -s -f generate/proto_makefile`，后者完成：

1. （Linux/mac）创建 `rpc` 软链接指向 `../rpc`，保证各服务共享 `apidoc/rpc` 下的公共代码；Windows 无软链，用拷贝同步。
2. 用 `protoc --gmsec_out=plugins=gmsec:./` 编译 `apidoc/proto/[service]/*.proto`。
3. 自动去掉生成代码中的 json `,omitempty` 标记（便于前端拿到完整字段）。
4. 将生成产物同步拷贝到 `../apidoc/rpc/[service]/`（供仓库级共享——**服务间引用对方的 proto 结构时，从这里引用**）。
5. 同时编译 `apidoc/proto/common/*.proto` 到 `apidoc/rpc/common/`。

### 生成产物及用途

| 产物 | 位置 | 用途 |
|---|---|---|
| `[name].pb.go` | `rpc/[service]/` | protobuf 标准生成代码：消息结构体（请求/响应）、序列化 |
| `[name].pb.gmsec.go` | `rpc/[service]/` | 框架扩展代码：gRPC 服务端注册函数（如 `proto.RegisterXxxServer`，Xxx 为服务名大驼峰）、客户端获取函数（`proto.GetXxxClient()`）、服务名函数（`proto.GetXxxName()`） |
| `apidoc/rpc/common/*.pb.go` | `apidoc/rpc/common/` | 公共消息结构，各服务共享 |
| `apidoc/rpc/[service]/*.pb.go` | `apidoc/rpc/[service]/` | 本服务产物的共享副本，供其他服务跨服务引用 |
| `docs/markdown/*.md`、`docs/swagger/swagger.json` | 服务目录 `docs/` | 由代码注释自动导出的 API 文档 |

**使用时机**：每次新增/修改 proto 文件后必须执行（在对应服务目录）；生成文件**禁止手工修改**（会被下次生成覆盖）；`make clear` 会删除 `rpc/` 下所有 .go，之后需重新 `make gen`。

---

## 4. 核心依赖库（团队重要库）

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

**时机**：需要用到框架新特性、修复了上游 bug 时；升级后必须 `go mod tidy` + `make build` 全量回归。多服务场景建议**同一批升级**，避免版本分叉导致行为不一致。

### 本地联调（改库源码时）

各服务 `go.mod` 中保留着注释掉的 replace 模板，调试库代码时打开即可，**提交前务必注释回去**（避免把本地路径带进版本库）：

```go
// replace github.com/xxjwxc/public => ../../public
// replace github.com/xxjwxc/ginrpc => ../../ginrpc
// replace github.com/gmsec/goplugins => ../../goplugins
// replace github.com/gmsec/micro => ../../micro
```

**约定**：不随意 fork 或替换这四个库；确需改动先提 issue/PR 走上游，避免各服务版本分叉。

---

## 5. proto 环境安装（仓库级）

```bash
# 仓库根目录执行
make install          # 安装 protoc 环境（mac 脚本）
make source_install   # 源码方式安装
# 对应脚本：proto_install_mac.sh / proto_install_linux.sh /
#          proto_install_windows_git_bash.sh（Windows 需在 git-bash 下执行）
```

---

## 6. 仓库级常见问题排查

| 现象 | 原因与处理 |
|---|---|
| `make new` 执行后服务没变化 | 目标目录已存在时跳过复制（防覆盖）——换服务名或先移除旧目录 |
| `make gen` 报 `protoc: command not found` | 未安装 proto 环境 → 仓库根目录 `make install`（Windows 用 git-bash 执行 windows 脚本） |
| Windows 下 `rpc` 目录不是软链 | 正常：软链仅在 Linux/mac 创建（`make gen` 中的 `ln -s` 分支），Windows 直接用拷贝同步 |
| 公共消息（`apidoc/proto/common/`）改动后其他服务没生效 | 修改方执行 `make gen` 会重新编译 common；引用方服务需各自重新 `make gen` + 编译 |
| 服务间 proto 引用找不到包 | 跨服务引用走 `apidoc/rpc/` 共享产物；确认被引用服务已执行过 `make gen` 且产物已同步 |
| `make clear` 后编译失败 | clear 会删除 `rpc/` 下所有 .go，需重新 `make gen` |
| 服务间调用失败 | 客户端需先 `micro.SetClientServiceName(proto.GetXxxName(), "服务注册名")`，且注册名与服务端 `micro.WithName` 一致 |
| 多服务端口冲突 | 各服务在 `conf/config.yml` 的 `port` 必须互不相同（grpc 与 gin 共用该端口） |

---

## AI 助手行为约束（仓库级，重要）

1. 命令执行位置要判断准确：仓库根目录命令（`make new`、`make install`）与服务目录命令（`make gen/orm/build/windows/linux/mac/run`）不可混用。
2. 修改任何 proto 后，必须提醒/协助执行 `make gen`；生成文件（`rpc/`、`apidoc/rpc/`、`docs/`）一律不手工编辑。
3. proto 文件位置在 `apidoc/proto/[service]/`，不要在服务目录里新建 proto；公共消息只加在 `apidoc/proto/common/`。
4. 涉及 `make clear`、删除文件、覆盖已有服务等破坏性操作，先向用户确认再执行。
5. 四个核心库（public / ginrpc / micro / goplugins）属于团队自维护库：改其行为应先说明原因并确认是否走上游 PR，不在业务服务里打补丁式规避。
6. 服务目录内的业务实现、@Router、ORM、测试、部署等规范，遵循对应服务目录下的 `SKILL.md`，两级文档不要互相重复维护。
