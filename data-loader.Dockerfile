ARG TARGETARCH
# Build the manager binary
FROM --platform=$TARGETARCH golang:1.23 as builder
ARG TARGETOS
ARG TARGETARCH

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -a -o data-loader ./cmd/data-loader

FROM python:3.12

RUN pip install --no-cache-dir "huggingface_hub[cli]"==0.24.6 modelscope==1.18.1

COPY --from=builder /workspace/data-loader /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/data-loader"]
