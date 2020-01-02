FROM node:10

ARG NPM_TOKEN  
COPY .npmrc .npmrc  
COPY package.json package.json
COPY . .  
RUN npm install
RUN npm install autopublish
RUN ls -la
RUN ./node_modules/.bin/autopublish .
RUN rm -f .npmrc

CMD npm start