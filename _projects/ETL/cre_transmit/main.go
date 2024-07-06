package main

import (
	"log"

	"github.com/Foodle123/go-diagrams/diagram"
	"github.com/Foodle123/go-diagrams/nodes/azure"
)

func main() {
	d, err := diagram.New(diagram.Filename("Creditreform"), diagram.Label("crefo"), diagram.Direction("LR"))
	if err != nil {
		log.Fatal(err)
	}

	//create crefo_nodes

	REQ_CREFO_PREP := azure.Database.SqlDatabases().Label("REQ_CREFO_PREP")

	spool_Crefo := azure.Database.SqlServers().Label("spool_crefo_data")

	crefo_out_de := azure.General.CSVFile().Label("Crefo_out_DE")
	regulatory_reporting_team := azure.General.Usericon().Label("Regulatory Reporting Team")
	crefo := azure.General.Crefo().Label("Creditreform")

	//connect nodes
	d.Connect(REQ_CREFO_PREP, spool_Crefo)
	d.Connect(spool_Crefo, crefo_out_de)
	d.Connect(crefo_out_de, regulatory_reporting_team, diagram.WithLabel("Mail"))
	d.Connect(regulatory_reporting_team, crefo)

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}
