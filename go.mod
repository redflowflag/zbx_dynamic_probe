module yanfeng.com/yfprobe

go 1.21.6

require (
	github.com/wenzhenxi/gorsa v0.0.0-20230530123828-0320cce15d81 // indirect
	github.com/yuin/gopher-lua v1.1.1 // indirect
	yfcrypt v0.0.0
	cmder v0.0.0
)

replace (
    yfcrypt => ./yfcrypt
	cmder => ./cmder
)

