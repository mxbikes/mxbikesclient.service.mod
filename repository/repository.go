package repository

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/morkid/paginate"
	"github.com/mxbikes/mxbikesclient.service.mod/models"
	"gorm.io/gorm"
)

type ModRepository interface {
	SearchMod(modName string, listQuery *models.ListQuery) *paginate.Page
	GetModByID(id string) *models.Mod
	Migrate() error
}

type postgresRepository struct {
	db *gorm.DB
}

func NewRepository(c *gorm.DB) *postgresRepository {
	return &postgresRepository{db: c}
}

func (p *postgresRepository) GetModByID(id string) *models.Mod {
	var mod models.Mod
	p.db.First(&mod, "id = ?", id)
	return &mod
}

func (p *postgresRepository) SearchMod(modName string, listQuery *models.ListQuery) *paginate.Page {
	// List query to string
	vals, _ := query.Values(listQuery)

	model := p.db.Where("LOWER(name) LIKE LOWER(?)", modName).Model(&models.Mod{})
	pg := paginate.New(&paginate.Config{
		DefaultSize: 10,
	})
	page := pg.Response(model, &http.Request{URL: &url.URL{RawQuery: vals.Encode()}}, &[]models.Mod{})

	return &page
}

func (p *postgresRepository) Migrate() error {
	p.db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	return p.db.AutoMigrate(&models.Mod{})
}
