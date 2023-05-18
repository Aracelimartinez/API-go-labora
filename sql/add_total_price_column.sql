ALTER TABLE items ADD COLUMN total_price NUMERIC(10, 2);

UPDATE items SET total_price = ROUND((price * quantity)::numeric, 2);
