FROM ubuntu:jammy

WORKDIR /usr/local/src
RUN apt update
RUN apt install -y git ffmpeg make g++ vim wget
RUN git clone https://github.com/ggerganov/whisper.cpp.git -b v1.2.0 --depth 1

WORKDIR /usr/local/src/whisper.cpp
# whisper.cpp setup
RUN bash ./models/download-ggml-model.sh base.en
RUN make
RUN make samples
RUN ./main -f samples/gb0.wav
