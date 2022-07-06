# stub makefile
test:
	cd ./libs/tokens/ && go test ./...
hooks:
	cp ./scripts/pre-commit .git/hooks/pre-commit
