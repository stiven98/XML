# Start from the latest golang base image
FROM golang:latest as builder
LABEL maintainer="Aca Simic <simicaca98@gmail.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o agent-service .

######## Start a new stage from scratch #######
FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /app/agent-service .
#COPY --from=builder /app/wait_for_postgres.sh
# Command to run the executable
EXPOSE 8086
CMD ["./agent-service"]