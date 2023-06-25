FROM node:20-alpine
RUN mkdir /web
ADD . /web
WORKDIR /web
RUN npm install
CMD ["npm", "run", "dev", "--", "--host", "--port", "8081"]
EXPOSE 8081