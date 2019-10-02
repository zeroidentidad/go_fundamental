package readconfig

import (
	"../mod-asisten-core/models"
	"github.com/jinzhu/gorm"
)

func a() {
	models.GetAllGrupKeys(new(gorm.DB))
}
