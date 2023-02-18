echo "Running docker-entrypoint.sh"
reflex -r '\.go$' -s -- sh -c 'echo "Running development server" && go mod tidy && go run cmd/main.go'
