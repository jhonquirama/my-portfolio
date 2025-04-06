
# syntax=docker/dockerfile:1

# image where make compiled goland app
ARG  ID_AWS_ACCOUNT_ECR
FROM ${ID_AWS_ACCOUNT_ECR}.dkr.ecr.us-east-1.amazonaws.com/blt-msa-ha-checkout-docker-base:golang-1.22.1-alpine as builder
RUN apk update && apk upgrade && apk --no-cache add libc-dev glib-static gcc build-base ca-certificates gcc musl-dev git go
RUN git clone --depth 1 --branch v2.4.0 https://github.com/edenhill/librdkafka.git \
  && cd librdkafka \
  && ./configure --prefix=/usr/local --libdir=/usr/local/lib --disable-lz4 --disable-ssl --disable-sasl --disable-zstd \
  && make \
  && make install



ARG APPNAME
ARG RESTPORT
RUN mkdir -p {/app,/app/build,/app/scripts}
WORKDIR /app
ADD . ./
CMD go get .
RUN make aws_build


# image more small for run de golang app


FROM ${ID_AWS_ACCOUNT_ECR}.dkr.ecr.us-east-1.amazonaws.com/blt-msa-ha-checkout-docker-base:linux-alpine

RUN apk add --no-cache \
         python3 \
         py3-pip \
     && pip3 install --upgrade pip \
     && pip3 install --no-cache-dir \
         awscli \
     && rm -rf /var/cache/apk/*


RUN aws --version

RUN mkdir docs
COPY --from=builder /app/build/ .
COPY --from=builder /app/entrypoint.sh .
COPY --from=builder /app/docs/ ./docs

RUN ["chmod", "+x", "entrypoint.sh"]
# entrypoint configure container an run dinamic the app golang

CMD [ "/entrypoint.sh" ]
#CMD ["/bin/sh"]