#--------------------#
# Setup server image #
#--------------------#

FROM ubuntu:latest

WORKDIR /app

COPY . .

RUN apt update && apt install -y wget git curl cargo

#--------------#
# Setup golang #
#--------------#

RUN wget https://dl.google.com/go/go1.16.5.linux-amd64.tar.gz && tar -xvf go1.16.5.linux-amd64.tar.gz && mv go /usr/local

ENV GOROOT=/usr/local/go

ENV GOPATH=$HOME/go

ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

#----------------------#
# Setup depth_analyzer #
#----------------------#

RUN cargo install depth_analyzer

#--------------#
# Start server #
#--------------#

RUN chmod +x runMiDaS.sh

RUN go test .

RUN go install .

ENTRYPOINT /go/bin/blindsight

EXPOSE 3000
