FROM golang:1.14 as builder
RUN mkdir /build 
ADD . /build/
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]