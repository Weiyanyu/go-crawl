FROM alpine

WORKDIR $GOPATH/src/go-crawl-engine
COPY . $GOPATH/src/go-crawl-engine

CMD ["./go-crawl-engine"]

# 编译二进制文件
# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-crawl-engine -tags prod main.go  

# build 命令
# docker build -t go-crawl-engine .

# run
# docker run --link go-crawl-workerserver:go-crawl-workerserver --link go-crawl-itemserver:go-crawl-itemserver --name go-crawl-engine go-crawl-engine