server:
	nodemon --watch './**/*go' --signal STGTERM --exec APP_ENV=dev 'go' run cmd/hub_logging/main.go

run:
	APP_ENV=dev 'go' run cmd/hub_logging/main.go
