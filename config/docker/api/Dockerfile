FROM golang:1.24

# Set work directory
WORKDIR /go/src/app

# Copy the source code
COPY . .

# Port
EXPOSE 8000

# Build Go app
RUN go build -o main cmd/main.go

# Run the executable
CMD ["./main"]