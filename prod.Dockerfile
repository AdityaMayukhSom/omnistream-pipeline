FROM golang:1.22.3-alpine3.19 AS build-stage
# creates a directory inside root called app and cd into that directory.
WORKDIR /app
# copies the go.mod and go.sum from our application into that directory.
COPY go.mod go.sum ./ 
# installs the required modules from go mod file.
RUN go mod download
# copies everything else(the source code) except files mentioned in .dockerignore.
COPY . .
# compiles the code and  creates an executable file named pixelated-pipeline.
# the exe file produced is a standalone execuatable file.
RUN CGO_ENABLED=0 GOOS=linux go build -o /pixelated-pipeline

# multistage building: using some components of the previous build step into the upcoming steps.
FROM alpine:3.19 AS build-release-stage
# sets the root directirectory as the working directory.
WORKDIR /
# copies the application.yml from host system to the /etc/secrets/application.yml in the container OS
COPY application.yml /etc/secrets/application.yml
# copies the pixelated-pipeline from the previous build stage to this one.
COPY --from=build-stage /pixelated-pipeline /pixelated-pipeline
EXPOSE 8080
# USER nonroot:nonroot
# finally starts the backend.
ENTRYPOINT [ "/pixelated-pipeline" ]
