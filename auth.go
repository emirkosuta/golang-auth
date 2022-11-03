package main2

// import (
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/joho/godotenv"
// )

// func setUpEnv() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		fmt.Println("Could not load .env file")
// 		os.Exit(1)
// 	}

// 	fmt.Println(os.Getenv("JWT_TOKEN_SECRET"))
// }

// func Home(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "super secret area")
// }

// var SECRET = []byte(os.Getenv("JWT_TOKEN_SECRET"))

// func createToken() (string, error) {
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["exp"] = time.Now().Add(time.Hour).Unix()

// 	tokenStr, err := token.SignedString(SECRET)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return "", err
// 	}

// 	return tokenStr, nil
// }

// func validateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		if r.Header["Bearer"] != nil {
// 			token, err := jwt.Parse(r.Header["Bearer"][0], func(t *jwt.Token) (interface{}, error) {
// 				_, ok := t.Method.(*jwt.SigningMethodHMAC)
// 				if !ok {
// 					w.WriteHeader(http.StatusUnauthorized)
// 					w.Write([]byte("Not authorized"))
// 				}
// 				return SECRET, nil
// 			})

// 			if err != nil {
// 				w.WriteHeader(http.StatusUnauthorized)
// 				w.Write([]byte("Not authorized"))
// 			}

// 			if token.Valid {
// 				next(w, r)
// 			}
// 		} else {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte("Not authorized"))
// 		}
// 	})
// }

// func main() {
// 	setUpEnv()

// 	http.HandleFunc("/api", Home)

// 	http.ListenAndServe(":3500", nil)
// }
