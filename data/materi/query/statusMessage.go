package query

type StatusMessage struct {
	ProcessID string `json:"processId"`
	Stage     string `json:"stage"`   // e.g., "Process 1", "Process 2"
	Title     string `json:"title"`   // e.g., Judul materi/ebook
	Message   string `json:"message"` // e.g., "Started", "Completed"
	Status    string `json:"status"`  // e.g., "start", "done", "completed", "error"
}
