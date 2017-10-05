package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type ResponseReg struct {
	Data []Region `json:"data"`
}

type ResponseMuseum struct {
	Data []Museum `json:"data"`
}

type Region struct {
	Nama string `json:"nama"`
	Kode string `json:"kode_wilayah"`
}

type Museum struct {
	ID                string `json:"museum_id"`
	KodePengelola     string `json:"kode_pengelolaan"`
	Nama              string `json:"nama"`
	SDM               string `json:"sdm"`
	AlamatJalan       string `json:"alamat_jalan"`
	DesaKelurahan     string `json:"desa_kelurahan"`
	Kecamatan         string `json:"kecamatan"`
	KabupatenKota     string `json:"kabupaten_kota"`
	Propinsi          string `json:"propinsi"`
	Lintang           string `json:"lintang"`
	Bujur             string `json:"bujur"`
	Koleksi           string `json:"koleksi"`
	SumberDana        string `json:"sumber_dana"`
	Pengelola         string `json:"pengelola"`
	Tipe              string `json:"tipe"`
	Standar           string `json:"standar"`
	TahunBerdiri      string `json:"tahun_berdiri"`
	Bangunan          string `json:"bangunan"`
	LuasTanah         string `json:"luas_tanah"`
	StatusKepemilikan string `json:"status_kepemilikan"`
}

var dirPath *string
var Limit *int

func main() {
	dirPath = flag.String("output", "./Kota", "a String")
	Limit = flag.Int("concurrent_limit", 5, "an int")
	flag.Parse()

	urlProv := "http://jendela.data.kemdikbud.go.id/api/index.php/CWilayah/wilayahGET"
	saveCSV(urlProv)
}

func saveCSV(urlProv string) {
	if _, err := os.Stat(*dirPath); os.IsNotExist(err) {
		os.Mkdir(*dirPath, os.ModePerm)
	}
	bodyProv := getBody(urlProv)
	dataProv := parseBodyReg(bodyProv)
	getKota(dataProv, urlProv)
}

func getKota(dataProv []Region, urlProv string) {
	var dataKota []Region

	log.SetFlags(log.Ltime) // format log output hh:mm:ss

	wg := sync.WaitGroup{}
	queue := make(chan []Region)

	for worker := 0; worker < *Limit; worker++ {
		wg.Add(1)

		go func(worker int) {
			defer wg.Done()

			for work := range queue {
				getMuseum(work) // blocking wait for work
			}
		}(worker)
	}

	for _, p := range dataProv {
		urlKota := urlProv + "?mst_kode_wilayah=" + p.Kode
		bodyKota := getBody(urlKota)
		dataKota = parseBodyReg(bodyKota)
		queue <- dataKota
		// go getMuseum(dataKota)
	}
}

func getMuseum(dataKota []Region) {
	for _, val := range dataKota {
		file, err := os.Create(*dirPath + "/" + val.Nama + ".csv")
		checkError("Cannot create file!", err)
		defer file.Close()

		w := csv.NewWriter(file)
		defer w.Flush()

		urlDataMuseum := "http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/searchGET?kode_kab_kota=" + val.Kode
		bodyMuseum := getBody(urlDataMuseum)
		dataMuseum := parseBodyMuseum(bodyMuseum)

		for _, m := range dataMuseum {
			var record []string
			record = append(record, m.ID)
			record = append(record, m.KodePengelola)
			record = append(record, m.Nama)
			record = append(record, m.SDM)
			record = append(record, m.AlamatJalan)
			record = append(record, m.DesaKelurahan)
			record = append(record, m.Kecamatan)
			record = append(record, m.KabupatenKota)
			record = append(record, m.Propinsi)
			record = append(record, m.Lintang)
			record = append(record, m.Bujur)
			record = append(record, m.Koleksi)
			record = append(record, m.SumberDana)
			record = append(record, m.Pengelola)
			record = append(record, m.Tipe)
			record = append(record, m.Standar)
			record = append(record, m.TahunBerdiri)
			record = append(record, m.Bangunan)
			record = append(record, m.LuasTanah)
			record = append(record, m.StatusKepemilikan)
			err = w.Write(record)
			checkError("Cannot write to file", err)
		}
	}
}

func parseBodyMuseum(body []byte) []Museum {
	var r ResponseMuseum
	if err := json.Unmarshal(body, &r); err != nil {
		fmt.Println(err)
	}

	museum := r.Data
	return museum
}

func parseBodyReg(body []byte) []Region {
	var r ResponseReg
	if err := json.Unmarshal(body, &r); err != nil {
		fmt.Println(err)
	}

	prov := r.Data
	return prov
}

func getBody(url string) []byte {
	museumClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
	}

	res, err := museumClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	return body
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
