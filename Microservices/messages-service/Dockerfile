# FROM golang:1.15-alpine3.13
# WORKDIR /app
# RUN apk update \
#     && apk add socat \
#     && apk add build-base
# COPY go.mod .
# COPY go.sum .
# RUN go mod download
# COPY . .
# RUN go build -o messages-service
# EXPOSE 8002
# CMD ["./messages-service"]

# Start from the latest golang base image
FROM golang:latest as builder
LABEL maintainer="Aca Simic <simicaca98@gmail.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o messages-service .

######## Start a new stage from scratch #######
FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /app/messages-service .
#COPY --from=builder /app/wait_for_postgres.sh
# Command to run the executable
EXPOSE 8002
CMD ["./messages-service"]