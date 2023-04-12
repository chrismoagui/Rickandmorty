FROM golang:alpine as builder

ENV GO111MODULE=on


RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:3.4
RUN apk --no-cache add ca-certificates
RUN apk update && apk add bash
RUN /bin/sh -c "apk add --no-cache bash"
RUN apk --no-cache add curl jq
RUN apk --no-cache add mysql mysql-client

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
# COPY --from=builder /app/main .
# COPY --from=builder /app/.env .       

EXPOSE 8080

COPY init.sh ./init.sh

#Command to run the executable
ENTRYPOINT [ "/bin/sh", "./init.sh" ]
CMD ["./main"]

