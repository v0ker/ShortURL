# ShortURL
This is a simple URL shortener written in Go using MySQL and Redis.

## Installation
1. Clone the repository
2. Execute sql script in `data/init.sql`
3. run
```shell
> go build cmd/app
> app -config configs/config.yaml  
```
4. create a new short url
```shell
curl -X POST \
  http://127.0.0.1:8090/api/url \
  -H 'Content-Type: application/json' \
  -d '{
	"url":"https://github.com/v0ker/ShortURL",
	"ttl": 300
}'
```
4. Open the short url from action 3 in browser  
