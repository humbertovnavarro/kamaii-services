# stub makefile
build:
	cd ./services/account-balance && go build -o ../../out/bin/account-balance

test:
	cd ./libs/tokens/ && go test ./...
	
hooks:
	cp ./scripts/pre-commit .git/hooks/pre-commit
