package main

import (
	"bufio"
	"dingo/model"
	"dingo/utils"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// 生成用户 调用后在users表生成一条用户数据
func main() {
	// 加载配置文件
	err := godotenv.Load()
	if err != nil {
		log.Fatal("配置文件加载失败。")
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入你的用户名: ")
	name, err := inputReader.ReadString('\n')
	if err != nil {
		log.Fatal("用户名创建失败")
	}
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("请输入你的密码: ")
	pwd, err := inputReader.ReadString('\n')
	if err != nil {
		log.Fatal("密码生成失败")
	}

	fmt.Printf("你的账号为：" + name)
	fmt.Printf("你的密码为：" + pwd)

	name = strings.TrimSuffix(name, "\n")
	pwd = strings.TrimSuffix(pwd, "\n")
	pwd = utils.Encrypt(pwd)
	// 连接MYSQL数据库
	model.ConnectMysql()

	user := model.User{Username: name, Password: pwd}

	result := model.DB.Create(&user) // 通过数据的指针来创建

	if result.Error != nil {
		fmt.Println("用户创建失败：" + result.Error.Error())
	} else {
		fmt.Println("用户创建成功：")
		fmt.Println(user)
	}

}
