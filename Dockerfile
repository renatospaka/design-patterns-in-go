FROM golang:1.18.0-stretch

## UPDATE THE OS
RUN apt-get update && \
    apt-get install build-essential golang-github-spf13-viper-dev -y  && \
    apt-get install -y tzdata 

WORKDIR /app

## SET ENVIRONMENT
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ENV TZ America/Sao_Paulo

## START A PROJECT
RUN go mod init github.com/renatospaka/design-pattern-go

# ## COPY NECESSARY FILES
# COPY go.* ./

# ## INSTALL MY STANDARD LIBRARIES 
# RUN go get github.com/spf13/viper && \
#     go get github.com/spf13/cobra && \
#     go get github.com/satori/go.uuid && \
#     go get github.com/gofiber/fiber/v2 && \
#     go get github.com/gofiber/jwt/v2 && \
#     go get github.com/gofiber/jwt/v2 && \
#     go get github.com/stretchr/testify

# ## TIDY THE PROJECT
# RUN go mod download && \
#     go mod tidy

## KEEP THE CONTAINER RUNNiNG
CMD ["tail", "-f", "/dev/null"]