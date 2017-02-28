FROM node
MAINTAINER dmoraschi

# Prepare app directory
RUN mkdir -p /usr/src/app

# Install dependencies
WORKDIR /usr/src/app

# Expose the app port and app volume
EXPOSE 3000
VOLUME /usr/src/app

# Start the app
CMD npm install && npm start