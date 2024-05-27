# sets up variables to be used inside dockerfile
# https://therdnotes.com/sharing-arg-in-multi-stage-dockerfile
ARG USER_NAME=docker-container-user
ARG USER_UID=6969

FROM golang:1.22.3-alpine3.19 AS build-stage
ARG USER_NAME
ARG USER_UID

RUN adduser -u ${USER_UID} --disabled-password --no-create-home ${USER_NAME}
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
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -s" -o ./pixelated-pipeline

# multistage building: using some components of the previous build step into the upcoming steps.
FROM scratch AS build-release-stage
ARG USER_NAME

# sets the root directirectory as the working directory.
WORKDIR /

# https://medium.com/@lizrice/non-privileged-containers-based-on-the-scratch-image-a80105d6d341
COPY --from=build-stage /etc/passwd /etc/passwd

# current user does not need to be root as database is shifted to cloud
# USER ${USER_NAME} # this non root user cannot access /etc/secrets, need to use groups to give access

# copies the pixelated-pipeline from the previous build stage to this one.
COPY --from=build-stage /app/pixelated-pipeline /pixelated-pipeline
COPY --from=build-stage /app/public/ /public/

# exposes port 8080 out of the container
EXPOSE 8080

# finally starts the backend.
CMD [ "/pixelated-pipeline" ]
