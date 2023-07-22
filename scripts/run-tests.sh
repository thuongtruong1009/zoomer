go test ./pkg/validators/... ./pkg/utils/... ./pkg/cache/... ./pkg/interceptor/... -v -coverprofile=logs/coverage -covermode=atomic

go test -race -timeout 30s ./pkg/helpers -run ^TestParallelize$
