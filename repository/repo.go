package repository

import (
	"gorm.io/gorm"
	"jawaban/service"
)

type Area struct {    //penulisan type struct harus seperti ini
	ID int64 `gorm:"column:id;primaryKey;"`
	AreaValue float64 `gorm:"column:area_value"`
	AreaType string `gorm:"column:type"`
}

type areaRepository struct {
	DB *gorm.DB
}

//NewGormDBRepository Generate Gorm DB user repository
func NewGormDBRepository(db *gorm.DB)service.Repository {
	return &areaRepository{
		db,
	}
}

// func (_r *AreaRepository) InsertArea(param1 int32, param2 int64, types []string, ar *Area)	(err error)
func (_r *areaRepository) InsertArea(param1 int64, param2 int64, _type string)	(err error) { /* tulisan  type di tambah _ karena tidak bisa digunakan, di hilangkan paremeter pemanggilan struct model*/
			//inst := _r.DB.Model(ar) // tidak usah dipakai karena tidak di pakai dimana-mana
			var ar Area // mendeklarasikan variabel struct model
			var area int64  // var bertulis V besar harus huruf kecil karena bukan termasuk nama variabel dan sesuaikan tipe date dengan parameter input int 64

			switch _type {
				case "persegi panjang":
					area = param1 * param2  // lebih baik variabel area tidak mendeklarasikan var lagi karena sudah di tulis di line sebelumnya
					ar.AreaValue = float64(area) //di rubah float 64 agar general support saat di pakai di berbagai kondisi
					ar.AreaType = _type
					err = _r.DB.Create(&ar).Error // tulisan create diganti Create
					if err != nil {
					return err
					}
				case "persegi":
					area = param1 * param2 // lebih baik variabel area tidak mendeklarasikan var lagi karena sudah di tulis di line sebelumnya
					ar.AreaValue = float64(area) //di rubah float 64 agar general support saat di pakai di berbagai kondisi
					ar.AreaType = "persegi"
					err = _r.DB.Create(&ar).Error // tulisan create diganti Create
					if err != nil {
					return err
				}
				case "segitiga":
					area = param1 * param2
					ar.AreaValue = 0.5 *  float64(area)
					ar.AreaType = "segitiga"
					err = _r.DB.Create(&ar).Error // tulisan create diganti Create
					if err != nil {
					return err
					}
				default:
					ar.AreaValue = 0
					ar.AreaType = "undefined data"
					err = _r.DB.Create(&ar).Error // tulisan create diganti Create
					if err != nil {
					return err
					}
			}
			return nil
}