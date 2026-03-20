echo "Building dependencies"
docker run --rm `
    -v ${pwd}/..\api-definitions\dashboard\proto:/proto `
    -v ${pwd}\lib\generated\grpc:/grpc `
    dart:stable `
    sh -c "apt-get update -qq && apt-get install -y -qq protobuf-compiler && dart pub global activate protoc_plugin && protoc --plugin=protoc-gen-dart=/root/.pub-cache/bin/protoc-gen-dart --proto_path=/proto --dart_out=grpc:/grpc /proto/*.proto"
