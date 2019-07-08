FROM golang:1.12.6-alpine3.10 as build-base

RUN apk update && apk add --no-cache git

RUN mkdir /banyan
WORKDIR /banyan

COPY go.mod .
COPY go.sum .

RUN go mod download




FROM build-base as build-env

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -o banyan




FROM scratch

COPY --from=build-env /banyan/banyan /banyan

CMD [ "/banyan" ]