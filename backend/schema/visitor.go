package schema

//Visitor represents visitors table in database
type Visitor struct {
	NoIdentitas     uint64 `gorm:"uniqueIndex" json:"no_identitas"`
	Nama            string `json:"nama"`
	TempatLahir     string `json:"tempat_lahir"`
	TglLahir        uint64 `json:"tgl_lahir"`
	JenisKelamin    string `json:"jenis_kelamin"`
	Alamat          string `json:"alamat"`
	Handphone       string `json:"handphone"`
	TlpRumah        string `json:"tlp_rumah"`
	Email           string `gorm:"uniqueIndex" json:"email"`
	KdPropinsi      string `json:"kd_propinsi"`
	KdKotaKabupaten string `json:"kd_kota_kabupaten"`
	KdKecamatan     string `json:"kd_kecamatan"`
	KdKelurahan     string `json:"kd_kelurahan"`
	Kodepos         string `json:"kodepos"`
	KdJenisId       string `json:"kd_jenis_id"`
	PhotoDiriktp    string `json:"photo_diri_ktp"`
	Password        string `gorm:"->;<-;not null" json:"-"`
	Token           string `gorm:"-" json:"token,omitempty"`
}
