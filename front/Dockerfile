FROM node:14.16-alpine3.12 as build
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
COPY package.json /app
RUN npm install
RUN npm install @angular/cli@11.2.12
COPY . /app
EXPOSE 4200
CMD ng serve --host 0.0.0.0