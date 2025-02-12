
ARG  ID_AWS_ACCOUNT_ECR
FROM ${ID_AWS_ACCOUNT_ECR}.dkr.ecr.us-east-1.amazonaws.com/blt-msa-ha-checkout-docker-base:golang-1.21.4-alpine as builder
RUN apk update && apk upgrade && apk --no-cache add libc-dev glib-static gcc build-base ca-certificates gcc musl-dev git go 
RUN git clone --depth 1 --branch v2.4.0 https://github.com/edenhill/librdkafka.git \
  && cd librdkafka \
  && ./configure --prefix=/usr/local --libdir=/usr/local/lib --disable-lz4 --disable-ssl --disable-sasl --disable-zstd \
  && make \
  && make install

ARG APPNAME
ARG RESTPORT
ARG REPO_TOKEN
RUN mkdir -p {/app,/app/build,/app/scripts}
WORKDIR /app
ADD . ./
RUN git settings --global url.https://$REPO_TOKEN@github.com/jhonquirama/.insteadOf https://github.com/jhonquirama
ENV GOPRIVATE=github.com/jhonquirama/*
RUN make aws_build

FROM ${ID_AWS_ACCOUNT_ECR}.dkr.ecr.us-east-1.amazonaws.com/blt-msa-alpine-bsae-doker:latest


RUN apk add --no-cache \
         python3 \
         py3-pip \
         jq \
         tzdata \
     && pip3 install --upgrade pip \
     && pip3 install --no-cache-dir \
         awscli \
         yq \
     && rm -rf /var/cache/apk/*

ENV TZ=America/Bogota
RUN aws --version 

RUN mkdir doc
COPY --from=builder /app/build/ .
COPY --from=builder /app/entrypoint.sh .
COPY --from=builder /app/doc/ ./doc

RUN ["chmod", "+x", "entrypoint.sh"]

CMD [ "/entrypoint.sh" ]
#CMD ["/bin/sh"]
