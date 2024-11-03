#!/bin/sh

for i in {15..23}; do
  for j in {1..25}; do
    kotlin /Users/adamjonas/.config/.aoc/jars/current.jar -y $i -d $j -c
    # sleep 1
  done
done
