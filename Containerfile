FROM registry.access.redhat.com/ubi9/go-toolset:1.24 as builder

COPY . .

RUN mkdir -p out && go build -buildvcs=false -o out/ollama-metrics .

# Build the final image
FROM registry.access.redhat.com/ubi9/ubi-minimal

ARG BUILD_DATE="1970-01-01T00:00:00Z"
ARG VCS_REF="unknown"
ARG VERSION="dev"

LABEL org.opencontainers.image.title="ollama-metrics" \
      org.opencontainers.image.description="Ollama metrics proxy - collects and exposes Prometheus metrics for Ollama" \
      org.opencontainers.image.url="https://github.com/clcollins/ollama-metrics" \
      org.opencontainers.image.source="https://github.com/clcollins/ollama-metrics" \
      org.opencontainers.image.revision="${VCS_REF}" \
      org.opencontainers.image.version="${VERSION}" \
      org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.vendor="clcollins" \
      org.opencontainers.image.licenses="MIT" \
      org.opencontainers.image.base.name="https://github.com/NorskHelsenett/ollama-metrics" \
      is.collins.cluster.upstream.source="https://github.com/NorskHelsenett/ollama-metrics" \
      is.collins.cluster.upstream.vendor="NorskHelsenett" \
      io.k8s.display-name="ollama-metrics" \
      io.k8s.description="Ollama metrics proxy - collects and exposes Prometheus metrics for Ollama" \
      is.collins.cluster.image.revision="${VCS_REF}" \
      is.collins.cluster.image.version="${VERSION}" \
      is.collins.cluster.image.created="${BUILD_DATE}" \
      is.collins.cluster.build.commit.id="${VCS_REF}" \
      is.collins.cluster.build.date="${BUILD_DATE}"

COPY --from=builder /opt/app-root/src/out/ollama-metrics /ollama-metrics

USER 1001
EXPOSE 8080
ENTRYPOINT ["/ollama-metrics"]
