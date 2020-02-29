package Controller

import (
	"encoding/json"
	"net/http"

	"github.com/dankusuma/learngolang/Data"
	"github.com/dankusuma/learngolang/Models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, req *http.Request) {
	var user Models.User
	var checkuser Models.User
	var ress Models.Response
	_ = json.NewDecoder(req.Body).Decode(&user)
	checkuser = Data.GetUserByEmailOrPhoneNumber(user)
	if CheckPasswordHash(user.Password, checkuser.Password) {
		w.WriteHeader(http.StatusOK)
		b, err := json.Marshal(checkuser)
		if err == nil {
			w.WriteHeader(http.StatusCreated)
			ress.Message = "Sukses"
			ress.Data = string(b)

		} else {
			w.WriteHeader(http.StatusInternalServerError)
			ress.Message = "Gagal"
			ress.Data = ""
		}
	} else {
		w.WriteHeader(http.StatusOK)
		ress.Message = "Username / Password Salah"
		ress.Data = ""
	}
	json.NewEncoder(w).Encode(ress)
}

func CreateCustomer(w http.ResponseWriter, req *http.Request) {
	var ress Models.Response
	var user Models.User
	var id = uuid.New()
	var checkuser Models.User
	_ = json.NewDecoder(req.Body).Decode(&user)
	user.ID = id.String()
	user.Password, _ = HashPassword(user.Password)
	checkuser = Data.GetUserByEmailOrPhoneNumber(user)
	if checkuser.ID == "" {
		Data.CreateUser(user)
		b, err := json.Marshal(user)
		if err == nil {
			w.WriteHeader(http.StatusCreated)
			ress.Message = "Sukses"
			ress.Data = string(b)

		} else {
			w.WriteHeader(http.StatusInternalServerError)
			ress.Message = "Gagal"
			ress.Data = ""
		}

	} else {
		w.WriteHeader(http.StatusOK)
		ress.Message = "Nomor HP / Email Sudah Digunakan"
		ress.Data = ""
	}

	json.NewEncoder(w).Encode(ress)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
