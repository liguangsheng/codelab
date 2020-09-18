@set GOARCH=386
@set GOOS=windows
@rsrc -manifest qqsm.exe.manifest -o rsrc.syso
@go build -ldflags="-H windowsgui" -o bin/qqsm.exe