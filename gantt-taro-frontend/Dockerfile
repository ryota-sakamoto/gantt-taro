FROM node:12.6.0
RUN apt-get update -y && \
    apt-get install git && \
    npm i vue-cli eslint -g
ENV NUXT_HOST 0.0.0.0
WORKDIR /src

CMD ["npm", "run", "dev"]