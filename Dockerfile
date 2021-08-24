FROM golang:1.17.0-alpine

COPY . /go/src/github.com/dirkarnez/go-generics-playground
WORKDIR /go/src/github.com/dirkarnez/go-generics-playground

# RUN go get .
# RUN go get -u github.com/dirkarnez/again@master

CMD go run -gcflags=-G=3  main.go

