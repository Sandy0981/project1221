package models

import (
	"context"
)

// Define the function CreatInventory, which belongs to the struct 'Conn'.
// This function takes in 3 parameters: a context `ctx` of type `Context`, `ni` of type `NewInventory`, and `userId` of type `uint`.
// This function will return an `Inventory` and an `error`.

func (d *Conn) CreateCompany(ctx context.Context, ni NewCompany, userId int) (Company, error) {
	// Create a new 'Inventory' struct named 'inv'.
	// Initialize it with parameters from the 'NewInventory' struct and the `userId` passed to the function.
	cmp := Company{
		CompanyName: ni.CompanyName,
		FoundedYear: ni.FoundedYear,
		Location:    ni.Location,
		//Jobs:        ni.Jobs,
	}
	err := d.db.Create(&cmp).Error
	if err != nil {
		return Company{}, err
	}

	// Successfully created the record, return the user.
	return cmp, nil

}

func (s *Conn) ViewCompanyAll(ctx context.Context, companyId string) ([]Company, error) {
	var cmp = make([]Company, 0, 10)
	tx := s.db.Find(&cmp)
	err := tx.Find(&cmp).Error
	if err != nil {
		return nil, err
	}
	return cmp, nil
}

func (s *Conn) ViewJob(ctx context.Context, companyId string) ([]Job, error) {
	var cmp = make([]Job, 0, 10)
	tx := s.db.Find(&cmp)
	err := tx.Find(&cmp).Error
	if err != nil {
		return nil, err
	}
	return cmp, nil
}
