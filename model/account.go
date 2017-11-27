package model

import "database/sql"

//Account  (TYPE)
type Account struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//Accounts (TYPE)
type Accounts struct {
	Accounts []*Account `json:"accounts"`
}

//DoesAccountResourceExist (POST)
func DoesAccountResourceExist(account *Account) bool {

	err := db.QueryRow("SELECT id, email, password FROM auth_accounts WHERE email=?", account.Email).Scan(&account.ID, &account.Email, &account.Password)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesAccountIDExist (POST)
func DoesAccountIDExist(ID int) bool {

	var id int
	err := db.QueryRow("SELECT id FROM auth_accounts WHERE id=?", ID).Scan(&id)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesAccountEmailExistForAnotherID (PUT)
func DoesAccountEmailExistForAnotherID(email string, id int) bool {

	var dbID int
	err := db.QueryRow("SELECT id FROM auth_accounts WHERE email=?", email).Scan(&dbID)

	if err == sql.ErrNoRows {
		return false
	}

	if dbID != id {
		return true
	}

	return false
}

//CreateAccount (POST)
func CreateAccount(account *Account) error {

	res, err := db.Exec("INSERT INTO auth_accounts VALUES(null, ?, ?)", account.Email, account.Password)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	account.ID = int(id)

	return nil
}

//GetAccounts (GET)
func GetAccounts() ([]Account, error) {
	rows, err := db.Query("SELECT id, email, password FROM auth_accounts")

	if err != nil {
		return nil, err
	}

	accounts := []Account{}

	for rows.Next() {
		defer rows.Close()

		var a Account
		if err := rows.Scan(&a.ID, &a.Email, &a.Password); err != nil {
			return nil, err
		}
		accounts = append(accounts, a)
	}

	return accounts, nil
}

//GetAccount (GET)
func GetAccount(account *Account) error {
	return db.QueryRow("SELECT email, password FROM auth_accounts WHERE id=?", account.ID).Scan(&account.Email, &account.Password)
}

//GetAccountByEmail (GET)
func GetAccountByEmail(account *Account) error {
	return db.QueryRow("SELECT id, email, password from auth_accounts where email=?",
		account.Email).Scan(&account.ID, &account.Email, &account.Password)
}

//UpdateAccount (PUT)
func UpdateAccount(account *Account) error {
	_, err :=
		db.Exec("UPDATE auth_accounts SET email=?,  password=? WHERE id=?", account.Email, account.Password, account.ID)

	return err
}

//DeleteAccount (DELETE)
func DeleteAccount(account *Account) error {
	_, err := db.Exec("DELETE FROM auth_accounts WHERE id=?", account.ID)

	return err
}
