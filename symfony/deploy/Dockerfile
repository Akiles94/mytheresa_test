FROM php:8.0-apache

RUN a2enmod rewrite
WORKDIR /var/www/html

ENV TZ=America/Guayaquil \
    DEBIAN_FRONTEND=noninteractive

COPY . .

RUN apt-get update && apt-get install -y git unzip zip

RUN echo 'deb [trusted=yes] https://repo.symfony.com/apt/ /' | tee /etc/apt/sources.list.d/symfony-cli.list && \
    apt-get update && \
    apt-get install -y symfony-cli

COPY --from=composer:2 /usr/bin/composer /usr/bin/composer

RUN composer install

EXPOSE 8000