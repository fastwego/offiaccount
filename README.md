# fastwego/offiaccount

A fast wechat offiaccount development sdk written in Golang

## 架构设计

![sdk](./doc/img/sdk.jpg)

## 演示 Demo

[https://github.com/fastwego/offiaccount-demo](https://github.com/fastwego/offiaccount-demo)

## 框架特点

### 符合直觉

作为第三方开发框架，尽可能贴合官方文档和设计，不引入新的概念，不给开发者添加学习负担

### 简洁而不过度封装

作为具体业务和微信公众号之间的中间层，专注于通道的角色：帮业务把配置/材料投递到公众号，将公众号响应透传回业务

至于 AccessToken 管理 和 消息加解密处理，框架内部完成得干净利落，不用开发者费心

### 官方文档就是最好的文档

每个接口的注释都附带官方文档的链接，让你省时省心

### 完备的单元测试

100% 覆盖每一个接口，让你每一次调用都信心满满

### 多账号支持

一套服务支持多个微信公众号账号

### 活跃的开发者社区

FastWeGo 是一整套完整的微信开发框架，包括公众号、开放平台、微信支付、企业微信、小程序、小游戏等微信服务，拥有庞大的开发者用户群体

## 参与贡献

欢迎提交 issue，一起让微信开发更快更好！

Faster we go together!