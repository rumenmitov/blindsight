FROM docker.io/library/golang

WORKDIR /app

RUN apt install -y git wget vim python3 cargo

RUN git clone https://github.com/rumenmitov/midas.git /midas

RUN wget -O /midas/weights https://github.com/isl-org/MiDaS/releases/download/v3_1/dpt_swin2_tiny_256.pt 

RUN mkdir /miniconda3

RUN wget https://repo.anaconda.com/miniconda/Miniconda3-latest-Linux-x86_64.sh -O ~/miniconda3/miniconda.sh

RUN bash /miniconda3/miniconda.sh -b -u -p /miniconda3

RUN rm -rf /miniconda3/miniconda.sh

RUN /miniconda3/bin/conda init bash

RUN cargo install depth_analyzer --locked

RUN cp /miniconda3/bin/* /usr/bin

RUN cp /home/root/.cargo/bin/depth_analyzer /usr/bin

WORKDIR /midas

RUN conda env create -f environment.yaml

RUN conda activate midas-py310

WORKDIR /app

COPY . .

RUN go test .

RUN go install .

ENTRYPOINT /go/bin/blindsight

EXPOSE 3000
