微信推送机器人

非常简单的实现，只有一个依赖。
上手简单，小白也能轻松使用。

依赖：https://github.com/tans/push-bot

效果上图：

使用方法：
1.微信关注：推送精灵
二维码：

2.创建企业微信群，邀请推送精灵进群

3.在群内发送一条消息，获取到token

4.替换common.go中的RoomUrl

5.go run main.go 或者go build