package main

import (
    "log"
	"github.com/Foodle123/go-diagrams/nodes/azure"
    "github.com/Foodle123/go-diagrams/diagram"
)

func main() {
	d, err := diagram.New(diagram.Filename("010_ana_etl"), diagram.Label("AnaCredit"), diagram.Direction("TB"))
	if err != nil {
		log.Fatal(err)
	}

//create staging_files nodes
	acr_crefo_sd_sql := azure.Database.SqlDatabases().Label("acr_crefo_sd")
	exp_kneiff_srap_sql := azure.Database.SqlDatabases().Label("exp_kneiff_srap")
	exp_kneiff_sap_sql := azure.Database.SqlDatabases().Label("exp_kneiff_sap")
	s022_kneiff_nostro_sql := azure.Database.SqlDatabases().Label("s022_kneiff_nostro")
	exp_kneiff_gts_sql := azure.Database.SqlDatabases().Label("exp_kneiff_gts")
	s023_acr_national_kennung_sql := azure.Database.SqlDatabases().Label("s023_acr_national_kennung")
	exp_gstslf_sap_sql := azure.Database.SqlDatabases().Label("exp_gstslf_sap")
	exp_gstslf_gts_sql := azure.Database.SqlDatabases().Label("exp_gstslf_gts")
	exp_gstiff_sap_sql := azure.Database.SqlDatabases().Label("exp_gstiff_sap")
	exp_gstiff_gts_sql := azure.Database.SqlDatabases().Label("exp_gstiff_gts")
	s020_acr_param_sql := azure.Database.SqlDatabases().Label("s020_acr_param")

//create anaCredit extract group
	extract := diagram.NewGroup("extract").Label("Extract").Add(acr_crefo_sd_sql, exp_kneiff_srap_sql, exp_kneiff_sap_sql, s022_kneiff_nostro_sql, exp_kneiff_gts_sql, s023_acr_national_kennung_sql, exp_gstslf_sap_sql, exp_gstslf_gts_sql, exp_gstiff_sap_sql, exp_gstiff_gts_sql, s020_acr_param_sql)
	d.Group(extract)

//create enrich nodes
	procedure_accpif := azure.Database.SqlServers().Label("enrich ACCPIF")
	procedure_acfdif := azure.Database.SqlServers().Label("enrich ACFDIF")
	procedure_acfdif_upd := azure.Database.SqlServers().Label("enrich ACFDIF_UPD")
	procedure_accfif := azure.Database.SqlServers().Label("enrich ACCFIF")
	procedure_acfaif := azure.Database.SqlServers().Label("enrich ACFAIF")
	procedure_acpdif := azure.Database.SqlServers().Label("enrich ACPDIF")
	procedure_acipif := azure.Database.SqlServers().Label("enrich ACIPIF")
	procedure_acifif := azure.Database.SqlServers().Label("enrich ACIFIF")
	procedure_acpcif := azure.Database.SqlServers().Label("enrich ACPCIF")

//create enrich group
	einrich := diagram.NewGroup("enrich").Label("Transform").Add(procedure_accpif, procedure_acfdif, procedure_acfdif_upd, procedure_accfif, procedure_acfaif, procedure_acpdif, procedure_acipif, procedure_acifif, procedure_acpcif)
	d.Group(einrich)

//create load nodes
	exp_accpif_sql := azure.Database.SqlDatabases().Label("EXP_ACCPIF")
	exp_acfdif_sql := azure.Database.SqlDatabases().Label("EXP_ACFDIF")
	exp_accfif_sql := azure.Database.SqlDatabases().Label("EXP_ACCFIF")
	exp_acfaif_sql := azure.Database.SqlDatabases().Label("EXP_ACFAIF")
	exp_acpdif_sql := azure.Database.SqlDatabases().Label("EXP_ACPDIF")
	exp_acipif_sql := azure.Database.SqlDatabases().Label("EXP_ACIPIF")
	exp_acifif_sql := azure.Database.SqlDatabases().Label("EXP_ACIFIF")
	exp_acpcif_sql := azure.Database.SqlDatabases().Label("EXP_ACPCIF")

//create load group
	load := diagram.NewGroup("load").Label("Load").Add(exp_accpif_sql, exp_acfdif_sql,exp_accfif_sql, exp_acfaif_sql, exp_acpdif_sql, exp_acipif_sql, exp_acifif_sql, exp_acpcif_sql)
	d.Group(load)

//connects for ACCPIF

	d.Connect(acr_crefo_sd_sql, procedure_accpif)
	d.Connect(exp_kneiff_srap_sql, procedure_accpif)
	d.Connect(exp_kneiff_sap_sql, procedure_accpif)
	d.Connect(s022_kneiff_nostro_sql, procedure_accpif)
	d.Connect(exp_kneiff_gts_sql, procedure_accpif)
	d.Connect(s020_acr_param_sql, procedure_accpif)

	d.Connect(procedure_accpif, exp_accpif_sql)

//connects for ACFDIF
	d.Connect(exp_accpif_sql, procedure_acfdif)

	d.Connect(acr_crefo_sd_sql, exp_acfdif_sql)

//connects for ACFDIF_UPD
	d.Connect(exp_acfdif_sql, procedure_acfdif_upd)
	d.Connect(s020_acr_param_sql, procedure_acfdif_upd)
	d.Connect(s023_acr_national_kennung_sql, procedure_acfdif_upd)
	d.Connect(exp_accpif_sql, procedure_acfdif_upd)

//connects for ACCFIF
	d.Connect(exp_acfdif_sql, procedure_accfif)
	d.Connect(exp_accpif_sql, procedure_accfif)

	d.Connect(procedure_accfif, exp_accfif_sql)




	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}
