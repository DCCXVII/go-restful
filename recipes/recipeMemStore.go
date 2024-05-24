package recipes

type MemStore struct {
	list map[string]Recipe
}

// Add implements main.recipeStore.
func (m *MemStore) Add(name string, recipe Recipe) error {
	panic("unimplemented")
}

// Get implements main.recipeStore.
func (m *MemStore) Get(name string) (Recipe, error) {
	panic("unimplemented")
}

// List implements main.recipeStore.
func (m *MemStore) List() (map[string]Recipe, error) {
	panic("unimplemented")
}

// Remove implements main.recipeStore.
func (m *MemStore) Remove(name string) error {
	panic("unimplemented")
}

// Update implements main.recipeStore.
func (m *MemStore) Update(name string, recipe Recipe) error {
	panic("unimplemented")
}

func NewMemStore() *MemStore {
	list := make(map[string]Recipe)
	return &MemStore{
		list,
	}
}
