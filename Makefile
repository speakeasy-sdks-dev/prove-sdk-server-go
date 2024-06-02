.PHONY: *

all: speakeasy


speakeasy: check-speakeasy
	speakeasy generate sdk --lang go -o . -s https://raw.githubusercontent.com/payfone/prove-openapi/main/orchestration.yaml

speakeasy-validate: check-speakeasy
	speakeasy validate openapi -s https://raw.githubusercontent.com/payfone/prove-openapi/main/orchestration.yaml

check-speakeasy:
	@command -v speakeasy >/dev/null 2>&1 || { echo >&2 "speakeasy CLI is not installed. Please install before continuing."; exit 1; }

tidy:
	go mod tidy