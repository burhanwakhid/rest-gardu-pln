package model

type GarduM struct {
	ID           uint    `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Kva          string  `json:"kva"`
	JmlPlgn      string  `json:"jmlPlgn"`
	StatusNyala  string  `json:"statusNyala"`
	Distribusi   string  `json:"distribusi"`
	Area         string  `json:"area"`
	Unit         string  `json:"unit"`
	NoTiang      string  `json:"noTiang"`
	JenisTrafo   string  `json:"jenisTrafo"`
	X            float64 `json:"x"`
	Y            float64 `json:"y"`
	Coverage     string  `json:"coverage"`
	KodeLaporan  string  `json:"kodeLaporan"`
	KodeLaporanP string  `json:"kodeLaporanP"`
	KodeLaporanJ string  `json:"kodeLaporanJ"`
}
