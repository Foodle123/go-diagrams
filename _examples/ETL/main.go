package main

import (
    "log"
	"github.com/Foodle123/go-diagrams/nodes/azure"
	//"../../nodes/azure"
    "github.com/Foodle123/go-diagrams/diagram"
    //"github.com/blushft/go-diagrams/nodes/aws"
    //"github.com/blushft/go-diagrams/nodes/generic"
)

func main() {
	d, err := diagram.New(diagram.Filename("ETL"), diagram.Label("ETL"), diagram.Direction("LR"))
	if err != nil {
		log.Fatal(err)
	}


//create RawFile nodes
	genericFile1 := azure.General.CSVFile().Label("RawFile")
	genericFile2 := azure.General.TXTFile().Label("RawFile")
	genericFile3 := azure.General.XMLFile().Label("RawFile")

//create Load node	
	rawFileStorage := azure.Storage.GeneralStorage().Label("Load")

//connect RawFiles with Load
	d.Connect(genericFile1, rawFileStorage)
	d.Connect(genericFile2, rawFileStorage)
	d.Connect(genericFile3, rawFileStorage)
	
	staging := azure.Database.SqlDatabases().Label("Staging")
	d.Connect(rawFileStorage, staging)
	enrich := azure.Compute.CloudServicesClassic().Label("Enrich")
	d.Connect(staging, enrich)
	core := azure.Database.SqlDatabases().Label("Core")
	d.Connect(enrich, core)
	move := azure.Migration.DatabaseMigrationServices().Label("Move")
	d.Connect(core, move)
	bais := azure.Migration.MigrationProjects().Label("BAIS")
	d.Connect(move, bais)
	
	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}
