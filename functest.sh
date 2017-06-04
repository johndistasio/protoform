#!/bin/bash
set -o errexit

CAULDRON=build/cauldron

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