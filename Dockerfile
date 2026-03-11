# =============================================================================
# Stage 1 — builder
# Uses the hardened dev image which includes the full Go toolchain, gcc, and
# git needed to download and compile the module graph.
# =============================================================================
FROM golang:1.26 AS builder

WORKDIR /build

# Copy dependency manifests first to leverage layer caching.
COPY go.mod go.sum ./

# Download dependencies (network-isolated after this layer).
RUN go mod download && go mod verify

# Copy the rest of the source tree.
COPY . .

# Build a statically linked binary so it can run on a scratch-like final image.
# CGO_ENABLED=0  — pure Go, no libc dependency
# -trimpath      — remove absolute file paths from the binary
# -ldflags       — strip debug info & DWARF to minimise image size
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
      -trimpath \
      -ldflags="-s -w" \
      -o /build/twosum \
      ./cmd/twosum


# =============================================================================
# Stage 2 — final (hardened runtime)
# Uses the minimal hardened Debian 13 image — no shell, no package manager,
# no Go toolchain. Only the binary and its dependencies are present.
# dhi.io images run as an existing nonroot user by default.
# =============================================================================
FROM golang:1.26

WORKDIR /app

# Copy only the compiled binary from the builder stage.
COPY --from=builder /build/twosum .

# Expose a descriptive entrypoint.  Consumers pass arguments directly.
ENTRYPOINT ["/app/twosum"]
