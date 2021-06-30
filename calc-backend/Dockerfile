FROM golang:1.16
RUN mkdir /calc-server
ADD . /calc-server
WORKDIR /calc-server
RUN go build -o main calc.go
CMD ["/calc-server/main"]

