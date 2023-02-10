# whisper.cpp-docker

this repository is just dockerinzed one that [ggerganov/whisper.cpp](https://github.com/ggerganov/whisper.cpp).

## how to use
- build on your machine

### build on your machine
#### requirements
- docker
- git

#### howto
- clone this repo
    - `git clone git@github.com:miyataka/whisper.cpp-docker.git`
    - `cd whisper.cpp-docker`
- docker build
    - `docker build .`
- download mp4 file if you need
- run copy_and_transcription.sh with args
    - `./copy-and-transcription.sh /path/to/your_video_file.mp4 ja`

## limitation
- only support mp4 file

## more detail
- go to [ggerganov/whisper.cpp](https://github.com/ggerganov/whisper.cpp) and
  read its document.
