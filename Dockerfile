# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# The alpine package manager use to fetch the current ca-certificates package
FROM alpine:latest as alpine
RUN apk --update add ca-certificates

# Start from scratch (use binary code)
FROM scratch

#
#ENV GOPATH /go/src/team-project

# Add Maintainer Info
LABEL maintainer="Team-Project <https://gitlab.com/golang-lv-388/team-project>"

# Add SSL root certificates.
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Add main binary file
ADD main /

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["/main"]