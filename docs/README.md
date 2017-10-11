手机客户端接口
==============

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
<!-- END doctoc -->

- [接口规范](#%E6%8E%A5%E5%8F%A3%E8%A7%84%E8%8C%83)
- [用户](#%E5%B7%A5%E4%BD%9C%E5%8F%B0%E6%8E%A5%E5%8F%A3)
  - [当前用户](#%E5%88%97%E8%A1%A8%E5%88%86%E9%A1%B5)
  - [登录](#%E8%AE%BF%E9%97%AE%E5%8E%86%E5%8F%B2%E5%88%86%E9%A1%B5)
  - [注册](#%E6%90%9C%E7%B4%A2%E5%88%86%E9%A1%B5)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

接口规范
--------

---

-	默认所有请求为 HTTP + JSON 格式
-	默认编码统一为 utf-8
-	时间类型格式为 iso8601, 例如 "2014-10-01T00:00:00+08"
-	非 `2xx` 的请求结果统一处理办法：服务器将返回 JSON 格式为： `{ "message": "{对应的提示信息}" }`，如需输出消息给用户，请使用 `message` 字段。如果返回结果不是 JSON 或读取不到 `message`，则使用本地消息。
-	可分页的请求支持 `page` 和 `limit` 参数：

| 名称    | 说明       | 类型           |
|---------|------------|----------------|
| `page`  | 页码       | Integer (>= 1) |
| `limit` | 每页结果数 | Integer (>= 1) |

用户接口
----------

---

### 当前用户信息

```http
GET /user/me HTTP/1.1
Content-Type: application/x-www-form-urlencoded
```

返回内容

```http
HTTP/1.1 200 ok
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{
  "id": "0",
  "username": "my username",
  "nickname": "my nickname",
}
```
---

### 注册

```http
POST /users/create HTTP/1.1
Content-Type: application/x-www-form-urlencoded
```
发送内容

| 名称          | 范围 | 说明     |
|---------------|------|----------|
| `username`       | 必填 | 用户名     |
| `password` | 选填 | 密码     |

返回内容

```http
HTTP/1.1 200 ok
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{
  "id": "0",
  "username": "my username",
  "nickname": "my nickname",
}


### 登录

```http
POST /users/login HTTP/1.1
Content-Type: application/x-www-form-urlencoded
```
发送内容

| 名称          | 范围 | 说明     |
|---------------|------|----------|
| `username`       | 必填 | 用户名     |
| `password` | 选填 | 密码     |

返回内容

```http
HTTP/1.1 200 ok
Content-Type: application/json
Cache-Control: no-store
Pragma: no-cache

{
  "id": "0",
  "username": "my username",
  "nickname": "my nickname",
}
```

更新日志
--------

-	10/11: 初始化
