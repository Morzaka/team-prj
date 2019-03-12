# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from scratch (use binary code)
FROM scratch

# Add Maintainer Info
LABEL maintainer="Team-Project <https://gitlab.com/golang-lv-388/team-project>"

# Add SSL root certificates. Depending on the operating system, these certificates can be in many different places.
# This one for linux, copy the ca-certificates.crt from our running machine into project root repository.
# COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ADD ca-certificates.crt /etc/ssl/certs/

# Add main binary file
ADD main /

# This container exposes port 8080 to the outside world
EXPOSE 8081

# Run the executable
CMD ["/main"]