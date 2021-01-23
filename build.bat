
rem Go Build
set GOOS=linux
set GOARCH=amd64

cd src
go build -o ..\build\program .

rem Docker Build
cd ..
docker build -f Dockerfile -t programservice:1.0 .

