#!/bin/bash
export TEST_CONFIG_PATH=$(pwd)/configs/config-test.yaml
for test_file in $(find ../ -name *_test.go); do go test -v $(dirname $test_file); done