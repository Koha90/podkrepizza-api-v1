package sqlite

import (
	"fmt"
)

// TODO: Вынести типы
type Pizza struct {
	ID          int
	Title       string
	Description string
	Price       float64
}

func (s *Storage) GetAllPizzas() ([]Pizza, error) {
	const op = "storage.sqlite.GetAllPizzas"

	rows, err := s.db.Query("SELECT * FROM Pizzas")
	if err != nil {
		return nil, fmt.Errorf("%s: query statement: %w", op, err)
	}
	defer rows.Close()

	var pizzas []Pizza

	for rows.Next() {
		var pizza Pizza
		if err := rows.Scan(&pizza.ID, &pizza.Title, &pizza.Description, &pizza.Price); err != nil {
			return nil, fmt.Errorf("%s: scan row: %w", op, err)
		}
		pizzas = append(pizzas, pizza)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: rows error: %w", op, err)
	}

	return pizzas, nil
}

// TODO: Вынести в отдельную функцию для администратора.
func (s *Storage) AddPizzas() error {
	const op = "sqlite.AddPizzas"

	pizzas := []struct {
		title       string
		description string
		ingredients []string
	}{
		{
			"Грибная",
			`Грибная пицца - это нежное сочетание куриного филе, грибов белых, лука, сливочного соуса и сыра моцарелла. Эта пицца представляет собой настоящий кулинарный шедевр, который обогащает ваши вкусовые рецепторы и приносит непередаваемое удовольствие.`,
			[]string{"Куриное филе", "грибы белые", "лук", "сливочный соус", "сыр моцарелла"},
		},
		{
			"Филадельфия",
			`Насыщенный вкус лосося и снежного краба в сочетании с нежным сливочным сыром и сыром моцарелла. Роскошный вкус Филадельфии, который покоряет сердца гурманов.`,
			[]string{"Лосось", "снежный краб", "сливочный сыр", "сливочный сыр", "сыр моцарелла"},
		},
	}

	for _, pizza := range pizzas {
		_, err := s.db.Exec(
			"INSERT INTO Pizzas (title, description, price) VALUES (?, ?, ?)",
			pizza.title,
			pizza.description,
			500,
		)
		if err != nil {
			return fmt.Errorf("%s: error insert pizzas: %w", op, err)
		}

		pizzaID, err := s.GetPizzaIDByTitle(pizza.title)
		if err != nil {
			return fmt.Errorf("%s: error getting pizza ID: %w", op, err)
		}

		for _, ingredient := range pizza.ingredients {
			var ingredientID int
			_, err := s.db.Exec("INSERT INTO Ingredients (title) VALUES (?)", ingredient)
			if err != nil {
				return fmt.Errorf("%s: error insert ingredients: %w", op, err)
			}

			ingredientID, err = s.GetIngredientIDByTitle(ingredient)
			if err != nil {
				return fmt.Errorf("%s: error getting ingredient ID: %w", op, err)
			}

			_, err = s.db.Exec(
				"INSERT INTO PizzaIngredients (pizza_id, ingredient_id) VALUES (?, ?)",
				pizzaID,
				ingredientID,
			)
			if err != nil {
				return fmt.Errorf("%s: error foreign pizza and ingredients: %w", op, err)
			}
		}
	}

	return nil
}

func (s *Storage) GetPizzaIDByTitle(title string) (int, error) {
	const op = "sqlite.GetPizzaIDByTitle"
	var pizzaID int

	err := s.db.QueryRow("SELECT COALESCE(id, 0) FROM Pizzas WHERE title = ?", title).Scan(&pizzaID)
	if err != nil {
		return 0, fmt.Errorf("%s: error getting pizza ID: %w", op, err)
	}

	return pizzaID, nil
}

func (s *Storage) GetIngredientIDByTitle(title string) (int, error) {
	const op = "sqlite.GetIngredientIDByTitle"
	var ingredientID int

	err := s.db.QueryRow("SELECT COALESCE(id, 0) FROM Ingredients WHERE title = ?", title).
		Scan(&ingredientID)
	if err != nil {
		return 0, fmt.Errorf("%s: error getting ingredient ID: %w", op, err)
	}

	return ingredientID, nil
}
