
    go mod init redirect
    go mod tidy
    CGO_ENABLED=0 go build -trimpath -ldflags "-s -w"
    ./redirect
