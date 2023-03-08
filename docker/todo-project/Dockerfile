FROM golang:1.19

# Pindah ke directory /app di dalam docker container
WORKDIR /app

# Copy semua file di dalam folder todo-project ke folder /app di dalam docker container
COPY . /app

RUN ["go", "build", "-o", "todo", "."]

EXPOSE 8000
ENTRYPOINT [ "./todo" ]
