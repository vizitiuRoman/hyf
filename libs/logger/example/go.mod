module github.com/vizitiuRoman/libs/logger/example

go 1.22.3

replace github.com/vizitiuRoman/libs/logger v0.0.0-00010101000000-000000000000 => ../

require (
	github.com/vizitiuRoman/libs/logger v0.0.0-00010101000000-000000000000
	go.uber.org/fx v1.21.0
	go.uber.org/zap v1.27.0
)

require (
	go.uber.org/dig v1.17.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
)
