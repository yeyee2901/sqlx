package datasource

import (

	"github.com/yeyee2901/sqlx/app/config"

	"github.com/jmoiron/sqlx"
)

type DataSource struct {
	Config *config.Config
	DB     *sqlx.DB
}

func NewDatasource(config *config.Config, db *sqlx.DB) *DataSource {
	return &DataSource{
		Config: config,
		DB:     db,
	}
}

// kalau select nya semua, bisa langsung pakai .Select(), karena bisa langsung
// di bind ke struct array. kasus ini untuk yang expected return result nya > 1
func (T *DataSource) GetAllUsers(users interface{}) (err error){
	q := `
        SELECT id,name,created_at FROM users;
    `
	if err := T.DB.Select(users, q); err != nil {
		return err
	}
	return 
}

// untuk pengambilan data yang sifatnya pasti kembali 1 row, maka lebih
// baik pakai .Get() , lalu di query nya dikasih LIMIT 1 juga supaya lebih
// aman. Asumsinya kalao .Get() mentalin error, berarti gaada data yang match
func (T *DataSource) GetUserById(user interface{}, id string) (err error) {
	var args []interface{}
	q := `
        SELECT id,name,created_at FROM users
        WHERE id=?
    `
	args = append(args, id)
	err = T.DB.Get(user, q, args...)

	return
}

// untuk data insertion, ada baiknya menggunakan MustBegin() ini adalah wrapper
// yg lebih robust (aslinya kan sql, yg lebih banyak dipake sekarang sqlx)
// kalau transaction nya gagal, berarti gin nya langsung panic (if the
// transactions are failing 1 after the other the app wont be usable anyway)
// Rollback() = cancel database transaction
func (T *DataSource) CreateUser(newUser interface{}) (userId int64, err error) {
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
func (T *DataSource) UpdateUserById(updatedData interface{}) (err error) {
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

// implementasi lain bisa pakai MustExec, argumen nya bentuknya di
// kasih placeholder berupa '?'
func (T *DataSource) DeleteUserById(id int) (rowsAffected int64, err error){
	tx := T.DB.MustBegin()
	q := `
        DELETE FROM users 
        WHERE id = ?
    `
	result := tx.MustExec(q, id)

    // jangan lupa di commit :)
    // jangan seperti seseorang yang tidak commit dalam relationship nya :)))
    err = tx.Commit()
    if err != nil {
        tx.Rollback()
        return
    }

    rowsAffected, err = result.RowsAffected()
    if err != nil {
        tx.Rollback()
        return
    }
	return
}
