FROM golang:1.19.0

## UPDATE THE OS
RUN apt-get update && \
    go install -v golang.org/x/tools/gopls@latest && \
    apt-get install -y tzdata 

## SET ENVIRONMENT
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ENV TZ America/Sao_Paulo

## COPY MOD FILES
WORKDIR /go/src
COPY go.mod go.sum ./

## VERIFY AND TIDY THE PROJECT
RUN go mod download && \
    go mod verify && \
    go mod tidy

## COPY NECESSARY FILES
COPY . .

## KEEP THE CONTAINER RUNNiNG
CMD ["tail", "-f", "/dev/null"]