# stub makefile
test:
	cd ./libs/tokens/ && go test ./...
install-hooks:
	cp pre-commit .git/hooks/pre-commit
