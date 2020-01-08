FROM alpine

WORKDIR $GOPATH/src/go-crawl-itemserver
COPY . $GOPATH/src/go-crawl-itemserver

CMD ["./go-crawl-itemserver"]

# 编译二进制文件
# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-crawl-itemserver -tags prod server.go  

# build 命令
# docker build -t go-crawl-itemserver .

# run
# docker run --link elasticsearch-server:elasticsearch-server -p 6000:6000 --name go-crawl-itemserver go-crawl-itemserver