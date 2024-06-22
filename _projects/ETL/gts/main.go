package main

import (
    "log"
	"github.com/Foodle123/go-diagrams/nodes/azure"
    "github.com/Foodle123/go-diagrams/diagram"
)

func main() {
	d, err := diagram.New(diagram.Filename("GTS"), diagram.Label("GTS"), diagram.Direction("TB"))
	if err != nil {
		log.Fatal(err)
	}

//create mft node
//	filesOrigin := azure.Analytics.DataLakeStoreGen1().Label("mft")

//create source_files nodes
	gts_kneiff_txt := azure.General.TXTFile().Label("gts_kneiff")
	gts_gstiff_txt := azure.General.TXTFile().Label("gts_gstiff")
	gts_gstslf_txt := azure.General.TXTFile().Label("gts_gstslf")

//create mft group
	mft_folder := diagram.NewGroup("mft").Label("mft").Add(gts_kneiff_txt, gts_gstiff_txt, gts_gstslf_txt)
	d.Group(mft_folder)

//create staging_files nodes
	STG_KNEIFF_GTS_sql := azure.Database.SqlDatabases().Label("STG_KNEIFF_GTS")
	STG_GSTSLF_GTS_sql := azure.Database.SqlDatabases().Label("STG_GSTSLF_GTS")
	STG_GSTIFF_GTS_sql := azure.Database.SqlDatabases().Label("STG_GSTIFF_GTS")

//create fis staging group
	gis_staging := diagram.NewGroup("staging").Label("Extract").Add(STG_KNEIFF_GTS_sql, STG_GSTSLF_GTS_sql, STG_GSTIFF_GTS_sql)
	d.Group(gis_staging)

//connect source_files with staging_files
	d.Connect(gts_kneiff_txt, STG_KNEIFF_GTS_sql)
	d.Connect(gts_gstiff_txt, STG_GSTSLF_GTS_sql)
	d.Connect(gts_gstslf_txt, STG_GSTIFF_GTS_sql)








//create enrich nodes
	procedure_kneiff := azure.Database.SqlServers().Label("enrich kneiff")
	procedure_gstiff := azure.Database.SqlServers().Label("enrich gstiff")
	procedure_gstslf := azure.Database.SqlServers().Label("enrich gstslf")

//create enrich group
	einrich_group := diagram.NewGroup("enrich").Label("Transform").Add(procedure_kneiff, procedure_gstiff, procedure_gstslf)
	d.Group(einrich_group)

//create export nodes
	EXP_KNEIFF_GTS_sql := azure.Database.SqlDatabases().Label("EXP_KNEIFF_GTS")
	EXP_GSTSLF_GTS_sql := azure.Database.SqlDatabases().Label("EXP_GSTSLF_GTS")
	EXP_GSTIFF_GTS_sql := azure.Database.SqlDatabases().Label("EXP_GSTIFF_GTS")

//create fis export group
	gts_export := diagram.NewGroup("export").Label("Load").Add(EXP_KNEIFF_GTS_sql, EXP_GSTSLF_GTS_sql, EXP_GSTIFF_GTS_sql)
	d.Group(gts_export)

//connect enrich procedures with staging_files
	d.Connect(STG_KNEIFF_GTS_sql, procedure_kneiff)

	d.Connect(EXP_GSTIFF_GTS_sql, procedure_gstiff)
	d.Connect(STG_GSTSLF_GTS_sql, procedure_gstslf)

	d.Connect(STG_GSTIFF_GTS_sql, procedure_gstiff)



//connect enrich procedures export_files
	d.Connect(procedure_kneiff, EXP_KNEIFF_GTS_sql)
	d.Connect(procedure_gstiff, EXP_GSTSLF_GTS_sql)
	d.Connect(procedure_gstslf, EXP_GSTIFF_GTS_sql)


	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}
