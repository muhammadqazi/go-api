server:
	go run src/cmd/main.go

fast:
	./air

func:
	chmod +x src/scripts/new-func.sh
	./src/scripts/new-func.sh $(name)

check-db:
	chmod +x src/scripts/check-db.sh
	./src/scripts/check-db.sh