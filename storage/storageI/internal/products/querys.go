package products

const (
	INSERT_QUERY      = "INSERT INTO products (name, type, count, price) VALUES (?,?,?,?)"
	GET_BY_NAME_QUERY = "SELECT id, name, type, count, price FROM products WHERE name = ?;"
	GET_ALL_QUERY     = "SELECT id, name, type, count, price FROM products"
	GET_BY_ID_QUERY   = "SELECT id, name, type, count, price FROM products WHERE id = ?"
	DELETE_QUERY      = "DELETE FROM products WHERE id = ?;"
	//UPDATE_QUERY      = "UPDATE products SET name=?, type=?, count=?, price=?, id_warehouse=? WHERE id=?;"
	//EXISTS_QUERY      = "SELECT id FROM products WHERE id=?;"
)
