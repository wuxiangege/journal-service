# journal-service

日记管理系统后端，基于 [go-zero](https://go-zero.dev/) + [GORM](https://gorm.io/) + MySQL，为 `journal-ui` 提供 REST API。

## 功能

- 用户登录（JWT）
- 日记 CRUD（标题、正文、日期、心情、标签、置顶）
- 列表筛选（关键词、标签、心情、近 7/30 天）
- 统计概览（本月篇数、平均字数、常用标签）

## 快速开始

### 1. 启动 MySQL

可用 Docker：

```bash
docker run -d --name journal-mysql \
  -e MYSQL_ROOT_PASSWORD=password \
  -e MYSQL_DATABASE=journal \
  -p 3307:3306 \
  mysql:8
```

按需修改 `etc/journal.yaml` 中的 `MySQL.DataSource`（默认 `127.0.0.1:3307`）。

### 2. 启动服务

```bash
cd journal-service
go run journal.go -f etc/journal.yaml
```

默认监听 `http://0.0.0.0:8080`（可在 `etc/journal.yaml` 的 `Port` 修改）。

首次启动会自动：

- 建表（`users`、`journal_posts`）
- 创建演示账号：`871240671@qq.com` / `123456`
- 导入与前端一致的 7 条示例日记

### 3. 启动前端

```bash
cd journal-ui
cp .env.example .env   # 首次
yarn dev
```

- 前端开发服务器默认：`http://localhost:5173`
- 本地 API 地址在 `journal-ui/.env` 中配置：`VITE_API_BASE=http://127.0.0.1:8080`
- 后端 `etc/journal.yaml` 的 `Cors.AllowOrigins` 已放行 `5173` 与 `8080`，直连后端无需 Vite 代理

## API 一览

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/v1/login` | 登录，返回 `token` |
| GET | `/api/v1/journal-posts` | 日记列表（需 JWT） |
| POST | `/api/v1/journal-posts` | 新建日记 |
| GET | `/api/v1/journal-posts/:id` | 单篇详情 |
| PUT | `/api/v1/journal-posts/:id` | 更新日记 |
| DELETE | `/api/v1/journal-posts/:id` | 删除日记 |
| GET | `/api/v1/stats` | 统计数据 |

请求头：`Authorization: Bearer <token>`

## 目录结构

```
journal.api          # API 定义（goctl 生成入口）
etc/journal.yaml     # 配置
internal/
  handler/           # HTTP 入口
  logic/             # 业务逻辑
  model/             # GORM 模型
  repo/              # 数据访问
  middleware/        # CORS 等
journal.go           # 主程序
```

## 重新生成代码

修改 `journal.api` 后：

```bash
goctl api go -api journal.api -dir . -style go_zero
```

生成后请保留手写逻辑（GORM、鉴权等）的改动。
