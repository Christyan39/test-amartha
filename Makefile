# execute unit test
test: 
	go test -short -p 2 -coverprofile=cover.out `go list ./... | grep -vE '(vendor|mock)'`

init:
	go mod init
	go mod tidy
	go mod vendor
