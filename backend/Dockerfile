FROM golang
MAINTAINER dmoraschi

# Prepare app directory
RUN mkdir -p /usr/src/app/vendor
RUN chmod 777 /usr/src/app/vendor
WORKDIR /go/src/app

# Install Glide
RUN curl https://glide.sh/get | sh

# Expose the app port and app volume
EXPOSE 3000
VOLUME /go/src/app

RUN git config --global http.sslVerify true

# Start the app
CMD glide install && go run main.go