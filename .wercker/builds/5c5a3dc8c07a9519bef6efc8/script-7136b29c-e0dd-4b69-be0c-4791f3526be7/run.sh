set -e
cd $WERCKER_SOURCE_DIR
export GO111MODULE=on
go build ./...
