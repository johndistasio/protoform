#!/bin/bash
#
# A functional test script for cauldron.
#
# @author John DiStasio <jndistasio@gmail.com>
# @copyright Copyright (c) 2017 John DiStasio
#
set -o errexit

CAULDRON=build/cauldron

[[ -e ${CAULDRON} ]] || { echo "${CAULDRON} missing" 2>&1; exit 1; }

diff <(${CAULDRON} -template examples/hello.tmpl name=John time=morning) \
    examples/rendered/hello.rendered \
    && echo "pass"

diff <(${CAULDRON} -template examples/resolv.conf.tmpl \
    nameservers='["10.20.30.40", "8.8.8.8"]' \
    domain=mydomain.com \
    options='{"rotate": "", "timeout": "5"}') \
    examples/rendered/resolv.conf.rendered \
    && echo "pass"

diff <(${CAULDRON} -template examples/treats.tmpl -json examples/treats.json) \
    examples/rendered/treats.rendered \
    && echo "pass"
