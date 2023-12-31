build:
	env GOOS:linux go build -o -ldflags="-s -w" -o bin/main main.go
deploy_prod : build
	serverless deploy --stage prod --aws-profile lurreyserverless