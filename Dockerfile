FROM golang:1.11

ENV APP_DIR=$GOPATH/src/github.com/jonathanhaposan/gotask
RUN mkdir -p ${APP_DIR}
COPY . ${APP_DIR}

WORKDIR ${APP_DIR}

RUN go get -d -v ./...
RUN go get gopkg.in/DATA-DOG/go-sqlmock.v1
RUN cd ${APP_DIR}/cmd/web && go build
RUN cd ${APP_DIR}/cmd/tcp && go build