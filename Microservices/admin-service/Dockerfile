FROM golang:latest as builder
LABEL maintainer="Aca Simic <simicaca9898@gmail.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o admin-service .
######## Start a new stage from scratch #######
FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /app/admin-service .
#COPY --from=builder /app/wait_for_postgres.sh
# Command to run the executable
EXPOSE 8089
CMD ["./admin-service"]