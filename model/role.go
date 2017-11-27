package model

import "database/sql"

//Role  (TYPE)
type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//Roles (TYPE)
type Roles struct {
	Roles []*Role `json:"roles"`
}

//DoesRoleResourceExist (POST)
func DoesRoleResourceExist(role *Role) bool {

	err := db.QueryRow("SELECT id, name FROM auth_roles WHERE name=?", role.Name).Scan(&role.ID, &role.Name)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesRoleIDExist (POST)
func DoesRoleIDExist(ID int) bool {

	var id int
	err := db.QueryRow("SELECT id FROM auth_roles WHERE id=?", ID).Scan(&id)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesRoleNameExistForAnotherID (PUT)
func DoesRoleNameExistForAnotherID(name string, id int) bool {

	var dbID int
	err := db.QueryRow("SELECT id FROM auth_roles WHERE name=?", name).Scan(&dbID)

	if err == sql.ErrNoRows {
		return false
	}

	if dbID != id {
		return true
	}

	return false
}

//CreateRole (POST)
func CreateRole(role *Role) error {

	res, err := db.Exec("INSERT INTO auth_roles VALUES(null, ?)", role.Name)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	role.ID = int(id)

	return nil
}

//GetRoles (GET)
func GetRoles() ([]Role, error) {
	rows, err := db.Query("SELECT id, name FROM auth_roles")

	if err != nil {
		return nil, err
	}

	roles := []Role{}

	for rows.Next() {
		defer rows.Close()

		var r Role
		if err := rows.Scan(&r.ID, &r.Name); err != nil {
			return nil, err
		}
		roles = append(roles, r)
	}

	return roles, nil
}

//GetRole (GET)
func GetRole(role *Role) error {
	return db.QueryRow("SELECT name FROM auth_roles WHERE id=?", role.ID).Scan(&role.Name)
}

//UpdateRole (PUT)
func UpdateRole(role *Role) error {
	_, err :=
		db.Exec("UPDATE auth_roles SET name=? WHERE id=?", role.Name, role.ID)

	return err
}

//DeleteRole (DELETE)
func DeleteRole(role *Role) error {
	_, err := db.Exec("DELETE FROM auth_roles WHERE id=?", role.ID)

	return err
}
