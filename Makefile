.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/kylerbot functions/kylerbot/*.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/ping functions/ping/*.go

# Install relevant packages for deployments via Serverless.
serverless-init:
	yarn add serverless serverless-domain-manager

# Remove relevant packages for deployments via Serverless.
serverless-clean:
	rm -rf node_modules package.json package.lock yarn.lock

clean:
	rm -rf ./bin ./vendor *.lock

# Deploy a clean build to AWS Lambda via Serverless.
deploy: clean build serverless-init
	sls deploy --verbose
	make serverless-clean
