FROM docker.cnb.cool/masx200/docker_mirror/golang:1.24.4-alpine-linux-amd64 AS build
ARG VERSION="dev"

# Set the working directory
WORKDIR /build
run sed -i 's#https\?://dl-cdn.alpinelinux.org/alpine#https://mirrors.tuna.tsinghua.edu.cn/alpine#g' /etc/apk/repositories
# Install git
RUN --mount=type=cache,target=/var/cache/apk \
    apk add git

env GO111MODULE=on
env export GOPROXY=https://goproxy.cn
# Build the server
# go build automatically download required module dependencies to /go/pkg/mod
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 go build -ldflags="-s -w -X main.version=${VERSION} -X main.commit=$(git rev-parse HEAD) -X main.date=$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
    -o /bin/github-mcp-server ./main.go

# Make a stage to run the app
FROM docker.cnb.cool/masx200/docker_mirror/distroless-base-debian12:latest-linux-amd64
# Set the working directory
WORKDIR /server
# Copy the binary from the build stage
COPY --from=build /bin/github-mcp-server .
# Set the entrypoint to the server binary
ENTRYPOINT ["/server/github-mcp-server"]
# Default arguments for ENTRYPOINT
CMD ["stdio"]
