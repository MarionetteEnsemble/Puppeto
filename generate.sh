#!/usr/bin/bash

SOURCE=${BASH_SOURCE[0]}
while [ -L "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
  DIR=$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )
  SOURCE=$(readlink "$SOURCE")
  [[ $SOURCE != /* ]] && SOURCE=$DIR/$SOURCE # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR=$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )
rm -rf $DIR/lib
mkdir $DIR/lib
antlr4 -long-messages -Dlanguage=Go -no-listener -visitor -package lib -o $DIR/lib -lib $DIR/grammars/ $DIR/grammars/Puppeto*.g4
find $DIR/lib -type f \
    -name '*.go' \
    -exec sed -i -e 's,github.com/antlr/antlr4/runtime/Go/antlr/v4,github.com/antlr4-go/antlr,g' {} \;
