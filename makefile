# stub makefile
test:
	cd ./libs/tokens/ && go test ./...
hooks:
	cp pre-commit .git/hooks/pre-commit
