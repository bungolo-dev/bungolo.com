set -euxo pipefail

mkdir -p "$(pwd)/functions"
GOBIN=$(pwd)/functions go install ./...
chmod +x "$(pwd)"/functions/*
go env
cp ./public/index.html "$(pwd)"/functions/index.html"