package model

type JsonGardu struct {
	ListAsset []struct {
		Kva          string `json:"kva,omitempty"`
		JmlPlgn      string `json:"jmlPlgn,omitempty"`
		StatusNyala  string `json:"statusNyala"`
		Distribusi   string `json:"distribusi"`
		Area         string `json:"area"`
		Unit         string `json:"unit"`
		NoTiang      string `json:"noTiang"`
		JenisTrafo   string `json:"jenisTrafo,omitempty"`
		X            string `json:"x"`
		Y            string `json:"y"`
		Coverage     string `json:"coverage"`
		KodeLaporan  string `json:"kodeLaporan"`
		KodeLaporanP string `json:"kodeLaporanP"`
		KodeLaporanJ string `json:"kodeLaporanJ"`
	} `json:"listAsset"`
}
