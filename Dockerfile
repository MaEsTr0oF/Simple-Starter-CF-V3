FROM node:20
WORKDIR /src
COPY package*.json ./
RUN npm install
COPY . ./
RUN npm run build
CMD ["sh", "-c", "echo 'Build completed. Copy files from ./dist'"]