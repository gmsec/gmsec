# syntax=docker/dockerfile:1

# ---- 构建阶段:固定 Go 版本,与 example/go.mod 的 go 1.26.10 对齐 ----
FROM golang:1.26.10 AS build
WORKDIR /src
COPY . .
WORKDIR /src/example
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/example .

# ---- 运行阶段:非 root,只拷二进制与配置模板 ----
FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=build /out/example /app/example
COPY --from=build /src/example/conf/config.yml /app/conf/config.yml
# 敏感项通过环境变量注入(见 .env.example),不要把真实口令打进镜像。
USER 65532:65532
EXPOSE 82
ENTRYPOINT ["/app/example"]
