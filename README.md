## http日志服务器
* 基于 [zap](https://github.com/uber-go/zap) 的http高性能日志服务器

### require
* zap
* go 1.11+

### config
* 项目目录下 `cp config.json.example config.json`
* 配置上自己的ApiToken 以及 数据库配置

### build
* 根据自己的环境，调整编译的命令
* `go build -o logServer.exe server.go`

### run 
* `./start.sh`

### client

```php
$logContent = [
    'event'=>'日志写入测试',
    'key'=>'log_write_test_key',
    'time'=>date('Y-m-d H:i:s'),
    'request'=>'{"id":"1"}',
    'response'=>'{"username":"samuel"}',
];
$logContent = [
    'value'=>json_encode($logContent, JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE),
];

$curl = curl_init();
curl_setopt_array($curl, array(
  CURLOPT_PORT => "8002",
  CURLOPT_URL => "http://laravel.suhanyu.top:8002/log",
  CURLOPT_RETURNTRANSFER => true,
  CURLOPT_ENCODING => "",
  CURLOPT_MAXREDIRS => 10,
  CURLOPT_TIMEOUT => 30,
  CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
  CURLOPT_CUSTOMREQUEST => "POST",
  CURLOPT_POSTFIELDS => http_build_query($logContent),
  CURLOPT_HTTPHEADER => array(
    "Content-Type: application/x-www-form-urlencoded",
  ),
));
$response = curl_exec($curl);
$err = curl_error($curl);
curl_close($curl);

if ($err) {
  echo "cURL Error #:" . $err;
} else {
  echo $response;
}
```

* 成功时的响应：

```json
{"status":"1","msg":"ok!"}
```

### other
* 保证当前目录可写


## thanks
* [zap](https://github.com/uber-go/zap)
