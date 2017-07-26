FROM golang:1.8

ENV PROJECT $GOPATH/src/bitbucket.org/cdncache/taskmanager
WORKDIR $PROJECT

RUN go get github.com/kardianos/govendor

COPY . $PROJECT

RUN make get-dep && \
    go test $(go list ./... | grep -v /vendor/) && \
    go build -o /bin/taskmanager

EXPOSE 8082
CMD ["taskmanager"]

#### STAGE 2 ####
# Docker >= 17.05.0-ce
FROM zyfdedh/alpine:3.6

COPY --from=0 /bin/taskmanager /bin/taskmanager

EXPOSE 8082
CMD ["taskmanager"]
