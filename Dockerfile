FROM golang:1.12.6-alpine3.10 as build-env

RUN apk add git

RUN mkdir /banyan
WORKDIR /banyan

COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build

FROM scratch

COPY --from=build-env /go/bin/banyan /banyan

ENTRYPOINT ["/banyan"]