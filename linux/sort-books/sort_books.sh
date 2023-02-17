#!/bin/sh

BOOK_CATALOG=$1
TARGET_FOLDER=$2

# Apabila TARGET_FOLDER tidak ada, maka kita buat folder tersebut
if [ ! -d "$TARGET_FOLDER" ]; then
    mkdir "$TARGET_FOLDER"
fi

# Read setiap line dari file yang ingin di baca
while read -r line; do
    SOURCE_FILE="books/$line.pdf"
    TARGET_FILE="$TARGET_FOLDER/$line.pdf"
    if [ -e "$SOURCE_FILE" ]; then
        echo "Copy $SOURCE_FILE to $TARGET_FILE"
        cp "$SOURCE_FILE" "$TARGET_FILE"
    fi
done < "$BOOK_CATALOG"
