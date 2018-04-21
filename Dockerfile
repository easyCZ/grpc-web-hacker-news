FROM node

COPY . /code/
WORKDIR code/app
RUN ls
RUN npm install

ENTRYPOINT npm run start
