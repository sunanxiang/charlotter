# Charlotter

微信订阅号接入chatGPT服务。

### 使用方法

1、准备一台可以访问openai 的linux服务器，安装用到的环境git、go等，然后拉取仓库
```shell
git clone https://github.com/sunanxiang/charlotter.git
```

2、准备好openAI token 以及微信开放平台上自己设置的token，在config 文件中修改相应的值。

3、
```shell
// go mod 
go mod tidy

// 直接运行
go run main.go

// 后台运行
go build .
nohup ./charlotter >log.txt 2>&1 &

```
