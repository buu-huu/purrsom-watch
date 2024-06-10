@echo off

set PROJECT_NAME=purrsom_watch
set GOOS=linux
set GOARCH=amd64

if "%GOOS%"=="windows" (
    set FILE_EXT=.exe
) else (
    set FILE_EXT=
)

go build  -C ..\cmd\watch -o ..\..\bin\%PROJECT_NAME%%FILE_EXT%

if %ERRORLEVEL% equ 0 (
    echo Build successful!
) else (
    echo Build failed!
)

pause
