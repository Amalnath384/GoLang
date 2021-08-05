package student

import (
	"awesomeProject/Practice28-Mockery-Rest-Mongo/pkg/model"
	"awesomeProject/Practice28-Mockery-Rest-Mongo/pkg/repository"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type StudentService struct {
}

func NewStudentService() *StudentService {
	return &StudentService{}
}

var jwtKey = []byte("secret_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func (p *StudentService) ListStudent(w http.ResponseWriter, r *http.Request) {
	//cookie, err := r.Cookie("token")
	//if err != nil {
	//	if err == http.ErrNoCookie {
	//		w.WriteHeader(http.StatusUnauthorized)
	//		fmt.Fprintln(w,"no cookie found")
	//		return
	//	}
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//res := TokenValidation(cookie.Value)
	//fmt.Println(res)

	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	res, err := repository.Repo.ListStudent(ctx)
	if err != nil {
		http.Error(w, "Failed to create in database", http.StatusBadRequest)
		return
	}
	//var Student model.StudentDetails
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

func (p *StudentService) GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	params := mux.Vars(r)
	id := params["id"]
	res, err := repository.Repo.GetStudent(ctx, id)
	if err != nil {
		http.Error(w, "Failed to create in database", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (p *StudentService) CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var Student model.StudentDetails
	err := Student.DecodeFromJSON(r.Body)
	if err != nil {
		http.Error(w, "Failed to Decode", http.StatusBadRequest)
		return
	}
	if err = Student.Validate(); err != nil {
		http.Error(w, "failed to validate struct", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	//deadline := time.Now().Add(time.Duration(50) * time.Second)
	//ctx, cancel := context.WithDeadline(ctx, deadline)
	//defer cancel()

	//ctx := context.Background()
	//ctx1,cancel := context.WithTimeout(ctx,time.Second * 5)
	//defer cancel()

	res, err := repository.Repo.CreateStudent(ctx, &Student)
	if err != nil {
		http.Error(w, "Failed to create in database", http.StatusBadRequest)
		return
	}
	fmt.Println(res)

	w.WriteHeader(http.StatusCreated)
	Student.EncodeToJSON(w)
	return

}

func (p *StudentService) DeleteStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	params := mux.Vars(r)
	id := params["id"]
	err := repository.Repo.DeleteStudent(ctx, id)
	if err != nil {
		http.Error(w, "Failed to create in database", http.StatusBadRequest)
		return
	}
	res := model.StudentDetails{

		Msg: "Student details deleted",
	}
	json.NewEncoder(w).Encode(res.Msg)
}

func (p *StudentService) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var Student model.StudentDetails
	w.Header().Set("Content-Type", "application/json")
	err := Student.DecodeFromJSON(r.Body)
	if err != nil {
		http.Error(w, "Failed to Decode", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	params := mux.Vars(r)
	id := params["id"]
	res, err := repository.Repo.GetStudent(ctx, id)
	if err != nil {
		http.Error(w, "Failed to create in database", http.StatusBadRequest)
		return
	}
	Student.Id = res.Id
	res, err = repository.Repo.UpdateStudent(ctx, &Student)
	if err != nil {
		http.Error(w, "Failed to create in database", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (p *StudentService) Login(w http.ResponseWriter, r *http.Request) {
	var credentials model.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[credentials.Username]

	if !ok || expectedPassword != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &model.Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

}

func (p *StudentService) Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &model.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Username)))

}

func (p *StudentService) Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &model.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	expirationTime := time.Now().Add(time.Minute * 5)

	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "refresh_token",
			Value:   tokenString,
			Expires: expirationTime,
		})

}

func TokenValidation(token string) bool {
	claim := &model.Claims{}
	tkn, err := jwt.ParseWithClaims(token, claim,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Println("signature is not valid")
			return false
		}
		log.Println("error while parsing request")
		return false
	}

	if !tkn.Valid {
		log.Println("token is not valid")
		return false
	}

	if tkn.Valid {
		return true
	}
	return true
}
