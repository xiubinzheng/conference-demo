\rm -rf main
\rm -rf handler.zip
GOOS=linux GOARCH=amd64 go build -o main main.go
zip handler.zip main
aws lambda update-function-code --function-name conference-demo-go --zip-file fileb://handler.zip