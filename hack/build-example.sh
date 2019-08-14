#!/bin/bash

commitSHA=$(git describe --dirty --always)
dateStr=$(date +%s)

go build -ldflags "-X main.commit=${commitSHA} -X main.date=${dateStr}" -o ./example/example ./example/