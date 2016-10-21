FROM scratch
COPY ca.crt /opt/
COPY localhost.crt /opt/
COPY localhost.key /opt/
COPY server /opt/
EXPOSE 51000
ENTRYPOINT ["/opt/server"]
CMD ["/opt/server","/opt/localhost.key","/opt/localhost.crt","/opt/ca.crt"]
