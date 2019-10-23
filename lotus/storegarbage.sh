#!/bin/bash
for ((i=0;i<14400;i++));do
{
sleep 1;lotus-storage-miner store-garbage
}
done
