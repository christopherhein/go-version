#!/bin/bash

PREV_VERSION=$(git describe --tags `git rev-list --tags --max-count=1`)
NEW_HASH=$(git rev-parse --verify HEAD)
NEXT_TAG=$(cat VERSION)

OUTPUT=$(git log ${PREV_VERSION}..${NEW_HASH} --no-merges --pretty=format:'* [view commit](http://github.com/christopherhein/go-version/commit/%H) %s' --reverse)

echo "<!-- START ${NEXT_TAG} -->" >> "CHANGELOG.md" 
echo "## ${NEXT_TAG}" >> "CHANGELOG.md"
echo "" >> "CHANGELOG.md"
echo "${OUTPUT}" >> "CHANGELOG.md"
echo "<!-- END ${NEXT_TAG} -->" >> "CHANGELOG.md"
echo "" >> "CHANGELOG.md"