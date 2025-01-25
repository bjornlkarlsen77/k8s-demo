### Step 1: Build stage
FROM golang:1.23-alpine as builder

WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.* .
RUN go mod download

# Copy the application source code and build the binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello

### 
## Step 2: Runtime stage
FROM scratch
WORKDIR /app
# Copy only the binary from the build stage to the final image
COPY --from=builder /app/hello .

# Set the entry point for the container
CMD ["./hello"]