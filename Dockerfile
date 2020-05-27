FROM golang:1.14 as builder
RUN mkdir /build 
ADD . /build/
# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .
# Download all the dependencies
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]