#!/bin/bash
export TEST_CONFIG_PATH=$(pwd)/tests/configs/config-test.yaml
export CONFIG_FILE_PATH=$(pwd)/configs/config-dev.yaml
export RBAC_FILE_PATH=$(pwd)/tests/configs/rbac.yaml
go test -v ./... > ./test.log
fails=$(grep FAIL: test.log | wc -l)
failedOnes=$(grep FAIL: test.log)
passes=$(grep PASS: test.log | wc -l)
printf "\nTEST SUMMARY\nFAILS: %d\nPASS: %d\n" $fails $passes
if [ "$fails" -ne "0" ]; then
  tests=""
  for TEST in $failedOnes
  do
    if [[ $TEST == Test* ]]; then
    tests=$tests"\n"$TEST
    fi
  done
  printf "\nFAILED TESTS: $tests\n"
fi
rm -f test.log