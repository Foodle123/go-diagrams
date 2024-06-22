package main

import (
    "log"
	"github.com/Foodle123/go-diagrams/nodes/azure"
    "github.com/Foodle123/go-diagrams/diagram"
)

func main() {
	d, err := diagram.New(diagram.Filename("FIS"), diagram.Label("FIS"), diagram.Direction("TB"))
	if err != nil {
		log.Fatal(err)
	}

//create mft node
//	filesOrigin := azure.Analytics.DataLakeStoreGen1().Label("mft")

//create source_files nodes
	fis_kneiff_csv := azure.General.CSVFile().Label("fis_kneiff")
	fis_gstiff_csv := azure.General.CSVFile().Label("fis_gstiff")
	fis_gstslf_csv := azure.General.CSVFile().Label("fis_gstslf")
	fis_zstiff_csv := azure.General.CSVFile().Label("fis_zstiff")

//create mft group
	mft_folder := diagram.NewGroup("mft").Label("mft").Add(fis_kneiff_csv, fis_gstiff_csv, fis_gstslf_csv, fis_zstiff_csv)
	d.Group(mft_folder)

//create staging_files nodes
	STG_KNEIFF_FIS_sql := azure.Database.SqlDatabases().Label("STG_KNEIFF_FIS")
	STG_GSTSLF_FIS_sql := azure.Database.SqlDatabases().Label("STG_GSTSLF_FIS")
	STG_GSTIFF_FIS_sql := azure.Database.SqlDatabases().Label("STG_GSTIFF_FIS")
	STG_ZSTIFF_FIS_sql := azure.Database.SqlDatabases().Label("STG_ZSTIFF_FIS")

//connect source_files with staging_files
	d.Connect(fis_kneiff_csv, STG_KNEIFF_FIS_sql)
	d.Connect(fis_gstiff_csv, STG_GSTSLF_FIS_sql)
	d.Connect(fis_gstslf_csv, STG_GSTIFF_FIS_sql)
	d.Connect(fis_zstiff_csv, STG_ZSTIFF_FIS_sql)

//create fis staging group
	fis_staging := diagram.NewGroup("staging").Label("Extract").Add(STG_KNEIFF_FIS_sql, STG_GSTSLF_FIS_sql, STG_GSTIFF_FIS_sql, STG_ZSTIFF_FIS_sql)
	d.Group(fis_staging)






//create enrich nodes
	procedure_kneiff := azure.Database.SqlServers().Label("enrich kneiff")
	procedure_gstiff := azure.Database.SqlServers().Label("enrich gstiff")
	procedure_gstslf := azure.Database.SqlServers().Label("enrich gstslf")
	procedure_zstiff := azure.Database.SqlServers().Label("enrich zstiff")

//create enrich group
	einrich_group := diagram.NewGroup("enrich").Label("Transform").Add(procedure_kneiff, procedure_gstiff, procedure_gstslf,procedure_zstiff)
	d.Group(einrich_group)
	//einrich_group.Heighth = 2

//connect enrich procedures with staging_files
	d.Connect(STG_KNEIFF_FIS_sql, procedure_kneiff)
	
	d.Connect(STG_GSTSLF_FIS_sql, procedure_gstiff)
	d.Connect(STG_GSTSLF_FIS_sql, procedure_gstslf)
	d.Connect(STG_GSTSLF_FIS_sql, procedure_zstiff)
	
	d.Connect(STG_GSTIFF_FIS_sql, procedure_gstiff)
	d.Connect(STG_GSTIFF_FIS_sql, procedure_gstslf)
	d.Connect(STG_GSTIFF_FIS_sql, procedure_zstiff)
	
	d.Connect(STG_ZSTIFF_FIS_sql, procedure_zstiff)


//create invis enrich group cluster for margin
//	invis_einrich_group := diagram.NewGroup("invis_enrich").Add(procedure_kneiff, procedure_gstiff, procedure_gstslf,procedure_zstiff)
//	d.Group(einrich_group)



//create export nodes
	EXP_KNEIFF_FIS_sql := azure.Database.SqlDatabases().Label("EXP_KNEIFF_FIS")
	EXP_GSTSLF_FIS_sql := azure.Database.SqlDatabases().Label("EXP_GSTSLF_FIS")
	EXP_GSTIFF_FIS_sql := azure.Database.SqlDatabases().Label("EXP_GSTIFF_FIS")
	EXP_ZSTIFF_FIS_sql := azure.Database.SqlDatabases().Label("EXP_ZSTIFF_FIS")

//create fis export group
	fis_export := diagram.NewGroup("export").Label("Load").Add(EXP_KNEIFF_FIS_sql, EXP_GSTSLF_FIS_sql, EXP_GSTIFF_FIS_sql,EXP_ZSTIFF_FIS_sql)
	d.Group(fis_export)

//connect enrich procedures export_files
	d.Connect(procedure_kneiff, EXP_KNEIFF_FIS_sql)
	d.Connect(procedure_gstiff, EXP_GSTSLF_FIS_sql)
	d.Connect(procedure_gstslf, EXP_GSTIFF_FIS_sql)
	d.Connect(procedure_zstiff, EXP_ZSTIFF_FIS_sql)



	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}
