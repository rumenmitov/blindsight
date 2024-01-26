#--------------------#
# Setup server image #
#--------------------#

FROM arch:latest

WORKDIR /app

COPY . .

RUN pacman -Sy wget git curl cargo

#--------------#
# Setup golang #
#--------------#

RUN pacman -Sy go

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
