@ECHO OFF
SET c=%1%.go
echo Opening %c% in editor
start notepad++ %c% & goto :exit 
:exit