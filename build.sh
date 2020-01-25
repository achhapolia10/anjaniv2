echo "Creating binaries for Inverntory-Manager"

go build -o "imanager64" .
if [ $? -eq 0 ]; then
  echo "Executable Generated Successfully"
  else
    echo "Executable Genration Failed"
fi

env GOOS=windows GOARCH=amd64 go build -o "imanager_64.exe" .
if [ $? -eq 0 ]; then
  echo "Executable Generated Successfully"
  else
    echo "Executableration Failed"
fi

env GOOS=windows GOARCH=amd64 go build -o "imanager_32.exe" .
if [ $? -eq 0 ]; then
  echo "Executable Generated Successfully"
  else
    echo "Executable Genration Failed"
fi

