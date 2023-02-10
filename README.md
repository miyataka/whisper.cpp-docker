# whisper.cpp-docker

this repository is just dockerinzed one that [ggerganov/whisper.cpp](https://github.com/ggerganov/whisper.cpp).

## requirements
- docker
- git

## how to use
1. clone this repo
    - `git clone git@github.com:miyataka/whisper.cpp-docker.git`
    - `cd whisper.cpp-docker`
2. docker build or pull pre-built image
    - `docker build .`
    - or
    - `docker pull miyataka/whisper.cpp-docker:latest`
3. place your audio file to same directory
    - `cp /path/to/your_video_file.mp4 ./`
4. run copy_and_transcription.sh with args
    - `./copy-and-transcription.sh <your_video_file_name> <language_keyword>`
        - language_keyword list is [here](https://github.com/ggerganov/whisper.cpp/blob/cfc06bf8dfea876dded28cff647bd98267dfc718/whisper.cpp#L119)
        - language_keyword is optional.
    - So
    - `./copy-and-transcription.sh ./your_video_file.mp4`
    - or
    - `./copy-and-transcription.sh ./your_video_file.mp4 ja`

## limitation
- only test in mp4, mp3 file

## more detail
- go to [ggerganov/whisper.cpp](https://github.com/ggerganov/whisper.cpp) and
  read its document.
