#!/usr/bin/env bash

set -euo pipefail

EXISTING_GO_GENERATE_COMMENTS="
pkg/ccl/sqlproxyccl/throttler/service.go://go:generate mockgen -package=throttler -destination=mocks_generated.go -source=service.go . Service
pkg/ccl/sqlproxyccl/denylist/service.go://go:generate mockgen -package=denylist -destination=mocks_generated.go -source=service.go . Service
pkg/ccl/sqlproxyccl/tenant/directory.go://go:generate mockgen -package=tenant -destination=mocks_generated.go . DirectoryClient,Directory_WatchEndpointsClient
pkg/cmd/roachprod/vm/aws/config.go://go:generate go-bindata -mode 0600 -modtime 1400000000 -pkg aws -o embedded.go config.json old.json
pkg/cmd/roachprod/vm/aws/config.go://go:generate gofmt -s -w embedded.go
pkg/cmd/roachprod/vm/aws/config.go://go:generate goimports -w embedded.go
pkg/cmd/roachprod/vm/aws/config.go://go:generate terraformgen -o terraform/main.tf
pkg/geo/wkt/wkt.go://go:generate sh generate.sh
pkg/kv/kvserver/concurrency/lock_table.go://go:generate ../../../util/interval/generic/gen.sh *lockState concurrency
pkg/kv/kvserver/spanlatch/manager.go://go:generate ../../../util/interval/generic/gen.sh *latch spanlatch
pkg/roachpb/batch.go://go:generate go run -tags gen-batch gen/main.go
pkg/security/certmgr/cert.go://go:generate mockgen -package=certmgr -destination=mocks_generated.go -source=cert.go . Cert
pkg/cmd/roachtest/prometheus/prometheus.go://go:generate mockgen -package=prometheus -destination=mock_generated.go -source=prometheus.go . cluster
pkg/security/securitytest/securitytest.go://go:generate go-bindata -mode 0600 -modtime 1400000000 -pkg securitytest -o embedded.go -ignore README.md -ignore regenerate.sh test_certs
pkg/security/securitytest/securitytest.go://go:generate gofmt -s -w embedded.go
pkg/security/securitytest/securitytest.go://go:generate goimports -w embedded.go
pkg/server/api_v2.go://go:generate swagger generate spec -w . -o ../../docs/generated/swagger/spec.json --scan-models
pkg/sql/conn_fsm.go://go:generate ../util/fsm/gen/reports.sh TxnStateTransitions stateNoTxn
pkg/sql/opt/optgen/lang/expr.go://go:generate langgen -out expr.og.go exprs lang.opt
pkg/sql/opt/optgen/lang/expr.go://go:generate langgen -out operator.og.go ops lang.opt
pkg/sql/schemachanger/scop/backfill.go://go:generate go run ./generate_visitor.go scop Backfill backfill.go backfill_visitor_generated.go
pkg/sql/schemachanger/scop/mutation.go://go:generate go run ./generate_visitor.go scop Mutation mutation.go mutation_visitor_generated.go
pkg/sql/schemachanger/scop/validation.go://go:generate go run ./generate_visitor.go scop Validation validation.go validation_visitor_generated.go
pkg/util/interval/generic/doc.go:  //go:generate ../../util/interval/generic/gen.sh *latch spanlatch
pkg/util/interval/generic/example_t.go://go:generate ./gen.sh *example generic
pkg/util/log/channels.go://go:generate go run gen/main.go logpb/log.proto channel.go channel/channel_generated.go
pkg/util/log/channels.go://go:generate go run gen/main.go logpb/log.proto log_channels.go log_channels_generated.go
pkg/util/log/channels.go://go:generate go run gen/main.go logpb/log.proto logging.md ../../../docs/generated/logging.md
pkg/util/log/channels.go://go:generate go run gen/main.go logpb/log.proto severity.go severity/severity_generated.go
pkg/util/timeutil/zoneinfo.go://go:generate go run gen/main.go
"

EXISTING_BROKEN_TESTS_IN_BAZEL="
pkg/acceptance/BUILD.bazel
pkg/cmd/cockroach-oss/BUILD.bazel
pkg/cmd/github-post/BUILD.bazel
pkg/cmd/prereqs/BUILD.bazel
pkg/cmd/publish-artifacts/BUILD.bazel
pkg/cmd/roachtest/BUILD.bazel
pkg/cmd/teamcity-trigger/BUILD.bazel
"

git grep 'go:generate stringer' pkg | while read LINE; do
    dir=$(dirname $(echo $LINE | cut -d: -f1))
    type=$(echo $LINE | grep -o -- '-type[= ][^ ]*' | sed 's/-type[= ]//g' | awk '{print tolower($0)}')
    build_out=$(bazel query --output=build "//$dir:${type}_string.go")
    if [[ -z "$build_out" ]]; then
        echo 'Detected an autogenerated file that is not built inside the Bazel sandbox: '
        echo "  $dir/${type}_string.go, generated by: $LINE"
        echo 'Generate this file using the Bazel sandbox (see the utilities in build/STRINGER.bzl);'
        exit 1
    fi
done

# We exclude stringer and add-leaktest.sh -- the former is already all
# Bazelfied, and the latter can be safely ignored since we have a lint to check
# the same thing: https://github.com/cockroachdb/cockroach/issues/64440
git grep '//go:generate' -- './*.go' | grep -v stringer | grep -v 'add-leaktest\.sh' | while read LINE; do
    if [[ "$EXISTING_GO_GENERATE_COMMENTS" == *"$LINE"* ]]; then
	# Grandfathered.
	continue
    fi
    echo 'Detected an unknown //go:generate comment:'
    echo "$LINE"
    echo 'Please ensure that the equivalent logic to generate these files is'
    echo 'present in the Bazel build as well, then add the line to the'
    echo 'EXISTING_GO_GENERATE_COMMENTS in build/bazelutil/check-genfiles.sh.'
    echo 'Also see https://cockroachlabs.atlassian.net/wiki/spaces/CRDB/pages/1380090083/How+to+ensure+your+code+builds+with+Bazel'
    exit 1
done

git grep 'broken_in_bazel' pkg | grep BUILD.bazel: | grep -v pkg/BUILD.bazel | grep -v generate-test-suites | cut -d: -f1 | while read LINE; do
    if [[ "$EXISTING_BROKEN_TESTS_IN_BAZEL" == *"$LINE"* ]]; then
	# Grandfathered.
	continue
    fi
    echo "A new broken test in Bazel was added in $LINE"
    echo 'Ensure the test runs with Bazel, then remove the broken_in_bazel tag.'
    exit 1
done
