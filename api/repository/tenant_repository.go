package repository

import (
	"github.com/Dev4w4n/admin.e-masjid.my/api/model"
	"gorm.io/gorm"
)

type TenantRepository interface {
	FindAll() ([]model.Tenant, error)
	FindById(id int64) (model.Tenant, error)
	Upsert(tenant *model.Tenant) (model.Tenant, error)
	Delete(id int64) error
}

type TenantRepositoryImpl struct {
	Db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) TenantRepository {
	db.AutoMigrate(&model.Tenant{})
	return &TenantRepositoryImpl{Db: db}
}

func (repo *TenantRepositoryImpl) FindAll() ([]model.Tenant, error) {
	var tenantList []model.Tenant
	result := repo.Db.Find(&tenantList)

	if result.Error != nil {
		return nil, result.Error
	}
	return tenantList, nil
}

func (repo *TenantRepositoryImpl) FindById(id int64) (model.Tenant, error) {
	var tenant model.Tenant
	result := repo.Db.Where("id = ?", id).First(&tenant)

	if result.Error != nil {
		return model.Tenant{}, result.Error
	}

	return tenant, nil
}

func (repo *TenantRepositoryImpl) Delete(id int64) error {
	var tenant model.Tenant
	result := repo.Db.Where("id = ?", id).Delete(&tenant)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *TenantRepositoryImpl) Upsert(tenant *model.Tenant) (model.Tenant, error) {
	var result *gorm.DB

	if tenant.Id != 0 {
		result = repo.Db.Model(&model.Tenant{}).Where("id = ?", tenant.Id).Updates(&tenant)
	} else {
		result = repo.Db.Create(&tenant)
	}

	if result.Error != nil {
		return model.Tenant{}, result.Error
	}

	_tenant, err := repo.FindById(tenant.Id)
	if err != nil {
		return model.Tenant{}, err
	}

	return _tenant, nil
}
