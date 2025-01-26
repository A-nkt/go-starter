#! /bin/bash

sam deploy \
    --stack-name go-aws-lambda \
    --capabilities CAPABILITY_NAMED_IAM \
    --region ap-northeast-1 \
    --profile training \
    --resolve-s3
