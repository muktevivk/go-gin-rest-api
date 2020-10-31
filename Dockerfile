FROM golang as build 

ENV GO111MODULE=on

WORKDIR /todo

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN mkdir src

COPY src src

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./src/app

FROM scratch

COPY --from=build /todo/app /todo/

EXPOSE 8080
ENTRYPOINT ["/todo/app"]