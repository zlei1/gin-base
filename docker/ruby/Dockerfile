FROM ruby:2.7-alpine

RUN apk add --update --no-cache \
      vim \
      build-base \
      postgresql-dev \
      tzdata \
      && rm -rf /var/cache/apk/*

WORKDIR /deploy

RUN gem install bundler:2.1.4

COPY Gemfile* ./
RUN bundle config --global frozen 1 \
  && bundle install -j4 --retry 3

COPY application.rb .
COPY Rakefile .

RUN mkdir config
COPY config/database.yml ./config
