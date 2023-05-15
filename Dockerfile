FROM ubuntu:jammy

WORKDIR /usr/local/src
RUN apt update
RUN apt install -y git ffmpeg make g++ vim wget
RUN git clone https://github.com/ggerganov/whisper.cpp.git -b v1.4.0 --depth 1

# whisper.cpp setup
WORKDIR /usr/local/src/whisper.cpp
RUN bash ./models/download-ggml-model.sh medium
RUN make

ADD entrypoint.sh /usr/local/bin
