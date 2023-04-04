# Charlotter

微信订阅号接入chatGPT服务demo。

### 使用方法

1、准备一台可以访问openai 的linux服务器
```shell
git clone https://github.com/sunanxiang/charlotter.git
```

2、准备好openAI token 以及微信开放平台上自己设置的token，在config 文件中修改相应的值。

3、
```shell
// 直接运行
go run main.go

// 后台运行
go build .
nohup ./charlotter >log.txt 2>&1 &

```
