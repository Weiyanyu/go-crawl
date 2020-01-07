FROM alpine

WORKDIR $GOPATH/src/go-crawl-workerserver
COPY . $GOPATH/src/go-crawl-workerserver

CMD ["./go-crawl-workerserver"]

# 编译二进制文件
# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-crawl-workerserver server.go  

# build 命令
# docker build -t go-crawl-workerserver .

# run
# docker run -p 9000:9000 --name go-crawl-workerserver go-crawl-workerserver ./go-crawl-workerserver -port=9000