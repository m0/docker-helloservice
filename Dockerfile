FROM golang:1.7.6
WORKDIR /go/src/
COPY helloservice.go .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-s' helloservice.go

FROM scratch
MAINTAINER mo@oclab.net
COPY --from=0 /go/src/helloservice /helloservice
EXPOSE 8080
CMD [ "/helloservice" ]
