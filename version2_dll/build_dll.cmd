go build -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -ldflags "-w -s" -buildmode=c-shared -o Heic2Jpg.dll
upx -9 Heic2Jpg.dll