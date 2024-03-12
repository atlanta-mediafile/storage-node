#!/bin/bash

# Check version file
if [ -e "version" ]; then
    # Get version from file
    current_version=$(head -n 1 "version")
    # Verifi version regex
    if ! echo "$current_version" | grep -E "^[0-9]+\.[0-9]+\.[0-9]+$" >/dev/null; then
        echo "Invalid version: $current_version"
        exit 1
    fi
else
    echo "The file version does not exits."
    exit 1
fi

# get commit name from args
commit="$1"

# get commit prefix
commit_prefix="${commit%%:*}"

# check commit prefix
if [ "$commit_prefix" == "feat" ]; then
    # get second version
    second_version=$(echo "$current_version" | cut -d '.' -f 2)

    # add version
    ((new_second_version = second_version + 1))

    # get new version
    # shellcheck disable=SC2001
    new_version=$(echo "$current_version" | sed "s/\.[0-9]\{1,\}\.[0-9]*$/.$new_second_version.0/")

    # insert new version in file version
    echo "$new_version" > "version"
else
    # get second version
    third_version=$(echo "$current_version" | cut -d '.' -f 3)

    # add version
    ((new_third_version = third_version + 1))

    # get new version
    # shellcheck disable=SC2001
    new_version=$(echo "$current_version" | sed "s/\.[0-9]*$/.$new_third_version/")

    # insert new version in file version
    echo "$new_version" > "version"
fi
echo "$new_version"
#echo "{new_version}={$new_version}" >> $GITHUB_OUTPUT