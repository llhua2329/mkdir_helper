@echo off

setlocal

if exist install.bat goto ok
echo install.bat must bi run from its folder
goto end

:ok

set OLDGOPATH=%GOPATH%
set GOPATH=%~dp0

gofmt -w src

go install mkdir_helper

:end
echo finished