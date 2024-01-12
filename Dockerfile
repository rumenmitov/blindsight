FROM docker.io/library/golang

WORKDIR /app

COPY . .

RUN go get ./src

EXPOSE 3000

CMD ["go", "run", "./src"]
