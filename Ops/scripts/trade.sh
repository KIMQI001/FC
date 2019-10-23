#!/usr/bin/env bash

for ((i=0;i<14400;i++));do
{
go-filecoin client propose-storage-deal "${2}" "${3}" "${4}" 2880 --allow-duplicates --repodir="${1}";sleep 30
}
done