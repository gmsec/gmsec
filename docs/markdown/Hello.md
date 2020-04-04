

## [推荐查看工具](https://www.iminho.me/)

## 总览:
- [Hello]
- [Waiting to write...]

--------------------

#### 简要描述：

- [带注解路由(参考beego形式)]

#### 请求URL:

- http://localhost:8080/xxjwxc/block

#### 请求方式：

- post
- get

#### 请求参数:

- ` ReqTest ` : demo struct

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`user_name` | `是`|int32|带校验方式   |
|`password` | 否|[]string|   |
|`hellxxjwo` | 否|`main.Hello`|测试// 测试1111   |


#### 请求示例:
```
{
     "hellxxjwo": {
          "Index": 0
     },
     "password": [
          ""
     ],
     "user_name": 0
}
```

#### 返回参数说明:

- ` ReqTest ` : demo struct

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`user_name` | `是`|int32|带校验方式   |
|`password` | 否|[]string|   |
|`hellxxjwo` | 否|`main.Hello`|测试// 测试1111   |


#### 返回示例:
	
```
{
     "hellxxjwo": {
          "Index": 0
     },
     "password": [
          ""
     ],
     "user_name": 0
}
```

#### 备注:

- 带注解路由(参考beego形式)
	

--------------------
--------------------

#### 自定义类型:

#### ` main `


- ` Hello ` : ...

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`Index` | 否|int|   |




