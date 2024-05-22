FROM golang:1.22.3-alpine3.19
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# https://github.com/cosmtrek/air/issues/274#issuecomment-1635489650
# air needs to use polling instead of fsnotify to rebuild on file changes
# when working with two different file system, i.e. windows along with docker
CMD ["air", "-c", ".air.toml"]
