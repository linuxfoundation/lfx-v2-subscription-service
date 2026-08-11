# LFX V2 Subscription Service

LFX V2 Subscription Service

## Getting Started

1. Generate the Goa API code (this repo does not commit generated code — it must be generated once before the module will build):

   ```bash
   make apigen
   ```

   Commit the resulting `gen/` directory along with your other changes.
2. Implement your service logic in `internal/service/service.go` and wire any dependencies in `internal/container/container.go`.
