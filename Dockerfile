# cache go module dependencies
FROM golang:1.12.6-alpine3.10 as build-base

RUN apk update && apk add --no-cache git

RUN mkdir /banyan
WORKDIR /banyan

COPY go.mod .
COPY go.sum .

RUN go mod download



# run the build
FROM build-base as build-env

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo ./cmd/banyan



# create a minimal deployment image
FROM scratch

COPY --from=build-env /banyan/banyan /banyan

CMD [ "/banyan" ]