INSERT INTO items (customer_name, order_date, product, quantity, price)
VALUES ('Customer 1', '2023-05-04', 'Product A', 10, 100),
       ('Customer 2', '2023-05-05', 'Product B', 5, 50),
       ('Customer 3', '2023-05-06', 'Product C', 3, 20),
       ('Customer 4', '2023-05-07', 'Product D', 8, 75),
       ('Customer 5', '2023-05-08', 'Product E', 2, 200);

SELECT *
FROM items
WHERE quantity > 3 AND price > 50;
