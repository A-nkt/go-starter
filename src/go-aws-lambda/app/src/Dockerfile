FROM python:3.9.8-bullseye

ENV OS linux
ENV ARCH amd64
ENV GO_VERSION 1.21.2

ENV GOPATH /go
ENV GO111MODULE on
ENV GOPROXY https://proxy.golang.org,direct
ENV GOPRIVATE github.com/weathernews
ENV CGO_ENABLED 0
ENV GOOS $OS
ENV GOARCH $ARCH
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN set -ex \
        && wget https://golang.org/dl/go${GO_VERSION}.${OS}-${ARCH}.tar.gz \
        && tar -C /usr/local -xzf go${GO_VERSION}.${OS}-${ARCH}.tar.gz \
        && rm go${GO_VERSION}.${OS}-${ARCH}.tar.gz && mkdir -p "${GOPATH}/src" "${GOPATH}/bin" && chmod -R 777 "${GOPATH}"

RUN set -ex \
        && pip install --no-cache-dir --upgrade pip && pip install --no-cache-dir \
                aws-sam-cli

COPY ./id_rsa root/.ssh/id_rsa

RUN set -ex \
        && git config --global url."git@github.com:".insteadOf "https://github.com/" \
        && ssh -o StrictHostKeyChecking=no -T git@github.com || true
ENTRYPOINT [ "sam" ]
CMD [ "--version" ]