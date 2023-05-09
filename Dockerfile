# your-repo/Dockerfile
FROM golang:1.18
WORKDIR /app
COPY . . 
RUN go build -o /bin/app .
ENTRYPOINT ["/bin/app"]