FROM golang:alpine


ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY app/go.mod .
COPY app/go.sum .
RUN go mod download

COPY app .

# Build the application
RUN go build -o main .

WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .
COPY .env.docker .env
#RUN cat .env

# Export necessary port
EXPOSE 8080

# Command to run when starting the container
CMD ["/dist/main"]