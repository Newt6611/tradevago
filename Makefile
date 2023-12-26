run:
	go run cmd/main.go

binance:
	go run cmd/binance/main.go

max:
	go run cmd/max/main.go
maxbuild:
	go build -o max cmd/max/main.go

bito:
	go run cmd/bito/main.go
