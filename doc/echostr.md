## 服务器校验

在微信公众号后台配置好 服务地址 后，微信服务器会发送一个校验请求到该 url

![echostr](./img/echostr.jpg)

- 接收到微信 GET 请求后，通过框架提供的 `EchoStr` 可以解析出对应参数
- 如果校验成功，会输出 `echostr` 完成校验

完整用例请参见 [https://github.com/fastwego/offiaccount-demo](https://github.com/fastwego/offiaccount-demo)