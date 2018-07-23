#!/bin/bash
#
# A smoke test script for cauldron.
#
# @author John DiStasio <jndistasio@gmail.com>
# @copyright Copyright (c) 2017 John DiStasio
#
CAULDRON=build/cauldron

[[ -e ${CAULDRON} ]] || { echo "${CAULDRON} missing, try 'make build' first" 2>&1; exit 1; }

TMPFILE=$(mktemp)

trap "rm ${TMPFILE}" EXIT

RED='\033[1;31m'
GREEN='\033[1;32m'
RESET='\033[0m'

EXIT_CODE=0

title() {
  echo -n "${1}: "
}

result() {
    if [[ "${?}" -eq 0 ]]; then
        printf "${GREEN}pass${RESET}\n"
    else
        printf "${RED}fail${RESET}\n"
        EXIT_CODE=1
    fi
}

title "Simple Rendering"
diff <(${CAULDRON} -template examples/hello.tmpl name=John time=morning) \
    examples/rendered/hello.rendered
result $?

title "JSON Rendering"
diff <(${CAULDRON} -template examples/resolv.conf.tmpl \
    nameservers='["10.20.30.40", "8.8.8.8"]' \
    domain=mydomain.com \
    options='{"rotate": "", "timeout": "5"}') \
    examples/rendered/resolv.conf.rendered
result $?

title "External JSON Datasource Rendering"
diff <(${CAULDRON} -template examples/treats.tmpl -json examples/treats.json) \
    examples/rendered/treats.rendered
result $?

title "Render to File"
${CAULDRON} -template examples/hello.tmpl -file "${TMPFILE}" name=John time=morning \
 && diff "${TMPFILE}" examples/rendered/hello.rendered
result $?

exit $EXIT_CODE
