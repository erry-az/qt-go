#!/bin/bash

set -ev

$GOPATH/bin/godepgraph -horizontal -s -o github.com/erry-az/qt-go/internal/examples/showcases/wallet github.com/erry-az/qt-go/internal/examples/showcases/wallet | dot -Tpng -o depgraph.png