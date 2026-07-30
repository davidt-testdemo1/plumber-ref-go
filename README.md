# plumber-ref-go

Control fixture for the Plumber end-to-end campaign: the simplest thing that
can deploy successfully through the **Containers** preset.

- stdlib only, so the Trivy image scan has almost no surface
- binds `$PORT`, which the generated Dockerfile sets
- returns HTTP 200 at `/`, which the generated ALB target group health-checks
- `go 1.22`, matching the base image the container scaffold pins

If a deploy of this repo fails, the fault is in Plumber.
