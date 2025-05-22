# reconciliation-block

基于区块链的可信对账系统

## 一、项目基本信息
### 1. 项目简介
`reconciliation-block` 是一个基于区块链的可信对账系统，为数据的可信对账提供了安全可靠的解决方案。该项目采用了前后端分离的架构，前端使用 Vue 3 和 Element Plus 构建用户界面，后端使用 Go 语言和 Gin 框架进行开发，数据库使用 MySQL 和 Redis。

### 2. 技术栈
- **后端**：
  - **Go**：版本 1.16
  - **Gin**：版本 1.7.0
  - **GORM**：版本 1.22.5
  - **MySQL**：数据库存储
  - **Redis**：缓存和数据存储

- **前端**：
  - **Vue**：版本 3.2.25
  - **Element Plus**：版本 2.0.1

## 二、项目结构

### 1. 前端项目结构
```
├─api       （frontend APIs）
├─assets	（static files）
├─components（components）
├─router	（frontend routers）
├─store     （vuex state management）
├─style     （common styles）
├─utils     （frontend common utilitie）
└─view      （pages）
```

### 2. 后端项目结构
| 文件夹 | 说明 | 描述 |
| --- | --- | --- |
| `--excel` | excel导入导出默认路径 | excel导入导出默认路径 |
| `--page` | 表单生成器 | 表单生成器 打包后的dist |
| `--template` | 模板 | 模板文件夹,存放的是代码生成器的模板 |
| `router` | 路由层 | 路由层 |
| `service` | service层 | 存放业务逻辑问题 |
| `source` | source层 | 存放初始化数据的函数 |
| `utils` | 工具包 | 工具函数封装 |
| `--timer` | timer | 定时器接口封装 |
| `--upload` | oss | oss接口封装 |

## 三、使用方法

### 1. 前端项目启动
在 `reconciliation-block/web` 目录下执行以下命令：
```bash
# 安装依赖
npm install

# 开发环境启动，支持热更新
npm run serve
```

### 2. 后端项目启动
在 `reconciliation-block/server` 目录下执行以下命令：
```bash
# 使用 go.mod 管理依赖
# 安装 Go 模块
go list (go mod tidy)

# 构建服务器
go build

# 运行服务器
./server
```

### 3. Docker 部署
在 `reconciliation-block/deploy/docker-compose` 目录下，使用 `docker-compose.yaml` 文件进行部署：
```bash
docker-compose up -d
```

### 4. API 文档自动生成（使用 Swagger）
#### 安装 Swagger
```bash
go get -u github.com/swaggo/swag/cmd/swag
```

