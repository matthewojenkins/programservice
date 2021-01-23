FROM golang:1.15

COPY build/program /go/bin/programservice/program
COPY src/config/database.properties /go/bin/programservice/database.properties


# Run the service when the container starts.
WORKDIR /go/bin/programservice
ENTRYPOINT /go/bin/programservice/program

# Document that the service listens on port 8080.
EXPOSE 8080
