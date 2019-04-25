######## First stage: build the executable #######
# Accept the Go version for the image to be set as a build argument.
# Default to Go 1.12.4
ARG GO_VERSION=1.12.4

# Build the executable.
FROM golang:${GO_VERSION}-alpine AS builder

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
# Git is required for fetching the dependencies.
RUN apk add --no-cache ca-certificates git

# Add Maintainer Info
LABEL maintainer="team-project <https://gitlab.com/golang-lv-388/team-project>"

# Set the Current Working Directory inside the container
WORKDIR /team-project

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY ./go.mod ./go.sum ./
RUN go mod download

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o team-project .

######## Second stage: from scratch #######
FROM scratch AS final

WORKDIR /root/

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /team-project .

# Declare the port on which the webserver will be exposed.
EXPOSE 8080

# Run the compiled binary.
ENTRYPOINT ["./team-project"]