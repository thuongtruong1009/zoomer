go clean -testcache && go test ./pkg/validators/... ./pkg/utils/... ./pkg/cache/... ./pkg/interceptor/... ./pkg/helpers/... -v -coverprofile=logs/coverage.txt -covermode=atomic

go test -timeout 30s ./pkg/helpers -run ^TestParallelize$
