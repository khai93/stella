FROM alpine:3.14

RUN apk update

# go
RUN apk add --no-cache go

# node
RUN apk add --update nodejs npm
RUN npm install -g jest

# python
ENV PYTHONUNBUFFERED=1
RUN apk add --update --no-cache python3 && ln -sf python3 /usr/bin/python
RUN python3 -m ensurepip
RUN pip3 install --no-cache --upgrade pip setuptools

# c++
RUN apk add build-base