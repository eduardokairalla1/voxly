# --------- STAGE 1: Build Image ---------

# set base image
FROM golang:1.25-alpine AS build

# set working directory
WORKDIR /app

# copy dependencies files
COPY go.mod go.sum ./

# install dependencies
RUN go mod download

# copy source code and version
COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY pkg/ ./pkg/
COPY VERSION ./

# build optimized binary
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s -X main.Version=$(cat VERSION)" \
    -o reclamis ./cmd/reclamis/main.go

# --------- STAGE 2: Production Image ---------

# set production image
FROM alpine:3.23.3

# metadata
ARG VERSION
LABEL version="${VERSION}"

# set working directory
WORKDIR /app

# install runtime deps
RUN apk --no-cache add ca-certificates tzdata curl

# copy binary from build
COPY --from=build /app/reclamis .

# copy init script
COPY scripts/init ./init

# make executable
RUN chmod +x ./init ./reclamis

# expose port
EXPOSE 8080

# run app
CMD ["./init"]