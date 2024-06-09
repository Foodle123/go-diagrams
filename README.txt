#to generate the dot file
go run main.go

#to generate the png
dot -Tpng -O filename.dot

#in the nodes/assets folder
go generate

#powershell command to replace git path
Get-ChildItem -Path "C:\Users\cedri\Desktop\go-diagrams\go-diagrams" -Include *.go, *.mod -Recurse | ForEach-Object {(Get-Content $_.FullName) -replace 'foodle123', 'Foodle123' | Set-Content $_.FullName}
