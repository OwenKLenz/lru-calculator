FROM golang:1.16
RUN mkdir /lru-cache
ADD . /lru-cache
WORKDIR /lru-cache/src
RUN go build -o ../bin/main main.go
CMD ["/lru-cache/bin/main"]

