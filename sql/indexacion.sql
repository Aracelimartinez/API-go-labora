CREATE INDEX idx_items_product ON items (product);

\timing
SELECT * FROM items WHERE product = 'Product D';
