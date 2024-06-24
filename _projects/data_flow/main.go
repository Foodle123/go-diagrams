package main

import (
	"log"

	"github.com/Foodle123/go-diagrams/diagram"
	"github.com/Foodle123/go-diagrams/nodes/azure"
)

func main() {
	d, err := diagram.New(diagram.Filename("Dataflow"), diagram.Label("Dataflow"), diagram.Direction("LR"))
	if err != nil {
		log.Fatal(err)
	}

	wipro := azure.Compute.BatchAccounts().Label("wipro")

	//create source_files nodes
	fis_kneiff_csv_1 := azure.General.CSVFile().Label("example file01")
	gts_kneiff_csv_1 := azure.General.TXTFile().Label("example file02")
	dots_1 := azure.General.TXTFile().Label("example file03")

	//create mft group
	mft_folder := diagram.NewGroup("mft").Label("mft").Add(fis_kneiff_csv_1, gts_kneiff_csv_1, dots_1)
	d.Group(mft_folder)

	d.Connect(wipro, fis_kneiff_csv_1)
	d.Connect(wipro, gts_kneiff_csv_1)
	d.Connect(wipro, dots_1)

	//create checkfiles nodes
	checkfiles := azure.Compute.CloudServicesClassic().Label("check files")

	//create copyfiles nodes
	copyfiles_1 := azure.Compute.CloudServicesClassic().Label("copy files")

	//create source_files nodes temp
	fis_kneiff_csv_2 := azure.General.CSVFile().Label("example file01")
	gts_kneiff_csv_2 := azure.General.TXTFile().Label("example file02")
	dots_2 := azure.General.TXTFile().Label("example file03")

	//create temp group
	temp_folder := diagram.NewGroup("temp_folder").Label("temp folder").Add(fis_kneiff_csv_2, gts_kneiff_csv_2, dots_2)
	d.Group(temp_folder)

	//create copyfiles nodes
	copyfiles_2 := azure.Compute.CloudServicesClassic().Label("copy files")

	//connect enrich procedures with staging_files

	d.Connect(fis_kneiff_csv_1, checkfiles)
	d.Connect(gts_kneiff_csv_1, checkfiles)
	d.Connect(dots_1, checkfiles)

	d.Connect(checkfiles, copyfiles_1)

	d.Connect(fis_kneiff_csv_1, copyfiles_1)
	d.Connect(gts_kneiff_csv_1, copyfiles_1)
	d.Connect(dots_1, copyfiles_1)

	d.Connect(copyfiles_1, fis_kneiff_csv_2)
	d.Connect(copyfiles_1, gts_kneiff_csv_2)
	d.Connect(copyfiles_1, dots_2)

	d.Connect(fis_kneiff_csv_2, copyfiles_2)
	d.Connect(gts_kneiff_csv_2, copyfiles_2)
	d.Connect(dots_2, copyfiles_2)

	//create source_files nodes prod folder
	fis_kneiff_csv_3 := azure.General.CSVFile().Label("example file01")
	gts_kneiff_csv_3 := azure.General.TXTFile().Label("example file02")
	dots_3 := azure.General.TXTFile().Label("example file03")

	//create prod group
	prod_folder := diagram.NewGroup("production_folder").Label("production folder").Add(fis_kneiff_csv_3, gts_kneiff_csv_3, dots_3)
	d.Group(prod_folder)

	d.Connect(copyfiles_2, fis_kneiff_csv_3)
	d.Connect(copyfiles_2, gts_kneiff_csv_3)
	d.Connect(copyfiles_2, dots_3)

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}
