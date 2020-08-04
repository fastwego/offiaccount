## 消息/事件 处理

> 当普通微信用户向公众账号发消息时，微信服务器将POST消息的XML数据包到开发者填写的URL上。
>
>在微信用户和公众号产生交互的过程中，用户的某些操作会使得微信服务器通过事件推送的形式通知到开发者在开发者中心处设置的服务器地址，从而开发者可以获取到该信息。
>

![message](./img/message.jpg)

- 接收到微信推送过来的消息/事件后，通过框架提供的 `ParseXML` 可以解析出对应的消息/事件类型
- 框架会自动根据微信提供的参数判断是否开启了加密功能，如有则根据配置的 `EncodingAESKey` 解密消息
- 开发者可以根据获取的消息/事件类型，完成具体的业务逻辑
- 如果需要即时回复用户文本/语音/图文等消息，构造相应的回复消息类型后，通过框架提供的 `Response` 方法输出内容
- 框架会自动根据微信提供的参数判断是否开启了加密功能，如有则根据配置的 `EncodingAESKey` 加密消息后输出给微信
- 如果不需要即时回复用户消息，`Response` 会自动回复 `success` 告知微信服务器正常响应

完整用例请参见 [https://github.com/fastwego/offiaccount-demo](https://github.com/fastwego/offiaccount-demo)

### 消息类型

```go
{{#include ../type/type_message/type_message.go}}
```

### 回复消息类型

```go
{{#include ../type/type_message/type_reply_message.go}}
```

### 事件类型

```go
{{#include ../type/type_event/type_event.go}}
```

```go
{{#include ../type/type_event/type_card_event.go}}
```

```go
{{#include ../type/type_event/type_guide_event.go}}
```

```go
{{#include ../type/type_event/type_menu_event.go}}
```

```go
{{#include ../type/type_event/type_template_msg_event.go}}
```

```go
{{#include ../type/type_event/type_verify_event.go}}
```