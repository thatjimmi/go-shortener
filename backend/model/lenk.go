package model

func FindByShorten(shorten string) (*Lenk, error) {
	var lenk Lenk

	tx := db.Where("shorten = ?", shorten).First(&lenk)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &lenk, nil
}

func UpdateLenk(lenk *Lenk) error {
	tx := db.Save(lenk)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func GetAllLenks() ([]Lenk, error) {
	var lenks []Lenk

	tx := db.Find(&lenks)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return lenks, nil

}

func CreateLenk(lenk *Lenk) error {
	tx := db.Create(lenk)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
