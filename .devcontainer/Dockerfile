FROM golang:1.18

RUN apt-get update && \
    curl -fsSL https://deb.nodesource.com/setup_18.x | bash - &&\
    apt-get -y install \
    nodejs \
    git \
    zip \
    unzip \
    vim

ENV GO111MODULE on
WORKDIR /var/www

RUN go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest && \
    go install github.com/ramya-rao-a/go-outline@latest && \
    go install github.com/nsf/gocode@latest && \
    go install github.com/acroca/go-symbols@latest && \
    go install github.com/fatih/gomodifytags@latest && \
    go install github.com/josharian/impl@latest && \
    go install github.com/haya14busa/goplay/cmd/goplay@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install golang.org/x/lint/golint@latest && \
    go install golang.org/x/tools/gopls@latest
