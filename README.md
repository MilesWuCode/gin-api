# gin-api

## 套件

- framework

  - https://github.com/gin-gonic/gin
  - https://github.com/kataras/iris - 參考
  - https://github.com/go-chi/chi - 參考

- orm

  - https://github.com/go-gorm/gorm
  - https://github.com/go-gorm/dbresolver
  - https://entgo.io
  - https://atlasgo.io

- database

  - https://github.com/go-gorm/mysql
  - https://github.com/go-gorm/sqlite

- log

  - https://github.com/uber-go/zap
  - https://github.com/sirupsen/logrus - 參考

- env

  - https://github.com/spf13/viper

- validator

  - https://github.com/go-playground/validator
  - https://github.com/thedevsaddam/govalidator - 參考

- storage

  - https://github.com/spf13/afero
  - https://github.com/minio/minio-go
  - https://github.com/googleapis/google-cloud-go
  - https://github.com/aws/aws-sdk-go

- cache (route, query, ...etc)

  - https://github.com/gin-contrib/cache
  - https://github.com/chenyahui/gin-cache - 參考

- dependency injection

  - https://github.com/google/wire - 參考
  - https://github.com/uber-go/dig - 參考

- jwt

  - https://github.com/golang-jwt/jwt

- auth

  - https://github.com/go-oauth2/oauth2 - Server
  - https://github.com/golang/oauth2 - Client
  - https://github.com/volatiletech/authboss - 參考

- Role Based Access Control (RBAC), roles & permissions

  - https://github.com/harranali/authority

- crypto

  - https://github.com/golang/crypto

- event

  - https://github.com/ThreeDotsLabs/watermill

- graphql

  - https://github.com/99designs/gqlgen

## 待開發

- roles & permissions
- file manager
- event
- graphql - 另開專案
- grpc - 另開專案
- ent - 另開專案
- atlas

## 學習知識

- https://github.com/open-policy-agent/opa

## Docker

- redis

```sh
docker run --name redis \
-p 6379:6379 \
-v ${PWD}:/data \
--restart unless-stopped \
-d redis:latest
```

- mysql

```sh
docker run --name mysql -p 3306:3306 \
-v ${PWD}:/var/lib/mysql \
-e MYSQL_ROOT_PASSWORD=password \
-e MYSQL_DATABASE=gin-api \
--restart unless-stopped \
-d mysql:latest
```
