
rem Go Build
set GOOS=linux
set GOARCH=amd64

rem Docker Build
docker build -f Dockerfile -t cloud-sql-proxy:1.1 .
