go clean -testcache

go test ./pkg/validators/... ./pkg/utils/... ./infrastructure/cache/... ./pkg/interceptor/... ./pkg/helpers/... -v -coverprofile=logs/coverage.txt -covermode=atomic

go test -timeout 30s ./pkg/helpers -run ^TestParallelize$
