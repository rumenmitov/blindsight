#!/bin/bash

wget -O /MiDaS/weights/dpt_swin2_tiny_256.pt https://github.com/isl-org/MiDaS/releases/download/v3_1/dpt_swin2_tiny_256.pt

conda env create -f environment.yaml && conda init 
