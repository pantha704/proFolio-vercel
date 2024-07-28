// /api/user.go
package api

import (
	"backend/handlers"   // Importing the handlers package
	"backend/middleware" // Importing the middleware package

	"github.com/gorilla/mux" // Importing the mux package from Gorilla for HTTP routing
)

// Handler is the main entry point for Vercel
// func Handler(w http.ResponseWriter, r *http.Request) {
// 	router := mux.NewRouter()
// 	RegisterUserRoutes(router)
// 	router.ServeHTTP(w, r)
// }

// RegisterUserRoutes registers all the routes related to user operations

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/api/signup", handlers.SignUpHandler).Methods("POST") // Route for user signup
	router.HandleFunc("/api/signin", handlers.SignInHandler).Methods("POST") // Route for user signin
	router.HandleFunc("/api/skills", handlers.GetSkillsHandler).Methods("GET") 
	
	// Routes that require authentication
	authenticated := router.PathPrefix("/api").Subrouter()
	authenticated.Use(middleware.JwtVerify)

	// Add User
	authenticated.HandleFunc("/user", handlers.AddUserHandler).Methods("POST")

	// Get User
	authenticated.HandleFunc("/user", handlers.GetUserHandler).Methods("GET")                               // Route for getting a user
	authenticated.HandleFunc("/user/{id}", handlers.GetUserByIDHandler).Methods("GET")                      // Route for getting a user by ID
	authenticated.HandleFunc("/user/email/{email}", handlers.GetUserByEmailHandler).Methods("GET")          // Route for getting a user by email
	authenticated.HandleFunc("/user/username/{username}", handlers.GetUserByUsernameHandler).Methods("GET") // Route for getting a user by username

	// Update User
	authenticated.HandleFunc("/user/{id}", handlers.UpdateUserHandler).Methods("PATCH")                          // Route for updating a user by ID
	authenticated.HandleFunc("/user/email/{email}", handlers.UpdateUserByEmailHandler).Methods("PATCH")          // Route for updating a user by email
	authenticated.HandleFunc("/user/username/{username}", handlers.UpdateUserByUsernameHandler).Methods("PATCH") // Route for updating a user by username
	authenticated.HandleFunc("/user", handlers.AddUserHandler).Methods("POST")                                   // Route for adding a new user

	// Get User Skills
	authenticated.HandleFunc("/user/id/{id}/skills", handlers.GetSkillsByUserIDHandler).Methods("GET")               // Route for getting skills by id
	authenticated.HandleFunc("/user/username/{username}/skills", handlers.GetSkillsByUsernameHandler).Methods("GET") // Route for getting skills by username
	authenticated.HandleFunc("/user/email/{email}/skills", handlers.GetSkillsByEmailHandler).Methods("GET")          // Route for getting skills by email

	// AI Routes
	authenticated.HandleFunc("/cover-letter", handlers.GeminiCoverLetterHandler).Methods("POST")  // Route for generating a cover letter using Gemini
	authenticated.HandleFunc("/calc-chance", handlers.CalculateReplacementChance).Methods("POST") // Route for calculating replacement chance
	authenticated.HandleFunc("/resume-review", handlers.ResumeReview).Methods("POST")             // Route for reviewing a resume
	
	
	// Resume CRUD
	
	authenticated.HandleFunc("/makeDefault", handlers.SetDefaultResumeHandler).Methods("POST")
	authenticated.HandleFunc("/user/{userID}/resumes", handlers.AddResumeHandler).Methods("POST")
	authenticated.HandleFunc("/user/{userID}/resumes/{resumeID}", handlers.UpdateResumeHandler).Methods("PUT")
	authenticated.HandleFunc("/user/{userID}/resumes/{resumeID}", handlers.DeleteResumeHandler).Methods("DELETE")

	// This is Cover letter handler using Open AI (to be used only when OPENAI key is present)
	authenticated.HandleFunc("/cover-letter", handlers.OpenAICoverLetterHandler).Methods("POST")
}

// // Handler is the main entry point for Vercel
// func Handler(w http.ResponseWriter, r *http.Request) {
// 	router := routes.SetupRouter()
// 	router.ServeHTTP(w, r)
// }

// // Exported route handlers for Vercel

// // SignUp handles user registration
// func SignUp(w http.ResponseWriter, r *http.Request) {
// 	handlers.SignUpHandler(w, r)
// }

// // SignIn handles user authentication
// func SignIn(w http.ResponseWriter, r *http.Request) {
// 	handlers.SignInHandler(w, r)
// }

// // GetSkills retrieves all available skills
// func GetSkills(w http.ResponseWriter, r *http.Request) {
// 	handlers.GetSkillsHandler(w, r)
// }

// // AddUser handles adding a new user
// func AddUser(w http.ResponseWriter, r *http.Request) {
// 	middleware.JwtVerify(http.HandlerFunc(handlers.AddUserHandler)).ServeHTTP(w, r)
// }

// // GetUser retrieves user information
// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	middleware.JwtVerify(http.HandlerFunc(handlers.GetUserHandler)).ServeHTTP(w, r)
// }

// // GetUserByID retrieves a user by their ID
// func GetUserByID(w http.ResponseWriter, r *http.Request) {
// 	middleware.JwtVerify(http.HandlerFunc(handlers.GetUserByIDHandler)).ServeHTTP(w, r)
// }

// // GetUserByEmail retrieves a user by their email
// func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
// 	middleware.JwtVerify(http.HandlerFunc(handlers.GetUserByEmailHandler)).ServeHTTP(w, r)
// }

// // GetUserByUsername retrieves a user by their username
// func GetUserByUsername(w http.ResponseWriter, r *http.Request) {
// 	middleware.JwtVerify(http.HandlerFunc(handlers.GetUserByUsernameHandler)).ServeHTTP(w, r)
// }

// // UpdateUser updates user information
// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	middleware.JwtVerify(http.HandlerFunc(handlers.UpdateUserHandler)).ServeHTTP(w, r)
// }

// // UpdateUserByEmail updates a user by their email
// func UpdateUserByEmail(w http.ResponseWriter, r *http.Request) {
// 	middleware.JwtVerify(http.HandlerFunc(handlers.UpdateUserByEmailHandler)).ServeHTTP(w, r)
// }

// // UpdateUserByUsername updates a user by their username
// func UpdateUserByUsername(w http.ResponseWriter, r *http.Request) {
// 	middleware.JwtVerify(http.HandlerFunc(handlers.UpdateUserByUsernameHandler)).ServeHTTP(w, r)
// }

// // GetSkillsByUserID retrieves skills for a user by their ID
// func GetSkillsByUserID(w http.ResponseWriter, r *http.Request) {
// 	middleware.JwtVerify(http.HandlerFunc(handlers.GetSkillsByUserIDHandler)).ServeHTTP(w, r)
// }

// // GetSkillsByUsername retrieves skills for a user by their username
// func GetSkillsByUsername(w http.ResponseWriter, r *http.Request) {
// 	middleware.JwtVerify(http.HandlerFunc(handlers.GetSkillsByUsernameHandler)).ServeHTTP(w, r)
// }

// // GetSkillsByEmail retrieves skills for a user by their email
// func GetSkillsByEmail(w http.ResponseWriter, r *http.Request) {
// 	middleware.JwtVerify(http.HandlerFunc(handlers.GetSkillsByEmailHandler)).ServeHTTP(w, r)
// }

// // ... Continue with other route handlers ...

// // GeminiCoverLetter generates a cover letter using Gemini
// func GeminiCoverLetter(w http.ResponseWriter, r *http.Request) {
// 	middleware.JwtVerify(http.HandlerFunc(handlers.GeminiCoverLetterHandler)).ServeHTTP(w, r)
// }

// // CalculateReplacementChance calculates the chance of replacement
// func CalculateReplacementChance(w http.ResponseWriter, r *http.Request) {
// 	middleware.JwtVerify(http.HandlerFunc(handlers.CalculateReplacementChance)).ServeHTTP(w, r)
// }

// // ResumeReview handles resume review requests
// func ResumeReview(w http.ResponseWriter, r *http.Request) {
// 	middleware.JwtVerify(http.HandlerFunc(handlers.ResumeReview)).ServeHTTP(w, r)
// }

// // SetDefaultResume sets a resume as default
// func SetDefaultResume(w http.ResponseWriter, r *http.Request) {
// 	middleware.JwtVerify(http.HandlerFunc(handlers.SetDefaultResumeHandler)).ServeHTTP(w, r)
// }

// // AddResume adds a new resume
// func AddResume(w http.ResponseWriter, r *http.Request) {
// 	middleware.JwtVerify(http.HandlerFunc(handlers.AddResumeHandler)).ServeHTTP(w, r)
// }

// // UpdateResume updates an existing resume
// func UpdateResume(w http.ResponseWriter, r *http.Request) {
// 	middleware.JwtVerify(http.HandlerFunc(handlers.UpdateResumeHandler)).ServeHTTP(w, r)
// }

// // DeleteResume deletes a resume
// func DeleteResume(w http.ResponseWriter, r *http.Request) {
// 	middleware.JwtVerify(http.HandlerFunc(handlers.DeleteResumeHandler)).ServeHTTP(w, r)
// }

// // OpenAICoverLetter generates a cover letter using OpenAI
// func OpenAICoverLetter(w http.ResponseWriter, r *http.Request) {
// 	middleware.JwtVerify(http.HandlerFunc(handlers.OpenAICoverLetterHandler)).ServeHTTP(w, r)
// }