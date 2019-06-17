FROM scratch

WORKDIR $GOPATH/src/github.com/tiancai110a/gin-blog
COPY  . $GOPATH/src/github.com/tiancai110a/gin-blog

EXPOSE 8000

ENTRYPOINT ["./gin-blog"]
