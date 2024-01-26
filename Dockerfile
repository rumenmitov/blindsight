FROM docker.io/library/golang

WORKDIR /app

COPY . .

RUN go test .

RUN go install .

ENTRYPOINT /go/bin/blindsight

EXPOSE 3000
