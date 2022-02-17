#!/bin/bash

GEN_PATH="."
GOPATH=$(go env GOPATH)

GROUPS_VERSION="master"
USERS_VERSION="master"
ROBOTS_VERSION="master"
ROBOT_HEALTH_VERSION="master"
EVENTS_DATA_VERSION="master"

AUTHORIZATIONS_VERSION="master"
curl -o proto/authorizations.proto "https://raw.githubusercontent.com/slavayssiere-spoon/authorizations/$AUTHORIZATIONS_VERSION/proto/authorizations.proto"

GROUPS_VERSION="master"
curl -o proto/groups.proto "https://raw.githubusercontent.com/slavayssiere-spoon/groups/$GROUPS_VERSION/proto/groups.proto"

USERS_VERSION="master"
curl -o proto/users.proto "https://raw.githubusercontent.com/slavayssiere-spoon/users/$USERS_VERSION/proto/users.proto"

ROBOTS_VERSION="master"
curl -o proto/robots.proto "https://raw.githubusercontent.com/slavayssiere-spoon/robots/$ROBOTS_VERSION/proto/robots.proto"

HEALTH_VERSION="master"
curl -o proto/health.proto "https://raw.githubusercontent.com/slavayssiere-spoon/health/$HEALTH_VERSION/proto/health.proto"

EVENT_VERSION="master"
curl -o proto/IdentityEvent.proto "https://raw.githubusercontent.com/slavayssiere-spoon/event-data/$HEALTH_VERSION/proto/IdentityEvent.proto"

mkdir -p proto/google/api 
curl -o proto/google/api/annotations.proto "https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto"
curl -o proto/google/api/http.proto "https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto"

mkdir -p proto/google/protobuf
curl -o proto/google/protobuf/descriptor.proto "https://raw.githubusercontent.com/protocolbuffers/protobuf/master/src/google/protobuf/descriptor.proto"
curl -o proto/google/protobuf/empty.proto "https://raw.githubusercontent.com/protocolbuffers/protobuf/master/src/google/protobuf/empty.proto"
curl -o proto/google/protobuf/struct.proto "https://raw.githubusercontent.com/protocolbuffers/protobuf/master/src/google/protobuf/struct.proto"
curl -o proto/google/protobuf/timestamp.proto "https://raw.githubusercontent.com/protocolbuffers/protobuf/master/src/google/protobuf/timestamp.proto"


echo "gen bq_field"
protoc \
  -I proto \
  -I $GOPATH/src/include \
  --go_out=$GEN_PATH      --go_opt=paths=source_relative \
  proto/bq_field.proto

echo "gen bq_table"
protoc \
  -I proto \
  -I $GOPATH/src/include \
  --go_out=$GEN_PATH      --go_opt=paths=source_relative \
  proto/bq_table.proto


echo "protocs"

protoc \
  -I proto \
  -I $GOPATH/src/include \
  --go_out=$GEN_PATH      --go_opt=paths=source_relative \
  --bq-schema_out=. \
  proto/data.proto

protoc \
  -I proto \
  -I $GOPATH/src/include \
  --go_out=$GEN_PATH      --go_opt=paths=source_relative \
  --bq-schema_out=. \
  proto/ActionWithInteractionResponse.proto

protoc \
  -I proto \
  -I $GOPATH/src/include \
  --go_out=$GEN_PATH      --go_opt=paths=source_relative \
  --bq-schema_out=. \
  proto/FormReply.proto

protoc \
  -I proto \
  -I $GOPATH/src/include \
  --go_out=$GEN_PATH      --go_opt=paths=source_relative \
  --bq-schema_out=. \
  proto/PersonData.proto

protoc \
  -I proto \
  -I $GOPATH/src/include \
  --go_out=$GEN_PATH      --go_opt=paths=source_relative \
  --bq-schema_out=. \
  proto/ScenarioData.proto

protoc \
  -I proto \
  -I $GOPATH/src/include \
  --go_out=$GEN_PATH      --go_opt=paths=source_relative \
  --bq-schema_out=. \
  proto/SoundData.proto

protoc \
  -I proto \
  -I $GOPATH/src/include \
  --go_out=$GEN_PATH      --go_opt=paths=source_relative \
  --bq-schema_out=. \
  proto/TouchData.proto

protoc \
  -I proto \
  -I $GOPATH/src/include \
  --go_out=$GEN_PATH      --go_opt=paths=source_relative \
  --bq-schema_out=. \
  proto/InteractionState.proto

protoc \
  -I proto \
  -I $GOPATH/src/include \
  --go_out=$GEN_PATH      --go_opt=paths=source_relative \
  --bq-schema_out=. \
  proto/RobotSpeech.proto

protoc \
  -I proto \
  -I $GOPATH/src/include \
  --go_out=$GEN_PATH      --go_opt=paths=source_relative \
  --bq-schema_out=. \
  proto/GenericData.proto

protoc-go-inject-tag -input=./data.pb.go
protoc-go-inject-tag -input=./ActionWithInteractionResponse.pb.go
protoc-go-inject-tag -input=./FormReply.pb.go
protoc-go-inject-tag -input=./PersonData.pb.go
protoc-go-inject-tag -input=./ScenarioData.pb.go
protoc-go-inject-tag -input=./SoundData.pb.go
protoc-go-inject-tag -input=./TouchData.pb.go
protoc-go-inject-tag -input=./InteractionState.pb.go
protoc-go-inject-tag -input=./RobotSpeech.pb.go
protoc-go-inject-tag -input=./GenericData.pb.go

echo "mod tidy"
go mod tidy

echo "update"
go get -u ./...

echo "build"
go build

# export DEBUG="true"
# export PROJECT_ID=""
# export PUBSUB_SUB=""
# export DATASET_NAME=""
# export SCHEMA_FILE_PATH=""
# export PORT_HEALTH="8081"
# export REDIS_HOST=""
# export REDIS_PORT=6379

