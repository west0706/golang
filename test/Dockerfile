from west0706/go_on_ubuntu:latest

RUN apt install -y git

WORKDIR $GOPATH/src 

RUN git clone https://github.com/west0706/golang.git go 

WORKDIR go/test/src/gin

RUN ["/usr/local/go/bin/go", "get", "-u", "github.com/gin-gonic/gin"]

RUN ["/usr/local/go/bin/go", "build"]

RUN mv ./* /opt

WORKDIR /opt

CMD ["./gin"]
