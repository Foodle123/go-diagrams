package main

import (
    "log"
	"github.com/Foodle123/go-diagrams/nodes/azure"
    "github.com/Foodle123/go-diagrams/diagram"
)

func main() {
	d, err := diagram.New(diagram.Filename("ETL"), diagram.Label("ETL"), diagram.Direction("LR"))
	if err != nil {
		log.Fatal(err)
	}

//create FilesOrigin node
	filesOrigin := azure.Migration.MigrationProjects().Label("filesOrigin")

//create RawFile nodes
	genericFile1 := azure.General.CSVFile().Label("RawFile")
	genericFile2 := azure.General.TXTFile().Label("RawFile")
	genericFile3 := azure.General.XMLFile().Label("RawFile")

//create Load node	
	rawFileStorage := azure.Storage.GeneralStorage().Label("Load")

//connect RawFiles with Load
	d.Connect(filesOrigin, genericFile1)
	d.Connect(filesOrigin, genericFile2)
	d.Connect(filesOrigin, genericFile3)
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
	
	oracleServer := diagram.NewGroup("oracleServer").Label("Oracle Server").Add(genericFile1, genericFile2, genericFile3,rawFileStorage, staging, enrich, core, move, bais,)
	d.Group(oracleServer)

	bundesbank := azure.General.DeutscheBundesbankLogo().Label("Bundesbank")
	d.Connect(bais, bundesbank)
	
	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}
