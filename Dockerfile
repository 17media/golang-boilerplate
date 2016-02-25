# Base this docker container off the official golang docker image.
# Docker containers inherit everything from their base.
FROM golang:1.5.3-alpine

# Create a directory inside the container to store all our application and then make it the working directory.
RUN mkdir -p /go/src/app
WORKDIR /go/src/app

RUN apk update && apk add --no-cache git mercurial

# Copy the boilerplate directory (where the Dockerfile lives) into the container.
COPY . /go/src/app

# Download and install any required third party dependencies into the container.
RUN go get -d
RUN go install
RUN go build

RUN rm -rf /var/cache/apk/*

ENTRYPOINT ["./docker-entrypoint.sh"]

# Set the PORT environment variable inside the container
ENV PORT 8080

# Expose port 8080 to the host so we can access our application
EXPOSE 3000

EXPOSE 8080

# Now tell Docker what command to run when the container starts
CMD ["app"]
