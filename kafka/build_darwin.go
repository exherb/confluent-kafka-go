// +build !dynamic


// This file was auto-generated by librdkafka/setup.sh, DO NOT EDIT.

package kafka

// #cgo CFLAGS: -I${SRCDIR}
// #cgo LDFLAGS: ${SRCDIR}/librdkafka/librdkafka_darwin.a  -lm -lsasl2 -lz -ldl -lpthread
import "C"

// LibrdkafkaLinkInfo explains how librdkafka was linked to the Go client
const LibrdkafkaLinkInfo = "static darwin from librdkafka-static-bundle-v1.1.0-selfstatic-test14.tgz"
