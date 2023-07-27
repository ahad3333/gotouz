swag-gen:
	swag init -g api/api.go -o api/docs

swag:
	export PATH=$(go env GOPATH)/bin:$PATH
run:
	go run cmd/main.go

build:
	CGO_ENABLED=0 GOOS=darwin go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/cmd/main.go

