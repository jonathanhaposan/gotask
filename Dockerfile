FROM golang:1.11

RUN mkdir -p /go/src/github.com/jonathanhaposan/gotask/
WORKDIR /go/src/github.com/jonathanhaposan/gotask/
ADD . /go/src/github.com/jonathanhaposan/gotask/

RUN go get -d -v ./...
RUN go install 

CMD ["gotask"]