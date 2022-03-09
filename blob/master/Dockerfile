# The base go-image
FROM golang:latest AS build

# Set working directory
WORKDIR /app

# Copy all files from the current directory to the app directory
COPY ./ ./

# Run command as described:
# go build will build an executable file named server in the current directory
RUN go build -o templog .



FROM alpine:latest as app

# Set working directory
WORKDIR /app

COPY --from=build --chown=nobody /app/templog /app/templog

# Run the server executable
CMD [ "/app/templog" ]