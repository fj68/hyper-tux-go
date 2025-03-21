FROM mcr.microsoft.com/devcontainers/go AS build

WORKDIR /app

# Install requirements
RUN apt update && apt install -y \
    xvfb\
    libasound2-dev libgl1-mesa-dev libxcursor-dev libxi-dev libxinerama-dev libxrandr-dev libxxf86vm-dev

COPY go.mod ./
RUN go mod download

# Copy source codes
COPY . .

# run xvfb in background
RUN Xvfb :99 -screen 0 1024x768x24 > /dev/null 2>&1 &
ENV DISPLAY=:99
