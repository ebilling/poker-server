package pb

//go:generate protoc --go_out=plugins=grpc:${GOPATH}/src -I=${GOPATH}/src/github.com/ebilling/poker-server/pb poker.proto
