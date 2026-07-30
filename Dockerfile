# syntax=docker/dockerfile:1
# Tudovu-generated. Multi-stage, distroless runtime, non-root.
FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
# Fully static binary so the runtime image needs no libc or OS packages.
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/server .

# Distroless static: no shell, no package manager, no OS packages — the Trivy
# image scan has nothing to flag. Ships CA certs + a built-in nonroot user.
FROM gcr.io/distroless/static-debian12:nonroot AS runtime
WORKDIR /app
COPY --from=build /out/server /app/server
USER nonroot:nonroot
ENV PORT=3000
EXPOSE 3000
# No Docker HEALTHCHECK: distroless has no shell/wget. Your platform health-checks
# /healthz instead (the ECS/ALB target group, or `docker run` liveness).
ENTRYPOINT ["/app/server"]
