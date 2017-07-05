FROM  golang:1.8

MAINTAINER Enaho Murphy <enahomurphy@gmail.com>

RUN curl https://glide.sh/get | sh

EXPOSE 8085

COPY . /go/src/recipe

WORKDIR /go/src/recipe

CMD ["glide", "install"]
CMD ["go", "run", "main.go"]