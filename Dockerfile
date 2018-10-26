FROM golang

ARG app_env
ENV app_env ${app_env}


COPY ./app /go/src/app
WORKDIR  /go/src/app

RUN go get ./
RUN go build -i

CMD if [ "${app_env}" = "prod" ]; then app; else go get github.com/pilu/fresh && fresh; fi

EXPOSE 8080