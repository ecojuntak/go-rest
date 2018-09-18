package model

func Migrate() (err error) {
	err = db.AutoMigrate(&User{}).Error
	return err
}
