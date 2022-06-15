package manager

import (
	"enigmacamp.com/goorm/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Infra interface {
	DbConn() *gorm.DB
}
type infra struct {
	db  *gorm.DB
	cfg config.Config
}

func (i *infra) initDb() {
	//dbHost := "167.172.69.254"
	//dbPort := "5432"
	//dbUser := "smm2"
	//dbPassword := "batchTwo"
	//dbName := "smm2"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", i.cfg.Host, i.cfg.User, i.cfg.Password, i.cfg.Name, i.cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "wmb.",
		}})
	if err != nil {
		panic(err)
	}
	enigmaDb, err := db.DB()
	err = enigmaDb.Ping()
	if err != nil {
		panic(err)
	}
	i.db = db
}

func (i *infra) DbConn() *gorm.DB {
	return i.db
}
func NewInfra(config config.Config) Infra {
	infra := new(infra)
	infra.cfg = config
	infra.initDb()
	return infra
}
