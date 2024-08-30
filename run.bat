@echo off

REM Change to the directory containing this script
cd /d "%~dp0"

REM Compile the Go program
go build -o %TEMP%\my-http-server.exe app\server.go

REM Run the compiled program
%TEMP%\my-http-server.exe %*