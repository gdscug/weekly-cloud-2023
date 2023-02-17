#!/bin/sh
#
#
if [ -x $1 ]; then
    echo Masukkan category yang ingin dicari
    exit 1
fi

CATEGORY=$1
BOOKS=$(find books)

while read -r book; do
    echo "$book" | grep "-$CATEGORY"
done < "$BOOKS"
