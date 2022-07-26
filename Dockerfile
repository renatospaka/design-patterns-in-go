FROM golang:1.18.2-stretch

## UPDATE THE OS
RUN apt-get update && \
    apt-get install -y tzdata 

WORKDIR /app

## SET ENVIRONMENT
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ENV TZ America/Sao_Paulo

## COPY FILES
COPY . .

## TIDY THE PROJECT
RUN go mod download && \
    go mod tidy

## KEEP THE CONTAINER RUNNiNG
CMD ["tail", "-f", "/dev/null"]