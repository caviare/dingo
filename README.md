# dingo
本项目采用了一系列Golang中比较流行的组件，并且构建了一些常用的功能代码，可以以本项目为基础快速搭建服务。

## 依赖

本项目已经整合了许多开发API所必要的组件：

1. [Echo](https://echo.labstack.com/): 高性能、可扩展、极简的 Go web 框架 
2. [GORM](https://echo.labstack.com/): ORM工具。本项目需要配合Mysql使用
3. [jwt-go](https://github.com/golang-jwt/jwt): JSON Web Tokens
4. [godotenv](https://github.com/joho/godotenv):  从 `.env` 加载环境变量。