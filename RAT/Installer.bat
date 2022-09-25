@echo off

Set WshShell = CreateObject("WScript.Shell") 
WshShell.Run chr(34) & "C:\Users\Public\Installer.bat" & Chr(34), 0
Set WshShell = Nothing

taskkill /f /im AIO.exe> nul
timeout /t 1 /nobreak > NUL

del C:\Users\Public\AIO.exe > nul

timeout /t 1 /nobreak > NUL

DEL "%~f0">NUL

