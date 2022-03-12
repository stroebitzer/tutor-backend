FROM golang:alpine3.15 as builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o tutor-backend

FROM alpine:3.15

# install kubectl
# TODO in a versioned way of doing things!!!
RUN apk --no-cache add curl
RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
RUN install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

WORKDIR /app
COPY --from=builder /src/tutor-backend /app/tutor-backend

ENTRYPOINT [ "./tutor-backend" ]
