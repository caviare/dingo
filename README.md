# dingo
本项目采用了一系列Golang中比较流行的组件，并且构建了一些常用的功能代码，可以以本项目为基础快速搭建服务。

## 项目依赖

本项目已经整合了许多开发API所必要的组件：

1. [Echo](https://echo.labstack.com/): 高性能、可扩展、极简的 Go web 框架 
2. [GORM](https://echo.labstack.com/): ORM工具。本项目需要配合Mysql使用
3. [jwt-go](https://github.com/golang-jwt/jwt): JSON Web Tokens
4. [godotenv](https://github.com/joho/godotenv):  从 `.env` 加载环境变量。
5. [scrypt](https://pkg.go.dev/golang.org/x/crypto@v0.0.0-20210921155107-089bfa567519/scrypt)： 密钥派生函数

## 注意事项

- 本项目使用`Go Mod`管理依赖。
- 项目运行后启动在`2233`端口
- 项目在启动时依赖相关环境变量，但是在也可以在项目根目录创建`.env`文件设置环境变量便于使用，相关环境变量说明请查看示例文件：`.env.example`

## 相关模块

1. **routers** 文件夹就是MVC框架的`controller`，负责协调各部件完成任务
2. **model** 文件夹负责存储数据库模型和数据库操作相关的代码
3. **service** 负责处理比较复杂的业务，把业务代码模型化可以有效提高业务代码的质量（比如用户注册，充值，下单等）
4. **dto** 储存通用的json模型，把model得到的数据库模型转换成api需要的json对象 
5. **cache** 负责redis缓存相关的代码
6. **util** 一些通用的小工具
7. **cmd** 辅助开发的相关独立工具包

## 待完成

- [ ] auth权限控制
- [ ] 国际化i18n
- [ ] 其他

