#--------------------#
# Setup server image #
#--------------------#

FROM docker.io/archlinux

WORKDIR /app

COPY . .

RUN pacman -Syu

RUN pacman -S wget git curl cargo go

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
