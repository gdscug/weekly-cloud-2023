FROM golang:1.19

WORKDIR /app

COPY . .

RUN ["go", "build", "-o", "server", "main.go"]

EXPOSE 8000
ENTRYPOINT [ "./server" ]
