package oauth

import (
	"net/http"
	"net/url"
	"strings"
	"golang.org/x/oauth2"
)

func HandleLogin(w http.ResponseWriter, r *http.Request, oauthConf *oauth2.Config, oauthState string) {
	URL, err := url.Parse(oauthConf.Endpoint.AuthURL)
	if err != nil {
		log.Fatal("Parse: ", err)
	}
	parameters := url.Values{}
	parameters.Add("client_id", oauthConf.ClientID)
	parameters.Add("redirect_uri", oauthConf.RedirectURL)
	parameters.Add("response_type", "code")
	// parameters.Add("scope", "openid email")
	parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
	parameters.Add("state", oauthState)

	URL.RawQuery = parameters.Encode()
	url := URL.String()
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}
