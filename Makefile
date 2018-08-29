# go build command
gobuild:
	@go build -v -o punyasaya cmd/punyasaya/*.go

# go run command
gorun:
	make gobuild
	@./punyasaya --config_path='./config/tkp-app.{TKPENV}.yaml'

# deploy command
deploy:
	@echo "deploying"