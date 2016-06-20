FROM centurylink/ca-certs

EXPOSE 8080

COPY polyglot-co-demo-backend /
ADD templates /templates
ADD static /static

WORKDIR /

ENTRYPOINT ["/polyglot-co-demo-backend"]
