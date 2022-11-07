# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
FROM golang:1.18-bullseye as build-env

# install additional OS packages.
RUN apt update && \
    apt upgrade -y

RUN apt-get install -y \
    build-essential \
    ca-certificates \
    # must install cross compiler for arm64
    gcc-aarch64-linux-gnu

WORKDIR /usr/src

# Get Go dependencies
COPY go.mod ./go.mod
COPY go.sum ./go.sum
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    go mod download

# Copy rest of files
COPY . .

# compile dualityd to ARM64 architecture for final image
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    CGO_ENABLED=1 \
    CC=aarch64-linux-gnu-gcc \
    GOOS=linux \
    GOARCH=arm64 \
    go build -o build/dualityd_arm64 ./cmd/dualityd


# Final image build on small stable release of ARM64 Linux
FROM arm64v8/alpine:20220715

# Install ca-certificates
RUN apk add --update \
    # required for dualityd to work
    libc6-compat \
    # allow JSON parsing in startup shell scripts
    jq \
    # required for HTTPS to connect properly
    ca-certificates

# Install testnet utils when not on a production chain
ARG NETWORK=duality-1
RUN if [[ ! "$NETWORK" =~ "^duality-\d+$" ]]; \
    then \
        # install TOML editing tool dasel for complicated TOML edits \
        wget https://github.com/TomWright/dasel/releases/download/v1.27.3/dasel_linux_arm64.gz; \
        gzip -d dasel_linux_arm64.gz; \
        chmod 755 dasel_linux_arm64; \
        mv ./dasel_linux_arm64 /usr/local/bin/dasel; \
    fi

WORKDIR /usr/src

# Copy over binaries from the build-env
COPY --from=build-env /usr/src/build/dualityd_arm64 /usr/bin/dualityd

# Copy our configuration settings
COPY docker/fullnode/app.toml /root/.duality/config/app.toml
COPY docker/fullnode/client.toml /root/.duality/config/client.toml
COPY docker/fullnode/config.toml /root/.duality/config/config.toml

# expose all used ports
EXPOSE 26658
EXPOSE 26657
EXPOSE 26656
EXPOSE 9091
EXPOSE 9090
EXPOSE 6060
EXPOSE 1317

# add startup scripts and their dependencies
COPY scripts scripts
COPY networks networks

# default to serving the chain with default data and name
CMD ["sh", "./scripts/startup.sh"]
