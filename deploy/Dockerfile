FROM golang:1.19-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
COPY pkg ./pkg
ENV GOARCH=amd64
RUN go build -o /go/bin/api cmd/api/main.go
# RUN go build -o /go/bin/api cmd/api/main.go \
#     && go build -o /go/bin/migrate cmd/migrate/main.go


FROM alpine:latest  
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
EXPOSE 8080

ENV TZ=Asia/Bangkok
ENV APP_MODE=production

# COPY migrations /app/migrations
COPY --from=build /go/bin/api /app/api
# COPY --from=build /go/bin/migrate /app/migrate


CMD ["/app/api"] 
# FROM golang:1.19-alpine AS builder
# # Set version
# ARG VERSION=dev

# # Set necessary environmet variables needed for our image
# ENV GO111MODULE=on \
#     CGO_ENABLED=0 \
#     GOOS=linux \
#     GOARCH=amd64

# # Move to working directory (/build).
# WORKDIR /build

# # Copy and download dependency using go mod.
# COPY go.mod .
# COPY go.sum .
# RUN go mod download

# # Copy the code into the container.
# COPY cmd/api/main.go .
# # ADD cmd/ cmd/
# # ADD .env .env

# # Set necessary environment variables needed for our image 
# # and build the API server.
# # RUN go build -o main cmd/v1/main.go
# RUN go build -o main -ldflags=-X=main.version=${VERSION} /build/main.go

# FROM alpine:latest
# # Move to /dist directory as the place for resulting binary folder
# WORKDIR /dist/
# # Copy binary from build to dist folder
# COPY --from=builder /build .
# # Export necessary port
# EXPOSE 8080
# # Command to run when starting the container.
# ENTRYPOINT ["./main"]