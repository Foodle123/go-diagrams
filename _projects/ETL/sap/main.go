package main

import (
    "log"
	"github.com/Foodle123/go-diagrams/nodes/azure"
    "github.com/Foodle123/go-diagrams/diagram"
)

func main() {
	d, err := diagram.New(diagram.Filename("SAP"), diagram.Label("SAP"), diagram.Direction("SAP"))
	if err != nil {
		log.Fatal(err)
	}


//create source_files nodes
	sap_eu_c_kneiff_txt := azure.General.TXTFile().Label("sap_kneiff")
	sap_eu_s_biliff_txt := azure.General.TXTFile().Label("sap_biliff")
	sap_eu_c_gstiff_txt := azure.General.TXTFile().Label("sap_gstiff_c")
	sap_eu_s_gstiff_txt := azure.General.TXTFile().Label("sap_gstiff_s")
	sap_eu_c_gstslf_txt := azure.General.TXTFile().Label("sap_gstslf_c")
	sap_eu_s_gstslf_txt := azure.General.TXTFile().Label("sap_gstslf_s")

//create mft group
	mft_folder := diagram.NewGroup("mft").Label("mft").Add(sap_eu_c_kneiff_txt, sap_eu_s_biliff_txt, sap_eu_c_gstiff_txt, sap_eu_s_gstiff_txt, sap_eu_c_gstslf_txt, sap_eu_s_gstslf_txt)
	d.Group(mft_folder)

//create staging_files nodes
	STG_KNEIFF_SAP_sql := azure.Database.SqlDatabases().Label("STG_KNEIFF_SAP")
	STG_BILIFF_SAP_sql := azure.Database.SqlDatabases().Label("STG_BILIFF_SAP")
	STG_GSTIFF_SAP_sql_C := azure.Database.SqlDatabases().Label("STG_GSTIFF_SAP_C")
	STG_GSTIFF_SAP_sql_S := azure.Database.SqlDatabases().Label("STG_GSTIFF_SAP_S")
	STG_GSTSLF_SAP_sql_C := azure.Database.SqlDatabases().Label("STG_GSTSLF_SAP_C")
	STG_GSTSLF_SAP_sql_S := azure.Database.SqlDatabases().Label("STG_GSTSLF_SAP_S")

//connect source_files with staging_files
	d.Connect(sap_eu_c_kneiff_txt, STG_KNEIFF_SAP_sql)
	d.Connect(sap_eu_s_biliff_txt, STG_BILIFF_SAP_sql)
	d.Connect(sap_eu_c_gstiff_txt, STG_GSTSLF_SAP_sql_C)
	d.Connect(sap_eu_s_gstiff_txt, STG_GSTSLF_SAP_sql_S)
	d.Connect(sap_eu_c_gstslf_txt, STG_GSTIFF_SAP_sql_C)
	d.Connect(sap_eu_s_gstslf_txt, STG_GSTIFF_SAP_sql_S)

//create sap staging group
	sap_staging := diagram.NewGroup("staging").Label("Extract").Add(STG_KNEIFF_SAP_sql, STG_BILIFF_SAP_sql, STG_GSTIFF_SAP_sql_C, STG_GSTIFF_SAP_sql_S, STG_GSTSLF_SAP_sql_C, STG_GSTSLF_SAP_sql_S)
	d.Group(sap_staging)

//create enrich nodes
	procedure_kneiff := azure.Database.SqlServers().Label("enrich kneiff")
	procedure_biliff := azure.Database.SqlServers().Label("enrich biliff")
	procedure_gstiff := azure.Database.SqlServers().Label("enrich gstiff")
	procedure_gstslf := azure.Database.SqlServers().Label("enrich gstslf")


//create enrich group
	einrich_group := diagram.NewGroup("enrich").Label("Transform").Add(procedure_kneiff, procedure_biliff, procedure_gstiff, procedure_gstslf)
	d.Group(einrich_group)


//connect enrich procedures with staging_files
	d.Connect(STG_KNEIFF_SAP_sql, procedure_kneiff)
	d.Connect(STG_BILIFF_SAP_sql, procedure_biliff)

	d.Connect(STG_GSTIFF_SAP_sql_C, procedure_gstiff)
	d.Connect(STG_GSTIFF_SAP_sql_S, procedure_gstiff)

	d.Connect(STG_GSTSLF_SAP_sql_C, procedure_gstiff)
	d.Connect(STG_GSTSLF_SAP_sql_S, procedure_gstiff)
	d.Connect(STG_GSTSLF_SAP_sql_C, procedure_gstslf)
	d.Connect(STG_GSTSLF_SAP_sql_S, procedure_gstslf)




	


//create export nodes
	EXP_KNEIFF_SAP_sql := azure.Database.SqlDatabases().Label("EXP_KNEIFF_SAP")
	EXP_BILIFF_SAP_sql := azure.Database.SqlDatabases().Label("EXP_BILIFF_SAP")
	EXP_GSTIFF_SAP_sql := azure.Database.SqlDatabases().Label("EXP_GSTIFF_SAP")
	EXP_GSTSLF_SAP_sql := azure.Database.SqlDatabases().Label("EXP_GSTSLF_SAP")



//create sap export group
	sap_export := diagram.NewGroup("export").Label("Load").Add(EXP_KNEIFF_SAP_sql, EXP_GSTSLF_SAP_sql, EXP_GSTIFF_SAP_sql,EXP_BILIFF_SAP_sql)
	d.Group(sap_export)

//connect enrich procedures export_files
	d.Connect(procedure_kneiff, EXP_KNEIFF_SAP_sql)
	d.Connect(procedure_biliff, EXP_BILIFF_SAP_sql)
	d.Connect(procedure_gstiff, EXP_GSTSLF_SAP_sql)
	d.Connect(procedure_gstslf, EXP_GSTIFF_SAP_sql)




	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}
