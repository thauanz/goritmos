#!bin/sh

for file in `ls -1 *.go` ; do \
    echo $file
    go run $file ;
done
