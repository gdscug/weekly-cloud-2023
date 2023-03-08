#!/bin/bash

# Ambil argumen dari command
CATALOG=$1
TARGET_FOLDER=$2
SOURCE_FOLDER="books"

# Jika argumen pertama atau catalog text tidak diinput
if [[ -z $1 ]]; then
    echo "Please input catalog text"
    exit 1
fi
#
# Jika argumen kedua atau target folder tidak diinput
if [[ -z $2 ]]; then
    echo "Please input target folder"
    exit 1
fi

# Jika target folder belum dibuat, maka dibuat
if [[ ! -d $TARGET_FOLDER ]];then
    mkdir "$TARGET_FOLDER"
fi

# Loop semua book dalam catalog
while read -r line; do
    # File dari folder books
    SOURCE_FILE="$SOURCE_FOLDER/$line.pdf"
    
    # File tujuan untuk dicopy
    TARGET_FILE="$TARGET_FOLDER/$line.pdf"

    # Copy dengan command `cp`
    cp "$SOURCE_FILE" "$TARGET_FILE"
    echo Copying from "$SOURCE_FILE" to "$TARGET_FILE"
done < "$CATALOG"
