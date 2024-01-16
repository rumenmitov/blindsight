FROM docker.io/library/golang

WORKDIR /app

COPY . .

RUN go test ./src

RUN go install ./src

ENTRYPOINT /go/bin/src

EXPOSE 3000
