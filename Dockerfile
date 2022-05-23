FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/ilikedaima/gin-blog
COPY . $GOPATH/src/github.com/ilikedaima/gin-blog
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./gin-blog"]