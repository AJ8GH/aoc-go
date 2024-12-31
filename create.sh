#!/bin/sh

# for i in {15..23}; do
#   for j in {1..25}; do
#     kotlin /Users/adamjonas/.config/.aoc/jars/current.jar -y $i -d $j -c
#     # sleep 1
#   done
# done

for i in {16..25}; do
  kotlin /Users/adamjonas/.config/.aoc/jars/current.jar -y 24 -d $i -c
  # sleep 1
done
