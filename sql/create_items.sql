CREATE TABLE IF NOT EXISTS items (
    id INTEGER PRIMARY KEY,
    customer_name VARCHAR NOT NULL,
    order_date DATE NOT NULL,
    product VARCHAR NOT NULL,
    quantity INTEGER NOT NULL,
    price NUMERIC NOT NULL
);
