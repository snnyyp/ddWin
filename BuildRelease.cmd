: Goto the main function directory
cd main

: Build for Windows 32 Bit
go env -w GOOS=windows
go env -w GOARCH=386
go build -ldflags "-s -w" -o %GOPROJECT%\bin\ddWin\ddWin_0.1_windows_386.exe

: Build for Windows 64 Bit
go env -w GOOS=windows
go env -w GOARCH=amd64
go build -ldflags "-s -w" -o %GOPROJECT%\bin\ddWin\ddWin_0.1_windows_amd64.exe

: Go back to the previous directory
cd ..
