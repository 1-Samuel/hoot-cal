package owl

type Repository interface {
	Get() ([]Match, error)
}
