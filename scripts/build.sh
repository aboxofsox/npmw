#!/bin/bash

package="npmw"
platforms=("amd64/windows" "amd64/linux" "386/windows" "386/linux")

function has_command() {
  command -v "$1" >/dev/null 2>&1
}

function exists() {
  [ -d "./bin" ]
}

if ! exists {
  mkdir "bin"
}

for platform in "${platforms[@]}"; do
    IFS="/" read -ra split <<< "$platform"
    arch="${split[0]}"
    os="${split[1]}"

    export GOARCH="$arch"
    export GOOS="$os"

    out="$package-$os-$arch"

    if [ "$os" = "windows" ]; then
        out+=".exe"
    fi

    go build -o "./bin/$out"

    if [ "$os" = "windows" ]; then
        if has_command zip; then
           zip -r "./bin/${out%.exe}.zip" "./bin/$out"
           rm "./bin/$out"
        fi
    else
        if has_command tar; then
            tar -czf "./bin/$out.tgz" "./bin/$out"
            rm "./bin/$out"
        fi
    fi
done