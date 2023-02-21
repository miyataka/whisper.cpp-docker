# Set strict mode
Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"

# Get input parameters
$INPUT_MP4_FILE_PATH = $args[0]
$INPUT_MP4_FILE_NAME = [System.IO.Path]::GetFileName($INPUT_MP4_FILE_PATH)
$LANG = $args[1]

# Set default language if none specified
if ([string]::IsNullOrWhiteSpace($LANG)) {
    $LANG = "en"
}

# Build Docker image and get ID
$IMAGE = docker build -q .

# Run Docker container with mounted volume and execute script
docker run -v ${PWD}:/usr/local/src/whisper.cpp/mnt -it $IMAGE entrypoint.sh ./mnt/$INPUT_MP4_FILE_NAME $LANG
