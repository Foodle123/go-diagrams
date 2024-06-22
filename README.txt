#to generate the dot file
go run main.go

#to generate the png
dot -Tpng -O filename.dot

#in the nodes/assets folder
go generate

#powershell command to replace git path
Get-ChildItem -Path "C:\Users\cedri\Desktop\go-diagrams\go-diagrams" -Include *.go, *.mod -Recurse | ForEach-Object {(Get-Content $_.FullName) -replace 'foodle123', 'Foodle123' | Set-Content $_.FullName}

#workflow to add own assets
#1 add the png to the go-diagrams\assets\x\y\img.png
#2 edit the corrisponding file in go-diagrams\nodes\x\y.go
#  then add the function for the node
#3 run "go generate" in go-diagrams\nodes\assets
#4 push everything to the git



edit in the dot file

pad=0.5;
ranksep=2;