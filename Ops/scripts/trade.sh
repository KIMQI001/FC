#!/usr/bin/env bash

go-filecoin client propose-storage-deal "${2}" "${3}" "${4}" 2880 --allow-duplicates --repodir="${1}"