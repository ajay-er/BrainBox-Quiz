package models

type AdminLogin struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"min=8,max=20"`
}
type AdminDetailsResponse struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}
type DashBoardUser struct {
	TotalUsers   int
	BlockedUsers int
	Users        []SignupDetailResponse
}
type DashBoardCategory struct {
	TotalCategories int
	Category        []Category
}
type DashBoardQuiz struct {
	TotalQuizes int
}
type DashBoardQuestions struct {
	TotalQuestions int
}
type TotalAdminDashboard struct {
	DashBoardUser DashBoardUser
	DashBoardCategory  DashBoardCategory
	// DashBoardQuiz      DashBoardQuiz
	// DashBoardQuestions DashBoardQuestions
}
