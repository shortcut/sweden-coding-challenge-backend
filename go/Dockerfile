FROM golang:1.19.1-alpine3.16

RUN sed -i 's/https/http/' /etc/apk/repositories

ARG USER_NAME
ARG GROUP_NAME
ARG USER_ID
ARG GROUP_ID
ARG DOMAIN_DIR

WORKDIR $DOMAIN_DIR

RUN echo "Asia/Tehran" > /etc/timezone

RUN addgroup -g ${GROUP_ID} ${GROUP_NAME} \
    && adduser -G ${GROUP_NAME} -g ${GROUP_NAME} -s /bin/sh -D ${GROUP_NAME}

RUN mkdir -p $DOMAIN_DIR

RUN chown ${USER_NAME}:${GROUP_NAME} ${DOMAIN_DIR}

RUN apk add --no-cache git curl wget