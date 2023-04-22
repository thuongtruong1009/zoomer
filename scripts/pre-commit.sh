echo "Running gofmt..."
gofmt -s -w .

echo "Running pre-commit hooks..."
./scripts/run-tests.sh

if [ $? -ne 0 ]; then
    echo "Tests must passs before commit. Pre-commit hooks failed"
    exit 1
fi
