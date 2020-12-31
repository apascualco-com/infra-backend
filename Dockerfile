FROM arm64v8/golang:1.14
COPY . /opt/app
WORKDIR /opt/app/
RUN go build -o /bin/app ./cmd/infra && chmod +x /bin/app && rm -Rf /opt/app && mkdir /etc/backend
EXPOSE 80
CMD ["/bin/app"]
