package helper

import (
	"alurkerjaService/model/domain"
	"alurkerjaService/model/web"
)

func ToJenisDataResponse(jenisData domain.JenisData) web.JenisDataResponse {
	return web.JenisDataResponse{
		Id:        jenisData.Id,
		JenisData: jenisData.JenisData,
	}
}

func ToJenisDataResponses(jenisDatas []domain.JenisData) []web.JenisDataResponse {
	var responses []web.JenisDataResponse
	for _, jenisData := range jenisDatas {
		responses = append(responses, ToJenisDataResponse(jenisData))
	}
	return responses
}

func ToDataKinerjaPemdaResponse(dataKinerjaPemda domain.DataKinerjaPemda) web.DataKinerjaPemdaResponse {
	return web.DataKinerjaPemdaResponse{
		Id:                   dataKinerjaPemda.Id,
		JenisDataId:          dataKinerjaPemda.JenisDataId,
		JenisData:            dataKinerjaPemda.JenisData,
		NamaData:             dataKinerjaPemda.NamaData,
		RumusPerhitungan:     dataKinerjaPemda.RumusPerhitungan,
		SumberData:           dataKinerjaPemda.SumberData,
		InstansiProdusenData: dataKinerjaPemda.InstansiProdusenData,
		Target:               ToTargetResponses(dataKinerjaPemda.Target),
		Keterangan:           dataKinerjaPemda.Keterangan,
	}
}

func ToDataKinerjaPemdaResponses(dataKinerjaPemdas []domain.DataKinerjaPemda) []web.DataKinerjaPemdaResponse {
	var responses []web.DataKinerjaPemdaResponse
	for _, dataKinerjaPemda := range dataKinerjaPemdas {
		responses = append(responses, ToDataKinerjaPemdaResponse(dataKinerjaPemda))
	}
	return responses
}

func ToTargetResponse(target domain.Target) web.TargetResponse {
	return web.TargetResponse{
		Id:     target.Id,
		Target: target.Target,
		Satuan: target.Satuan,
		Tahun:  target.Tahun,
	}
}

func ToTargetResponses(targets []domain.Target) []web.TargetResponse {
	var responses []web.TargetResponse
	for _, target := range targets {
		responses = append(responses, ToTargetResponse(target))
	}
	return responses
}
