FROM ruby:2.7-alpine

RUN apk add --update --no-cache \
      vim \
      build-base \
      postgresql-dev \
      tzdata \
      && rm -rf /var/cache/apk/*

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

WORKDIR /deploy

RUN gem install bundler:2.1.4

COPY Gemfile* ./
RUN bundle config --global frozen 1 \
  && bundle install -j4 --retry 3

COPY application.rb .
COPY Rakefile .

COPY config/database.sample.yml config/database.yml

COPY db/migrate db/migrate
