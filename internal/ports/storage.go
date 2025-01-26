package ports

type Storage interface {
	Add(text string) (string, error)
	Remove(id int) (string, error)
	Update(id int, new string) (string, error)
	List(filter string) (string, error)
	UpdateStatus(id int, status string) (string, error)
}
