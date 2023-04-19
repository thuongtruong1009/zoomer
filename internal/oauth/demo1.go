package oauth

// go get firebase.google.com/go
// import (
//     "context"
//     "firebase.google.com/go/auth"
//     "firebase.google.com/go"
//     "github.com/labstack/echo/v4"
//     "google.golang.org/api/option"
// )

// func main() {
//     // Initialize Firebase app
//     ctx := context.Background()
//     opt := option.WithCredentialsFile("path/to/firebase_credentials.json")
// 	config := &firebase.Config{ProjectID: "my-project-id"}
//     app, err := firebase.NewApp(ctx, config, opt)
//     if err != nil {
// 		log.Fatalf("error initializing app: %v\n", err)
//     }

//     // Initialize Firebase Auth client
//     authClient, err := app.Auth(ctx)
//     if err != nil {
//         panic(err)
//     }

//     // Initialize Echo instance
//     e := echo.New()
//     //...
// }

// func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
//     return func(c echo.Context) error {
//         // Extract ID token from HTTP request
//         token := c.Request().Header.Get("Authorization")
//         if token == "" {
//             return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
//         }

//         // Verify ID token using Firebase Auth client
//         token = strings.TrimPrefix(token, "Bearer ")
//         decodedToken, err := authClient.VerifyIDToken(ctx, token)
//         if err != nil {
//             return echo.NewHTTPError(http.StatusUnauthorized, "Invalid ID token")
//         }

//         // Set authenticated user information in context
//         c.Set("userID", decodedToken.UID)
//         c.Set("email", decodedToken.Claims["email"])

//         return next(c)
//     }
// }

// func protectedRouteHandler(c echo.Context) error {
//     userID := c.Get("userID").(string)
//     email := c.Get("email").(string)

//     // Do something with authenticated user information

//     return c.String(http.StatusOK, "Authenticated User: "+email)
// }

// e.GET("/protected", protectedRouteHandler, authMiddleware)

