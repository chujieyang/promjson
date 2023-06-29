from amd64/alpine:3.16
WORKDIR /yusur
COPY main ./
EXPOSE 19101
ENTRYPOINT ["./main"]