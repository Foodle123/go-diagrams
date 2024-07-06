package main

import (
	"log"

	"github.com/Foodle123/go-diagrams/diagram"
	"github.com/Foodle123/go-diagrams/nodes/azure"
)

func main() {
	d, err := diagram.New(diagram.Filename("010_crefo_enrich_etl"), diagram.Label("crefo"), diagram.Direction("TB"))
	if err != nil {
		log.Fatal(err)
	}

	//create crefo_input_nodes

	EXP_KNEIFF_GSWS := azure.Database.SqlDatabases().Label("EXP_KNEIFF_GSWS")
	EXP_GSTSLF_GSWS := azure.Database.SqlDatabases().Label("EXP_GSTSLF_GSWS")
	EXP_KNEIFF_RTL_DE_SRS := azure.Database.SqlDatabases().Label("EXP_KNEIFF RTL_DE_SRS")
	EXP_GSTSLF_RTL_DE_SRS := azure.Database.SqlDatabases().Label("EXP_GSTSLF RTL_DE_SRS")
	EXP_KNEIFF_DLP_DE_SRS := azure.Database.SqlDatabases().Label("EXP_KNEIFF DLP_DE_SRS")
	EXP_GSTSLF_DLP_DE_SRS := azure.Database.SqlDatabases().Label("EXP_GSTSLF DLP_DE_SRS")
	EXP_KNEIFF_SAP := azure.Database.SqlDatabases().Label("EXP_KNEIFF_SAP")
	EXP_GSTSLF_SAP := azure.Database.SqlDatabases().Label("EXP_GSTSLF_SAP")
	s022_kneiff_nostro := azure.Database.SqlDatabases().Label("s022_kneiff_nostro")
	EXP_KNEIFF_SRAP := azure.Database.SqlDatabases().Label("EXP_KNEIFF_SRAP")
	EXP_GSTSLF_RTL_AT_SRS := azure.Database.SqlDatabases().Label("EXP_GSTSLF_RTL_AT_SRS")
	STG_KNEIFF_GSWS_AT := azure.Database.SqlDatabases().Label("STG_KNEIFF_GSWS_AT")
	STG_GSTSLF_GSWS_AT := azure.Database.SqlDatabases().Label("STG_GSTSLF_GSWS_AT")
	STG_KNEIFF_DLP_AT_SRS := azure.Database.SqlDatabases().Label("STG_KNEIFF_DLP_AT_SRS")
	STG_GSTSLF_DLP_AT_SRS := azure.Database.SqlDatabases().Label("STG_GSTSLF_DLP_AT_SRS")
	STG_KNEIFF_RTL_AT_SRS := azure.Database.SqlDatabases().Label("STG_KNEIFF_RTL_AT_SRS")
	STG_GSTSLF_RTL_AT_SRS := azure.Database.SqlDatabases().Label("STG_GSTSLF_RTL_AT_SRS")
	EXP_GSTSLF_GTS := azure.Database.SqlDatabases().Label("EXP_GSTSLF_GTS")

	//create crefo_enrich_node
	crefo_enrich := azure.Database.SqlServers().Label("crefo_enrich")

	//create crefo_output_node
	REQ_CREFO_PREP := azure.Database.SqlDatabases().Label("REQ_CREFO_PREP")

	//connect nodes
	d.Connect(EXP_KNEIFF_GSWS, crefo_enrich)
	d.Connect(EXP_GSTSLF_GSWS, crefo_enrich)
	d.Connect(EXP_KNEIFF_RTL_DE_SRS, crefo_enrich)
	d.Connect(EXP_GSTSLF_RTL_DE_SRS, crefo_enrich)
	d.Connect(EXP_KNEIFF_DLP_DE_SRS, crefo_enrich)
	d.Connect(EXP_GSTSLF_DLP_DE_SRS, crefo_enrich)
	d.Connect(EXP_KNEIFF_SAP, crefo_enrich)
	d.Connect(EXP_GSTSLF_SAP, crefo_enrich)
	d.Connect(s022_kneiff_nostro, crefo_enrich)
	d.Connect(EXP_KNEIFF_SRAP, crefo_enrich)
	d.Connect(EXP_GSTSLF_RTL_AT_SRS, crefo_enrich)
	d.Connect(STG_KNEIFF_GSWS_AT, crefo_enrich)
	d.Connect(STG_GSTSLF_GSWS_AT, crefo_enrich)
	d.Connect(STG_KNEIFF_DLP_AT_SRS, crefo_enrich)
	d.Connect(STG_GSTSLF_DLP_AT_SRS, crefo_enrich)
	d.Connect(STG_KNEIFF_RTL_AT_SRS, crefo_enrich)
	d.Connect(STG_GSTSLF_RTL_AT_SRS, crefo_enrich)
	d.Connect(EXP_GSTSLF_GTS, crefo_enrich)

	d.Connect(crefo_enrich, REQ_CREFO_PREP)

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}
