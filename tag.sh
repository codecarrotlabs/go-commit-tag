#! /bin/bash

####
# Create new tag with message
#
# @params string tag_name
# sh ./tag.sh v10.423.023
#####

NEW_TAG=$1

RECENT_TAG=$(git describe --abbrev=0)

git tag -a $NEW_TAG -m "https://github.com/codecarrotlabs/go-commit-tag/compare/$RECENT_TAG...$NEW_TAG"
