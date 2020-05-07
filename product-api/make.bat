@ECHO off    
WHERE /q swagger
IF %ERRORLEVEL% NEQ 0 go get -u github.com/go-swagger/go-swagger/cmd/swagger 

swagger generate spec -o ./swagger.yaml --scan-models