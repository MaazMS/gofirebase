## Step by step to build  Authenticate Rest API in GO with Firebase Authentication
1. Install go   
1. create directory 
1. initialize go mod  
1. go get packages   
   1. go get github.com/gin-gonic/gin
   1. go get github.com/jinzhu/gorm
1. Create REST API server in GO     
1. create main.go inside gofirebase  
    1. create a gin instance with gin.Default().  
    1. Created a database instance with function config.CreateDatabase() and it returns db instance of gorm 
    1. Set the database instance to gin context for all incoming requests using middleware in gin use r.Use()  
    1. Defined routes for finding and creating artists with r.GET() and r.POST()
    1. started the server at port 5000.   
    
1. Inside config package , create a file database.go   
    1. create a gorm database instance for sqlite3.  
    1. Used AutoMigrate to migrate our model Artist.  
1. Inside api package , create a file artist.go    
    1. Create a struct Artist which is the model for our database table artist.  
    1. CreateArtistInput is the struct for input request mapping.  
    1. function FindArtists   
       1. first extracted db from gin context, then used it to find all the artists.  
       1. It provide a response.  
   1. function CreateArtist  
      1. first extracted db from gin context,  then mapped input request to our CreateArtistInput struct.  
      1. It created a new artist with db and returned artist as a response.  
   
1. open up your terminal, go to the project directory and run the project `go run main.go`  
1. Create a Firebase Project
   1. Open up firebase console, and create a project [click here](https://console.firebase.google.com/u/0/)  
   1. click continue.  
   1. select location 
   1. accept policy   
   1. click done   
   
1. Generate Firebase Admin SDK GO   
   1. click setting on left side upper corner.  
   1. click on project setting  
   1. click on Service accounts option  
   1. choose language  
   1. click on Generate new private key   
   **Note**  Your private key gives access to your project's Firebase services. Keep it confidential and never store it in a public repository.    
   **Note**  Store this file securely, because your new key can't be recovered if lost.    
   
1. Configure and Initialize Firebase SDK
   1.  use the firebase admin SDK and firebase auth inside our GO a server, we need to add some packages.  
   1. go get firebase.google.com/go/v4
   1. go get firebase.google.com/go
1. config package create a new file firebase.go    
   1.  first we use absolute filepath to get the path to the serviceAccountKey.json from our current file.  
   1. We pass this credential clientOption to firebase.NewApp() along with other params, which creates the firebase app.  
   1. firebase app, we can then extract the firebase auth with, app.Auth() method.    
   1. main.go `  c.Set("firebaseAuth", firebaseAuth)` set firebase auth to gin context with a middleware to all incoming request.  
1. Validate API requests with Firebase Auth   
   1. Before validating the API requests, let's see how we make an API request with firebase authentication token.  
   1. The API request header needs to have: `Authorization: Bearer {{user_firebase_token}}`   
   1. There are multiple ways you could get user_firebase_token.
      1. Use a firebase package in your front end. Login the user with the package and user email/password.   
         Then you will be provided with the token, which then you set to your API request header.  
      1. Create a login API in your go server, validate the user, and then create tokens with the firebase auth.  
      1. token, err := firebaseAuth.CustomToken(context.Background(), "firebase_UID")