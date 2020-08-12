# Building the container image for the image.
ARG VERSION=alpine
# Sets the base image
FROM golang:${VERSION} AS baseimage
# FROM busybox:${VERSION} AS baseimage
RUN echo 'Building image...'

# Environment variable for the home directory
ENV home /server
WORKDIR ${home}
RUN pwd
ADD ./src/counter/* ./counter/
ADD ./src/user/* ./user/
ADD ./src/server.go ./server.go
ADD ./src/go.mod ./go.mod
ADD ./src/users.json ./
# COPY ./src/* .
# RUN pwd
RUN go build -o main server.go  

# FROM scratch
# COPY --from=baseimage /server/main /server/main

# ENTRYPOINT [ "/go/bin/server" ]
EXPOSE 8090
# CMD ["./main"]
ENTRYPOINT [ "/server/main" ]
