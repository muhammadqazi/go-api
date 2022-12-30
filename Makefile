server:
	go run src/cmd/main.go

func:
	chmod +x src/scripts/new-func.sh
	./src/scripts/new-func.sh $1