FROM golang:1.10-alpine as BUILD
ADD . /go/src/wic-server
RUN go install wic-server

FROM alpine:latest
COPY --from=BUILD /go/bin/wic-server .
CMD ["./wic-server"]
