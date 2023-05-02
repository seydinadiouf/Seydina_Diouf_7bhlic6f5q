package payload

type StudentToClass struct {
	StudentID       int    `json:"studentId"`
	SchoolClassName string `json:"schoolClassName"`
}
