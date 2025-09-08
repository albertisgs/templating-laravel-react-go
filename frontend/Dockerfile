# Base image Node.js
FROM node:18-alpine

# Set working directory
WORKDIR /app

# Copy package.json dan package-lock.json
COPY package*.json ./

# Install dependensi
RUN npm install

# Copy seluruh kode aplikasi
COPY . .

# Vite perlu expose port
EXPOSE 5173

# Jalankan dev server
CMD ["npm", "run", "dev", "--", "--host"]