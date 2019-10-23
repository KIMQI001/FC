#!/usr/bin/env bash
go-filecoin daemon \
  --cmdapiaddr=/ip4/127.0.0.1/tcp/"${2}" \
  --swarmlisten=/ip4/127.0.0.1/tcp/"${3}" \
  --repodir="${1}" &