#!/bin/bash
#

ANIMALS="Ayam Kelinci Jaguar"

for animal in $ANIMALS
do
    mkdir $animal
    for i in {1..10}
    do
        touch $animal/$animal-$i.txt
    done
done
