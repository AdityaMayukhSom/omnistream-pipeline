FROM golang:1.22.3-alpine3.19 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /pixelated-pipeline

FROM alpine:3.19 AS build-release-stage
WORKDIR /
COPY application.yml /etc/secrets/application.yml
COPY --from=build-stage /pixelated-pipeline /pixelated-pipeline
# USER nonroot:nonroot
ENTRYPOINT [ "/pixelated-pipeline" ]
