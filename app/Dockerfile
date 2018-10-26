FROM golang:1.9-alpine
LABEL maintainer="Javier lopez Lopez"

COPY . /go/src/app
COPY config/entrypoint.sh /entrypoint.sh
COPY config/realize.dev.yml /go/src/app/.realize.yaml

RUN apk add --no-cache ca-certificates \
        dpkg \
        gcc \
        git \
        musl-dev \
        bash

ENV GOPATH /go
ENV PATH =  ${GOPATH}/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "${GOPATH}/src" "${GOPATH}/bin" \
    && chmod -R 777 "${GOPATH}" \
    && chmod +x /entrypoint.sh

RUN go get github.com/tockins/realize github.com/gorilla/mux

WORKDIR ${GOPATH}
ENTRYPOINT [ "/entrypoint.sh" ]

CMD [ "realize", "start"]
