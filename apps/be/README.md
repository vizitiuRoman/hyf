## .env Example

```dotenv
LOCATION=remote
####################################################
###################### LOGGER ######################
LOGGER_LEVEL=debug
LOGGER_ENCODING=console
# zapcore.LevelEncoder.UnmarshalText : capital | capitalColor | color | lower
LOGGER_LEVEL_ENCODER=lower
####################################################
##################### DATABASE #####################
# LOCAL
DB_DSN=postgres://hyf:hyf@hyf_db:5432/hyf?sslmode=disable
####################################################
##################### CACHE #####################
REDIS_SENTINEL_ADDRESSES=hyf_redis:6379
REDIS_PASSWORD=hyf
```

## Run application

```bash
$ docker-compose -f ./var/docker/docker-compose.yaml
```

And if you see the following output then services is started ok.
```bash
Running 3/3
  Container docker-app      Started                                                                                                                6.3s
```

## Swagger API

swagger-ui looks like this:
![Demo-Api](swagger-ui.png)