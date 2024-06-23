FROM golang:1.22.3-alpine3.19
WORKDIR /app
# air-verse air is a command line utility for live reload in development stage.
RUN go install github.com/air-verse/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN swag init
# https://github.com/air-verse/air/issues/274#issuecomment-1635489650
# air needs to use polling instead of fsnotify to rebuild on file changes
# when working with two different file system, i.e. windows along with docker
CMD ["air", "-c", ".air.toml"]
