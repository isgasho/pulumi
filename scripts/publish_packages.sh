#!/bin/bash
# publish_packages.sh uploads our packages to package repositories like npm
set -o nounset -o errexit -o pipefail
ROOT=$(dirname $0)/..

if [[ "${TRAVIS_OS_NAME:-}" == "linux" ]]; then
    echo "Publishing NPM package to NPMjs.com:"
    NPM_TAG="dev"

    # If the package doesn't have a pre-release tag, use the tag of latest instead of
    # dev. NPM uses this tag as the default version to add, so we want it to mean
    # the newest released version.
    if [[ $(jq -r .version < "${ROOT}/sdk/nodejs/bin/package.json") != *-* ]]; then
        NPM_TAG="latest"
    fi

    pushd ${ROOT}/sdk/nodejs/bin && \
        npm publish --tag "${NPM_TAG}" && \
        npm info 2>/dev/null || true && \
        popd

    echo "Publishing Pip package to pulumi.com:"
    twine upload \
        --repository-url https://pypi.pulumi.com?token=${PULUMI_API_TOKEN} \
        -u pulumi -p pulumi \
        ${ROOT}/sdk/python/env/src/dist/*.whl
fi

${PULUMI_SDK}/scripts/build-sdk.sh $(${ROOT}/scripts/get-version) $(git rev-parse HEAD)
