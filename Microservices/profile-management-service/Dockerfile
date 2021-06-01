FROM golang:latest as builder
LABEL maintainer="Aleksandar Stevanovic <stevanovicaleksa98@gmail.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o profile-management-service .
######## Start a new stage from scratch #######
FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /app/profile-management-service .
#COPY --from=builder /app/wait_for_postgres.sh
# Command to run the executable
EXPOSE 8087
CMD ["./profile-management-service"]