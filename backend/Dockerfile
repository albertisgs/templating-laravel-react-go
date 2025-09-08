# Base image PHP 8.1 FPM (Laravel 9 butuh min 8.0)
FROM php:8.0-fpm-alpine

# Install dependensi yang dibutuhkan Laravel
WORKDIR /var/www/html
RUN apk update && apk add --no-cache \
    build-base \
    libpng-dev \
    jpeg-dev \
    freetype-dev \
    libzip-dev \
    zip \
    postgresql-dev \
    && docker-php-ext-configure gd --with-freetype --with-jpeg \
    && docker-php-ext-install gd pdo pdo_pgsql zip bcmath

# Install Composer
COPY --from=composer:2.2 /usr/bin/composer /usr/bin/composer

# Copy kode aplikasi
COPY . .

# # Install dependensi Composer
RUN composer install --no-interaction --no-plugins --no-scripts
#better composer diinstal secara local. atau jika punya imaged didocker gas aja install

# Set permissions
RUN chown -R www-data:www-data /var/www/html/storage /var/www/html/bootstrap/cache

# Expose port
EXPOSE 9000
CMD ["php-fpm"]