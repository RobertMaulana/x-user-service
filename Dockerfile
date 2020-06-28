FROM golang:alpine
RUN apk update && apk add --no-cache git
RUN adduser -D -g '' appuser
WORKDIR $GOPATH/src/user-service
COPY . .
RUN go get
RUN go build -o user-service
ENTRYPOINT ./user-service

# running in port
EXPOSE 8081

