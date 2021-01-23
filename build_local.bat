
rem Go Build
set GOOS=windows
set GOARCH=amd64

cd src
go build .

move programservice.exe ..\build
copy config\local_database.properties ..\build\database.properties

cd ..\build
programservice
