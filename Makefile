server:
	nodemon --watch './**/*go' --signal STGTERM --exec APP_ENV=dev 'go' run main.go
