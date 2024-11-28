FROM node:20

WORKDIR /src

COPY package*.json ./
RUN npm install

RUN npm install -g gulp-cli

COPY . ./

RUN npm run build

CMD ["sh", "-c", "echo 'Build completed. Copy files from ./dist'"]
