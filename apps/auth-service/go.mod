module github.com/vizitiuRoman/auth-service

go 1.23.4

replace github.com/vizitiuRoman/libs/config => ../libs/config

replace github.com/vizitiuRoman/libs/cache => ../libs/cache

replace github.com/vizitiuRoman/libs/logger => ../libs/logger

replace github.com/vizitiuRoman/libs/pgclient => ../libs/pgclient

require (
	github.com/Pallinder/go-randomdata v1.2.0
	github.com/friendsofgo/errors v0.9.2
	github.com/gofrs/uuid v4.2.0+incompatible
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	github.com/stretchr/testify v1.9.0
	github.com/vizitiuRoman/libs/cache v0.0.0
	github.com/vizitiuRoman/libs/config v0.0.0
	github.com/vizitiuRoman/libs/logger v0.0.0
	github.com/vizitiuRoman/libs/pgclient v0.0.0
	github.com/volatiletech/null/v8 v8.1.2
	github.com/volatiletech/sqlboiler/v4 v4.15.0
	github.com/volatiletech/strmangle v0.0.6
	go.uber.org/fx v1.19.3
	go.uber.org/zap v1.27.0
	golang.org/x/crypto v0.30.0
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4
	google.golang.org/grpc v1.70.0
	google.golang.org/protobuf v1.35.2
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.4.2 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/redis/go-redis/v9 v9.6.1 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/volatiletech/inflect v0.0.1 // indirect
	github.com/volatiletech/randomize v0.0.1 // indirect
	go.uber.org/config v1.4.0 // indirect
	go.uber.org/dig v1.16.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/net v0.32.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/tools v0.21.1-0.20240508182429-e35e4ccd0d2d // indirect
	golang.org/x/xerrors v0.0.0-20220609144429-65e65417b02f // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
