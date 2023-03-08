#!/bin/bash

FOLDER="certificates"
NAMES_FILE="names.txt"

if [[ ! -d $FOLDER ]]; then
    echo "Creating certificates folder..."
    mkdir $FOLDER
fi

while read -r name; do
    echo Creating "$name".html
    sed -e "s/NAMA_DISINI/$name/" cert_template.html > $FOLDER/"$name".html
done < "$NAMES_FILE"
