package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/webserver/types"
	"github.com/webserver/utils"
	"golang.org/x/crypto/bcrypt"
)

var UserDat []types.UserData

func SignUp(DB *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var u types.UserData
		json.Unmarshal(reqBody, &u)
		bytes, err := bcrypt.GenerateFromPassword([]byte(u.Pswd), 14)
		if err != nil {
			fmt.Println("Error Occured")
			return
		}
		insert, err := DB.Exec("INSERT INTO user VALUES (?,?,?,?,?,?,?,?,?,?,?)", u.Email, string(bytes), u.FirstName, u.LastName, u.OrgName, u.Inst, u.BuildNo, u.FloorNo, u.LabHead, u.LabAddress, u.Tel)

		if err != nil {
			panic(err.Error())
		}
		n, err := insert.RowsAffected()
		if n == 1 {
			fmt.Println("1 row inserted.")
		}

		token, err := utils.CreateToken(u.Email)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}
		json.NewEncoder(w).Encode(token)

	}
}
func Login(DB *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var u types.UserData
		json.Unmarshal(reqBody, &u)
		query, err := DB.Query("SELECT * FROM user WHERE email = ?", u.Email)
		if err != nil {
			panic(err.Error())
		}

		for query.Next() {
			var data types.UserData
			err = query.Scan(&data.Email, &data.Pswd)
			if err != nil {
				panic(err.Error())
			}
			err := bcrypt.CompareHashAndPassword([]byte(data.Pswd), []byte(u.Pswd))
			if err != nil {
				json.NewEncoder(w).Encode(false)
			} else {
				token, err := utils.CreateToken(u.Email)
				if err != nil {
					json.NewEncoder(w).Encode(err)
				}
				json.NewEncoder(w).Encode(token)
			}
		}
	}
}
func RegisterData(DB *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var data types.Data
		json.Unmarshal(reqBody, &data)
		fmt.Println(data.Description)
		insert, err := DB.Exec("INSERT INTO data VALUES (?,?,?,?,?)", data.Name, data.Description, data.Code, data.ProducedAt, time.Now())

		if err != nil {
			json.NewEncoder(w).Encode(err)
			panic(err.Error())
		}
		n, err := insert.RowsAffected()
		if n == 1 {
			fmt.Println("Data saved successfully")
		}
		json.NewEncoder(w).Encode(data)
	}
}
func QueryData(DB *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query, err := DB.Query("SELECT * FROM data")
		if err != nil {
			panic(err.Error())
		}
		var arrData []types.Data
		for query.Next() {
			var data types.Data
			err = query.Scan(&data.Name, &data.Description, &data.Code, &data.ProducedAt, &data.CreatedAt)
			if err != nil {
				panic(err.Error())
			}
			arrData = append(arrData, data)
		}
		json.NewEncoder(w).Encode(arrData)
	}
}
