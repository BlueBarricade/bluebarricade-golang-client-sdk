# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="li@bluebarricade.com"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN make -f makefile .
RUN ls -al

######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

RUN mkdir -p /root/config/
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/dist/ /root/
COPY --from=builder /app/dist/conf.local.yaml /root/config/

# Expose port 8080 to the outside world
EXPOSE 8080

ENV COMPUTOP_PASSWORD=none
ENV ENVIRONMENT=development

VOLUME /root/config
VOLUME /var/log

RUN apk add --no-cache bash

# Command to run the executable
ENTRYPOINT [ "/bin/bash", "./start-server.sh" ]
CMD ["-p", "/root/config/conf.local.yaml"] 