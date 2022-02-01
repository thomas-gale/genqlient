FROM ubuntu:18.04 
LABEL maintainer="Thomas Gale (gale.thomasjames@gmail.com)"

# Configure a SUDO non-root user (no password)
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && apt-get update \
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME
USER $USERNAME

# Versions
ENV GO_VERSION=1.17

# Go 
ENV GO_DIST=go$GO_VERSION.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

# Update, install base packages and tools
RUN sudo apt-get update \
    && sudo apt-get -y install git gcc zip unzip apt-transport-https \
    ca-certificates golang-cfssl \
    traceroute net-tools curl dnsutils \
    gnupg-agent \
    software-properties-common \
    #
    # Install go
    && sudo curl -OJ https://dl.google.com/go/$GO_DIST \
    && sudo tar -C /usr/local -xzf $GO_DIST \
    && sudo rm -rf $GO_DIST
