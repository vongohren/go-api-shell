package models

type ShoppingList struct {
  Id      string `gorethink:"id,omitempty"`
	Items   []Item
  Owner   string `json:"Owner" binding:"required"`
}
