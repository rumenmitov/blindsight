#--------------------#
# Setup server image #
#--------------------#

FROM ubuntu:latest

WORKDIR /app

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

#-------------#
# Setup MiDaS #
#-------------#

RUN apt-get install -y python3 python3-pip

RUN pip install conda

RUN wget https://repo.anaconda.com/miniconda/Miniconda3-latest-Linux-x86_64.sh 

RUN bash Miniconda3-latest-Linux-x86_64.sh -b -p $HOME/miniconda 

RUN rm Miniconda3-latest-Linux-x86_64.sh

ENV PATH="/home/username/miniconda/bin:$PATH"

RUN git clone https://github.com/isl-org/MiDaS.git /MiDaS

RUN wget -O /MiDaS/weights/dpt_swin2_tiny_256.pt https://github.com/isl-org/MiDaS/releases/download/v3_1/dpt_swin2_tiny_256.pt

RUN conda env create -f environment.yaml

RUN conda activate midas-py310

#--------------#
# Start server #
#--------------#

COPY . .

RUN chmod +x runMiDaS.sh

RUN go test .

RUN go install .

ENTRYPOINT /go/bin/blindsight

EXPOSE 3000
