# FROM golang:1.12.0-alpine3.9
# RUN mkdir /app
# ADD . /app
# WORKDIR /app
# ## Add this go mod download command to pull in any dependencies
# RUN go mod download
# ## Our project will now successfully build with the necessary go libraries included.
# RUN go build -o main .
# ## Our start command which kicks off
# ## our newly created binary executable
# CMD ["/app/main"]

# build stage
FROM golang:1.14-buster as builder

WORKDIR /app

COPY . .
# COPY go.mod .
# COPY go.sum .

RUN go mod download
# RUN install -v -mod=vendor ./cmd/...

RUN go build -o main ./cmd

# Run stage
FROM debian:buster

RUN apt-get update --fix-missing && \
apt-get install -yqq ca-certificates

COPY --from=builder /app .

# Copy the Pre-built binary file from the previous stage
# COPY --from=builder /app/main .

# RUN chmod +x ./main

# final stage
# FROM scratch
# COPY --from=builder /app/main /app/
EXPOSE 8080
CMD ["./main"]