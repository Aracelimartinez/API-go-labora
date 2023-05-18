package services

import (
	"sync"
	"testing"
)

func TestIncrementViewCounter(t *testing.T)  {
	var initialViewCounter int
	var err error
	var wg sync.WaitGroup
	itemId := "5"

	err = Db.QueryRow("SELECT view_counter FROM items WHERE id = $1", itemId).Scan(&initialViewCounter)
	if err != nil {
		t.Errorf("Error al seleccionar el item inicial: %v", err)
	}

	expectedViewCounter := initialViewCounter + 100

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			err := IncrementViewCounter(itemId)

			if err != nil {
				t.Errorf("Error al obtener el item: %v", err)
				return
			}
		}()
	}

	wg.Wait()

	var resultViewCounter int
	err = Db.QueryRow("SELECT view_counter FROM items WHERE id = $1", itemId).Scan(&resultViewCounter)
	if err != nil {
		t.Errorf("Error al seleccionar el item inicial: %v", err)
	}

	if resultViewCounter != expectedViewCounter {
		t.Errorf("El contador de vistas no se ha incrementado correctamente. Esperado: %d, Actual: %d", expectedViewCounter, resultViewCounter)
	}

}
