#!/bin/bash

SED_IFLAG=(-i'')
ERGO_SPEC_VERSION=master
GENERATOR_VERSION=v4.3.0
DIR=$(pwd)
PACKAGE_NAME=ergo_rest

GIT_USER_ID=ross-weir
GIT_REPO_ID=go-ergo-rest
  # -e GIT_USER_ID=${GIT_USER_ID} -e GIT_REPO_ID=${GIT_REPO_ID} \

docker run --user "$(id -u):$(id -g)" --rm -v ${DIR}:/local \
  openapitools/openapi-generator-cli:${GENERATOR_VERSION} generate \
  -i /local/codegen/openapi.yaml \
  -g go \
  -t /local/codegen/templates \
  --git-user-id ${GIT_USER_ID} \
  --git-repo-id ${GIT_REPO_ID} \
  --additional-properties packageName=${PACKAGE_NAME} \
  -o /local;

echo "fixing linting"

sed "${SED_IFLAG[@]}" 's/Api/API/g' *.go;
sed "${SED_IFLAG[@]}" 's/Json/JSON/g' *.go;
sed "${SED_IFLAG[@]}" 's/Id/ID/g' *.go;
sed "${SED_IFLAG[@]}" 's/Url/URL/g' *.go;

echo "fixing big int"

# import math/big into required files
BIGINT_FIX=("model_node_info" "model_pow_solutions")
for FILE in "${BIGINT_FIX[@]}" ; do
    sed "${SED_IFLAG[@]}" "s#package ${PACKAGE_NAME}#package ${PACKAGE_NAME}\n\nimport "'"math/big"'"\n#g" "${FILE}.go";
done

# replace types with big.Int
sed "${SED_IFLAG[@]}" 's/Difficulty \*int32/Difficulty *big.Int/g' model_node_info.go;
sed "${SED_IFLAG[@]}" 's/HeadersScore \*int32/HeadersScore *big.Int/g' model_node_info.go;
sed "${SED_IFLAG[@]}" 's/FullBlocksScore \*int32/FullBlocksScore *big.Int/g' model_node_info.go;
sed "${SED_IFLAG[@]}" 's/D float32/D big.Int/g' model_pow_solutions.go;

# fix unhandle-able types, tuples, etc
sed "${SED_IFLAG[@]}" 's/AnyOfstringinteger/interface{}/g' *.go;
sed "${SED_IFLAG[@]}" 's/AnyOfPaymentRequestBurnTokensRequestAssetIssueRequest/interface{}/g' *.go;
sed "${SED_IFLAG[@]}" 's/OneOfCommitmentWithSecretCommitmentSecretProven/interface{}/g' *.go;

# remove unneeded files
rm git_push.sh
rm .travis.yml

gofmt.exe -w *.go
