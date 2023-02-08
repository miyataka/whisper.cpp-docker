#!/bin/bash
set -ex

INPUT_MP4_FILE_PATH=$1
INPUT_MP4_FILE_NAME=`basename $INPUT_MP4_FILE_PATH`
LANG=$2

if [ x$LANG == x ]; then
    LANG=en
fi

IMAGE=`docker build -q .`

# cp $INPUT_MP4_FILE_PATH .

docker run -v $PWD:/usr/local/src/whisper.cpp/mnt -it $IMAGE entrypoint.sh ./mnt/$INPUT_MP4_FILE_NAME $LANG
