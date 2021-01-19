#!/usr/bin/env bash

CAULDRON="${PWD}/cauldron"

[[ -e ${CAULDRON} ]] || { echo "${CAULDRON} missing, try 'go build' first" 2>&1; exit 1; }

TMPFILE=$(mktemp)

trap "rm ${TMPFILE}" EXIT

RED='\033[1;31m'
GREEN='\033[1;32m'
RESET='\033[0m'

EXIT_CODE=0

result() {
    if [[ "${1}" -eq 0 ]]; then
        printf "${GREEN}pass${RESET} ${2}\n"
    else
        printf "${RED}fail${RESET} ${2}\n"
        EXIT_CODE=1
    fi
}

diff <(${CAULDRON} -template testdata/hello.tmpl name=John time=morning) testdata/hello.rendered
result $? "simple argument rendering"

diff <(${CAULDRON} -template testdata/resolv.conf.tmpl \
    nameservers='["10.20.30.40", "8.8.8.8"]' domain=mydomain.com options='{"rotate": "", "timeout": "5"}') \
    testdata/resolv.conf.rendered
result $? "json argument rendering"

diff <(${CAULDRON} -template testdata/treats.tmpl -json testdata/treats.json) testdata/treats.rendered
result $? "json datasource dendering"

${CAULDRON} -template testdata/hello.tmpl -file "${TMPFILE}" name=John time=morning \
 && diff "${TMPFILE}" testdata/hello.rendered
result $? "render to file"

exit $EXIT_CODE
