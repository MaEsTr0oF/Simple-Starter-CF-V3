FROM node:20

# Устанавливаем рабочую директорию
WORKDIR /src

COPY package*.json ./

# Устанавливаем зависимости
RUN npm install --legacy-peer-deps

COPY . ./

RUN npm run build

CMD ["sh", "-c", "echo 'Build completed. Copy files from ./dist'"]
