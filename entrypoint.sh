#!/bin/bash

INPUT_MP4_FILE_PATH=$1
LANG=$2

ffmpeg -i $INPUT_MP4_FILE_PATH -ar 16000 samples/test.wav

./main -m models/ggml-medium.bin --output-csv --output-file ./mnt/output -l $LANG -f samples/test.wav
