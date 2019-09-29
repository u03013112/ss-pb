#!/bin/bash
git pull &&
if [ ! -f ./version ]; then
    echo "1.0.0" > ./version
    fi
VERSION=$(cat ./version) && \
echo "Current version: $VERSION" && \
NEXT_VERSION=$(echo $VERSION | awk -F. -v OFS=. 'NF==1{print ++$NF}; NF>1{$NF=sprintf("%0*d", length($NF), ($NF+1)); print}') && \
echo $NEXT_VERSION > ./version && \
git add ./version && \
git commit -m "Released version v$NEXT_VERSION" && \
git tag -a v$NEXT_VERSION -m "Released version v$NEXT_VERSION" && \
git push origin v$NEXT_VERSION
git push
