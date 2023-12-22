package acc

import "testing"

func TestInMemoryStorage(t *testing.T) {
	storage := NewInMemoryStorage[int]()
	data := []Data[int]{
		{Original: 1, ID: "id1"},
		{Original: 2, ID: "id2"},
	}

	// Veri kaydetme ve yükleme işlemleri
	if err := storage.Save(data); err != nil {
		t.Errorf("Error saving data to in-memory storage: %v", err)
	}

	loadedData, err := storage.Load()
	if err != nil {
		t.Errorf("Error loading data from in-memory storage: %v", err)
	}

	if len(loadedData) != len(data) {
		t.Errorf("Expected %d items, but got %d", len(data), len(loadedData))
	}
}
