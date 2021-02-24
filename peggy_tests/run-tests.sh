#!/bin/bash
TEST_TYPE=$1
set -eu

# Run test entry point script
docker exec peggy_test_instance /bin/sh -c "pushd /shentu/ && tests/container-scripts/integration-tests.sh 1 $TEST_TYPE"