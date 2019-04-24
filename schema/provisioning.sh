#!/bin/sh -eux
cat $1 | sqlite3 $2
