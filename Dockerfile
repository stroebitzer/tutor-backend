FROM golang:alpine3.16 as builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o tutor-backend

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /src/tutor-backend ./tutor-backend
ENTRYPOINT [ "./tutor-backend" ]
