package main

import (
    "log"
	"github.com/Foodle123/go-diagrams/nodes/azure"
    "github.com/Foodle123/go-diagrams/diagram"
)

func main() {
	d, err := diagram.New(diagram.Filename("FSL"), diagram.Label("FSL"), diagram.Direction("TB"))
	if err != nil {
		log.Fatal(err)
	}

//create source_files nodes
	kneiff_fsl_csv := azure.General.CSVFile().Label("fsl_kneiff")
	gstiff_fsl_csv := azure.General.CSVFile().Label("fsl_gstiff")
	gstslf_fsl_csv := azure.General.CSVFile().Label("fsl_gstslf")

//create mft group
	mft_folder := diagram.NewGroup("mft").Label("mft").Add(kneiff_fsl_csv, gstiff_fsl_csv, gstslf_fsl_csv)
	d.Group(mft_folder)

//create staging_files nodes
	STG_KNEIFF_FSL_sql := azure.Database.SqlDatabases().Label("STG_KNEIFF_FSL")
	STG_GSTSLF_FSL_sql := azure.Database.SqlDatabases().Label("STG_GSTSLF_FSL")
	STG_GSTIFF_FSL_sql := azure.Database.SqlDatabases().Label("STG_GSTIFF_FSL")

//connect source_files with staging_files
	d.Connect(kneiff_fsl_csv, STG_KNEIFF_FSL_sql)
	d.Connect(gstiff_fsl_csv, STG_GSTSLF_FSL_sql)
	d.Connect(gstslf_fsl_csv, STG_GSTIFF_FSL_sql)

//create fsl staging group
	fsl_staging := diagram.NewGroup("staging").Label("Extract").Add(STG_KNEIFF_FSL_sql, STG_GSTSLF_FSL_sql, STG_GSTIFF_FSL_sql)
	d.Group(fsl_staging)

//create enrich nodes
	procedure_kneiff := azure.Database.SqlServers().Label("enrich kneiff")
	procedure_gstiff := azure.Database.SqlServers().Label("enrich gstiff")
	procedure_gstslf := azure.Database.SqlServers().Label("enrich gstslf")

//create enrich group
	einrich_group := diagram.NewGroup("enrich").Label("Transform").Add(procedure_kneiff, procedure_gstiff, procedure_gstslf)
	d.Group(einrich_group)

//connect enrich procedures with staging_files
	d.Connect(STG_KNEIFF_FSL_sql, procedure_kneiff)

	d.Connect(STG_GSTSLF_FSL_sql, procedure_gstslf)

	d.Connect(STG_GSTIFF_FSL_sql, procedure_gstiff)

//create export nodes
	EXP_KNEIFF_FSL_sql := azure.Database.SqlDatabases().Label("EXP_KNEIFF_FSL")
	EXP_GSTSLF_FSL_sql := azure.Database.SqlDatabases().Label("EXP_GSTSLF_FSL")
	EXP_GSTIFF_FSL_sql := azure.Database.SqlDatabases().Label("EXP_GSTIFF_FSL")

//create fsl export group
	fsl_export := diagram.NewGroup("export").Label("Load").Add(EXP_KNEIFF_FSL_sql, EXP_GSTSLF_FSL_sql, EXP_GSTIFF_FSL_sql)
	d.Group(fsl_export)

//connect enrich procedures export_files
	d.Connect(procedure_kneiff, EXP_KNEIFF_FSL_sql)
	d.Connect(procedure_gstiff, EXP_GSTSLF_FSL_sql)
	d.Connect(procedure_gstslf, EXP_GSTIFF_FSL_sql)

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}
