## promjson

本项目提供了一个 API，用于将指定的 Prometheus 的 Metrics 数据转换成 JSON 格式。

该服务启动后会暴露 http://[IP]:[PORT]/metricsjson 路由，同时提供参数 url 用于指定需要进行转换的地址。

#### Usage

* Query JSON Metrics

```
curl -XGET http://[IP]:[PORT]/metricsjson\?url\=http://localhost:19100/metrics
```

Response:

```
{"code":0,"data":"json data here","message":"success"}
```
