#!/bin/bash
# dependiendo ubicación instalación del Go SDK
go run /snap/go/current/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost