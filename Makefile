prebuild:
	go vet ./... \
	&& golint -set_exit_status $(go list ./...) \
	&& go test -short ./...