module github.com/vizitiuRoman/libs/config/example

go 1.23

toolchain go1.23.1

replace github.com/vizitiuRoman/libs/config v0.0.0-20230718194319-d2cd082fd377 => ../

require (
	github.com/vizitiuRoman/libs/config v0.0.0-20230718194319-d2cd082fd377
	go.uber.org/fx v1.22.1
	go.uber.org/zap v1.26.0
)

require (
	github.com/joho/godotenv v1.5.1 // indirect
	go.uber.org/config v1.4.0 // indirect
	go.uber.org/dig v1.17.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.15.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
