FROM golang as builder
ENV GO111MODULE=off
ENV GO15VENDOREXPERIMENT=1
ENV GITPATH=chatroom_server_golang
WORKDIR /
RUN mkdir -p /go/src/${GITPATH}
COPY ./ /go/src/${GITPATH}
RUN cd /go/src/${GITPATH} && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -v
FROM alpine
ENV apk –no-cache add ca-certificates
COPY --from=builder /go/bin/chatroom_server_golang /root/chatroom
CMD ["/root/chatroom"]