# go_blog

gin + vue3 开发的个人博客项目

## 本地测试

### 开发工具及版本

golang: 1.23.5

node: v22.13.0

docker: 27.4.0

编译器：vscode、goland、webstorm

### 启动容器

```bash
docker run -itd --name mysql -p 3306:3306 -e  MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=blog_db -d mysql

docker run --name es -p 127.0.0.1:9200:9200 -e "discovery.type=single-node" -e "xpack.security.http.ssl.enabled=false" -e "xpack.license.self_generated.type=trial" -e "xpack.security.enabled=false" -e ES_JAVA_OPTS="-Xms84m -Xmx512m" -d elasticsearch:8.17.0

docker run --name redis -p 6379:6379 -d redis
```

### 启动服务

```bash
# 进入 server 文件夹，修改配置文件 config.yaml
go mod tidy
# 初始化 mysql
go run main.go -sql
# 初始化 es
go run main.go -es
# 创建管理员
go run main.go -admin
# 运行后端
go run main.go

# 进入 web 文件夹
npm install
# 运行前端
npm run dev
```

