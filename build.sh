echo "Creating binaries for Inverntory-Manager"

go build -o "imanager64" .
if [ $? -eq 0 ]; then
  echo "Executable Generated Successfully"
  else
    echo "Executable Genration Failed"
fi

env GOOS=windows CGO_ENABLED=1 GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -o "imanager_64.exe" .
if [ $? -eq 0 ]; then
  echo "Executable Generated Successfully"
  else
    echo "Executableration Failed"
fi

env GOOS=windows CGO_ENABLED=1 GOARCH=386 CC=i686-w64-mingw32-gcc go build -o "imanager_32.exe" .
if [ $? -eq 0 ]; then
  echo "Executable Generated Successfully"
  else
    echo "Executable Genration Failed"
fi

