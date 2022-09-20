package datasource

import (

	"github.com/yeyee2901/sqlx/app/config"

	"github.com/jmoiron/sqlx"
)

type Datasource struct {
	Config *config.Config
	DB     *sqlx.DB
}

func NewDatasource(config *config.Config, db *sqlx.DB) *Datasource {
	return &Datasource{
		Config: config,
		DB:     db,
	}
}

// kalau select nya semua, bisa langsung pakai .Select(), karena bisa langsung
// di bind ke struct array. kasus ini untuk yang expected return result nya > 1
func (T *Datasource) GetAllUsers() ([]User, error) {
	q := `
        SELECT id,name,created_at FROM users;
    `
	var users []User

	if err := T.DB.Select(&users, q); err != nil {
		return nil, err
	}

	return users, nil
}

// untuk pengambilan data yang sifatnya pasti kembali 1 row, maka lebih
// baik pakai .Get() , lalu di query nya dikasih LIMIT 1 juga supaya lebih
// aman. Asumsinya kalao .Get() mentalin error, berarti gaada data yang match
func (T *Datasource) GetUserById(id string) (result User, err error) {
	var args []interface{}
	q := `
        SELECT id,name,created_at FROM users
        WHERE id=? LIMIT 1
    `
	args = append(args, id)
	err = T.DB.Get(&result, q, args...)

	return
}

// untuk data insertion, ada baiknya menggunakan MustBegin() ini adalah wrapper
// yg lebih robust (aslinya kan sql, yg lebih banyak dipake sekarang sqlx)
// kalau transaction nya gagal, berarti gin nya langsung panic (if the
// transactions are failing 1 after the other the app wont be usable anyway)
// Rollback() = cancel database transaction
func (T *Datasource) CreateUser(newUser *CreateUserReq) (userId int64, err error) {
	tx := T.DB.MustBegin()
	q := `
        INSERT INTO users (name) VALUES
        (
            :name
        )
    `
	r, err := tx.NamedExec(q, newUser)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return
	}

	// get user ID (best practice backend, saat create user, kembalikan ID yg
	// di generate)
	userId, err = r.LastInsertId()
	return
}

// kasus nya sama kayak insert, bedanya query nya update saja
func (T *Datasource) UpdateUserById(updatedData *UpdateUserByIdReq) (err error) {
	tx := T.DB.MustBegin()
	q := `
        UPDATE users 
        SET name = :name
        WHERE id = :id
    `
	tx.NamedExec(q, updatedData)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return
	}

	return nil
}
