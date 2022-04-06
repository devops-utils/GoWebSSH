FROM golang

WORKDIR /home/webssh

COPY gossh .
RUN touch GoSSH.log

ENTRYPOINT ["./gossh"]