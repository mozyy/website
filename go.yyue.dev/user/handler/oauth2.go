package handler

import (
	"crypto/rand"
	"crypto/rsa"
	"net/http"
	"time"

	"github.com/ory/fosite"
	"github.com/ory/fosite/compose"
	"github.com/ory/fosite/storage"
)

// This is the example storage that contains:
// * an OAuth2 Client with id "my-client" and secret "foobar" capable of all oauth2 and open id connect grant and response types.
// * a User for the resource owner password credentials grant type with username "peter" and password "secret".
//
// You will most likely replace this with your own logic once you set up a real world application.
var store = storage.NewExampleStore()

// This secret is being used to sign access and refresh tokens as well as
// authorization codes. It must be exactly 32 bytes long.
var secret = []byte("my super secret signing password")

func generateKey() *rsa.PrivateKey {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		panic("unable to create private key")
	}
	return privateKey
}

// check the api docs of compose.Config for further configuration options
var config = &compose.Config{
	AccessTokenLifespan: time.Minute * 30,
	// ...
}

var oauth2Provider = compose.ComposeAllEnabled(config, store, secret, generateKey())

// The authorize endpoint is usually at "https://mydomain.com/oauth2/auth".
func authorizeHandlerFunc(rw http.ResponseWriter, req *http.Request) {
	// This context will be passed to all methods. It doesn't fulfill a real purpose in the standard library but could be used
	// to abort database lookups or similar things.
	ctx := req.Context()

	// Let's create an AuthorizeRequest object!
	// It will analyze the request and extract important information like scopes, response type and others.
	ar, err := oauth2Provider.NewAuthorizeRequest(ctx, req)
	if err != nil {
		oauth2Provider.WriteAuthorizeError(rw, ar, err)
		return
	}

	// Now that the user is authorized, we set up a session. When validating / looking up tokens, we additionally get
	// the session. You can store anything you want in it.

	// The session will be persisted by the store and made available when e.g. validating tokens or handling token endpoint requests.
	// The default OAuth2 and OpenID Connect handlers require the session to implement a few methods. Apart from that, the
	// session struct can be anything you want it to be.
	mySessionData := &fosite.DefaultSession{
		Username: req.Form.Get("username"),
	}

	// It's also wise to check the requested scopes, e.g.:
	// if authorizeRequest.GetScopes().Has("admin") {
	//     http.Error(rw, "you're not allowed to do that", http.StatusForbidden)
	//     return
	// }

	// Now we need to get a response. This is the place where the AuthorizeEndpointHandlers kick in and start processing the request.
	// NewAuthorizeResponse is capable of running multiple response type handlers which in turn enables this library
	// to support open id connect.
	response, err := oauth2Provider.NewAuthorizeResponse(ctx, ar, mySessionData)
	if err != nil {
		oauth2Provider.WriteAuthorizeError(rw, ar, err)
		return
	}

	// Awesome, now we redirect back to the client redirect uri and pass along an authorize code
	oauth2Provider.WriteAuthorizeResponse(rw, ar, response)
}

// The token endpoint is usually at "https://mydomain.com/oauth2/token"
func tokenHandlerFunc(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	mySessionData := new(fosite.DefaultSession)

	// This will create an access request object and iterate through the registered TokenEndpointHandlers to validate the request.
	accessRequest, err := oauth2Provider.NewAccessRequest(ctx, req, mySessionData)
	if err != nil {
		oauth2Provider.WriteAccessError(rw, accessRequest, err)
		return
	}

	if mySessionData.Username == "super-admin-guy" {
		// do something...
	}

	// Next we create a response for the access request. Again, we iterate through the TokenEndpointHandlers
	// and aggregate the result in response.
	response, err := oauth2Provider.NewAccessResponse(ctx, accessRequest)
	if err != nil {
		oauth2Provider.WriteAccessError(rw, accessRequest, err)
		return
	}

	// All done, send the response.
	oauth2Provider.WriteAccessResponse(rw, accessRequest, response)

	// The client has a valid access token now
}
