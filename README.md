# firstgowebservice

had to enable a whole load of service apis

gcloud services enable cloudresourcemanager.googleapis.com \
    compute.googleapis.com \
    iam.googleapis.com \
    oslogin.googleapis.com \
    servicenetworking.googleapis.com \
    sqladmin.googleapis.com


# Run docker locally
# Create docker network for containers
docker network create servicenet
# run the cloud sql proxy container
docker run -d --name=sqlproxy --network=servicenet -v C:\Users\matt\IdeaProjects\programservice\infrastructure:/config -p 127.0.0.1:3305:3305 gcr.io/cloudsql-docker/gce-proxy:1.19.1 /cloud_sql_proxy -instances=workoutapp-270814:europe-west2:mjpostgresworkout=tcp:0.0.0.0:3305 -credential_file=/config/account.json

# run service which uses local docker
docker run -p 8080:8080 --name programservice --network=servicenet programservice:1.0

# run in K8s (after configuring minikube docker image repo location
@FOR /f "tokens=*" %i IN ('minikube -p minikube docker-env') DO @%i

kubectl create deployment programservice --image=programservice:1.0
kubectl expose deployment programservice --type=LoadBalancer --port=8090