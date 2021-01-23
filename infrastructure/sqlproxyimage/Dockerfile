FROM gcr.io/cloudsql-docker/gce-proxy:1.19.1

COPY config/account.json /config/account.json

# Run the service when the container starts.
WORKDIR /
ENTRYPOINT /cloud_sql_proxy -instances=workoutapp-270814:europe-west2:mjpostgresworkout=tcp:0.0.0.0:3305 -credential_file=/config/account.json

# Document that the service listens on port 8080.
EXPOSE 3305
