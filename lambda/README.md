# Go Lambda

## steps
go mod init main
go mod tidy
go get github.com/aws/aws-lambda-go/lambda
GOOS=linux GOARCH=amd64 go build -o main
