package handlers

import (
  "log"
  "fmt"
  "os"
  "errors"
  "net/http"
  "time"
  "encoding/json"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
  "github.com/Snorlock/shoppingApi/db"
  "github.com/Snorlock/mux"
  re "github.com/dancannon/gorethink"
  jwt "github.com/dgrijalva/jwt-go"
)

func init () {
  goth.UseProviders(
    facebook.New(os.Getenv("FACEBOOK_SECRET"), os.Getenv("FACEBOOK_APP_SECRET"), fmt.Sprintf("http://%s/auth/facebook/callback", os.Getenv("DNS_HOSTNAME"))),
  )
  gothic.GetProviderName = getProviderName
}

func BeginAuthHandler(env *db.Env,  w http.ResponseWriter, r *http.Request) error {
  gothic.BeginAuthHandler(w, r)
  return nil;
}

func CallBack(env *db.Env, w http.ResponseWriter, r *http.Request) error {
	// print our state string to the console. Ideally, you should verify
	// that it's the same string as the one you set in `setState`
	log.Printf("State: ", gothic.GetState(r))
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return err
	}
  users, err2 := re.DB("list_api").Table("users").Filter(map[string]interface{}{
    "Email":user.Email,
  }).Run(env.DBSession)
  if err2 != nil {
    return err2
  }
  var usersArr []interface{}
  err3 := users.All(&usersArr)
  if err3 != nil {
      return err3
  }
  if len(usersArr) > 0 {
    log.Printf("User exists")
  } else {
    log.Printf("User does not exist, creating")
    _, err4 := re.DB("list_api").Table("users").Insert(user).RunWrite(env.DBSession)
    if err4 != nil {
        return err4
    }
  }

  token := jwt.New(jwt.SigningMethodHS256)
  token.Claims["id"] = user.UserID
  token.Claims["iat"] = time.Now().Unix()
  token.Claims["exp"] = time.Now().Add(time.Second * 3600 * 24).Unix()
  jwtString, err5 := token.SignedString([]byte("mysupersecretkey"))
  log.Printf(jwtString)
  if err5 != nil {
      return err5
  }
  jsonToken := Token{jwtString}
  js, err := json.Marshal(jsonToken)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return nil
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)

  // http.Redirect(w, r, "/", 301)
  return nil
}

type Token struct {
  Bearer string
}

func getProviderName(req *http.Request) (string, error) {
	provider := req.URL.Query().Get("provider")
	if provider == "" {
		params := mux.Vars(req)
		provider = params["provider"]
	}
	if provider == "" {
		return provider, errors.New("you must select a provider")
	}
	return provider, nil
}
