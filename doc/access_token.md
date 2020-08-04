## AccessToken 管理

### 背景知识

> access_token是公众号的全局唯一接口调用凭据，公众号调用各接口时都需使用access_token。开发者需要进行妥善保存。
> 
> access_token的存储至少要保留512个字符空间。
> 
> access_token的有效期目前为2个小时，需定时刷新，重复获取将导致上次获取的access_token失效。
> 
> 公众平台的API调用所需的access_token的使用及生成方式说明：
> 
> 1、建议公众号开发者使用中控服务器统一获取和刷新access_token，其他业务逻辑服务器所使用的access_token均来自于该中控服务器，不应该各自去刷新，否则容易造成冲突，导致access_token覆盖而影响业务；
> 
> 2、目前access_token的有效期通过返回的expire_in来传达，目前是7200秒之内的值。中控服务器需要根据这个有效时间提前去刷新新access_token。在刷新过程中，中控服务器可对外继续输出的老access_token，此时公众平台后台会保证在5分钟内，新老access_token都可用，这保证了第三方业务的平滑过渡；
> 
> 3、access_token的有效时间可能会在未来有调整，所以中控服务器不仅需要内部定时主动刷新，还需要提供被动刷新access_token的接口，这样便于业务服务器在API调用获知access_token已超时的情况下，可以触发access_token的刷新流程。
> 
> 4、对于可能存在风险的调用，在开发者进行获取 access_token调用时进入风险调用确认流程，需要用户管理员确认后才可以成功获取。具体流程为：
> 
> 开发者通过某IP发起调用->平台返回错误码[89503]并同时下发模板消息给公众号管理员->公众号管理员确认该IP可以调用->开发者使用该IP再次发起调用->调用成功。
> 
> 如公众号管理员第一次拒绝该IP调用，用户在1个小时内将无法使用该IP再次发起调用，如公众号管理员多次拒绝该IP调用，该IP将可能长期无法发起调用。平台建议开发者在发起调用前主动与管理员沟通确认调用需求，或请求管理员开启IP白名单功能并将该IP加入IP白名单列表。
> 
> 公众号和小程序均可以使用AppID和AppSecret调用本接口来获取access_token。AppID和AppSecret可在“微信公众平台-开发-基本配置”页中获得（需要已经成为开发者，且帐号没有异常状态）。
> 
> **调用接口时，请登录“微信公众平台-开发-基本配置”提前将服务器IP地址添加到IP白名单中，点击查看设置方法，否则将无法调用成功。**

### 单实例服务

![single](./img/access_token_single.jpg)

当只有一个服务实例的时候，可以简化 AccessToken 刷新机制：

- 从本地 AccessToken Cache 获取
- 如果不存在 或者 已过期，那么从微信服务器刷新&更新缓存
- 本地缓存默认使用文件方式，存放在系统临时目录下，可以通过`SetAccessTokenCacheDriver` 方法修改为内存或其他方式

**这是 fastwego/offiaccount 框架的默认刷新机制**

### 多实例服务

![multi](./img/access_token_multi.jpg)

多个服务实例，如果都直接从微信服务器刷新 AccessToken ，会导致冲突

因为每次从微信服务器刷新 AccessToken 都会导致旧的 AccessToken 不可用，这样会就影响其他服务器

此时，需要引入 AccessToken Manager 中控服务，定期去微信服务器获取 AccessToken，例如：

- 启动定时任务，每 2 小时从微信服务器获取 AccessToken
- 将最新的 AccessToken 存储到 Redis 中，修改公众号实例的 AccessToken 获取机制 `SetGetAccessTokenHandler(f GetAccessTokenFunc)`，这样公众号实例每次调用微信 API 都会先从 Redis 中获取 AccessToken
- 显然，这种架构会增加网络延时 以及 引入 AccessToken Manager 单点风险，如果对服务可用性要求很高，可以考虑引入 Redis 主从等高可用架构

