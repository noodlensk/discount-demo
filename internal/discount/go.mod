module github.com/noodlensk/discount-demo/internal/discount

go 1.17

require (
	github.com/deepmap/oapi-codegen v1.9.0
	github.com/go-chi/chi/v5 v5.0.5
	github.com/go-chi/render v1.0.1
	github.com/google/uuid v1.3.0
	github.com/noodlensk/discount-demo/internal/common v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
	go.uber.org/zap v1.19.1
)

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
)

replace github.com/noodlensk/discount-demo/internal/common => ../common/
