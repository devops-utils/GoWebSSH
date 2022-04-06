FROM golang

WORKDIR /home/webssh

COPY gossh .
RUN touch GoSSH.log
RUN mkdir /home/webssh/log

ENTRYPOINT ["./gossh"]