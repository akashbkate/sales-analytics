CREATE TABLE customers (
    id VARCHAR PRIMARY KEY,
    name TEXT,
    email TEXT,
    address TEXT
);

CREATE TABLE products (
    id VARCHAR PRIMARY KEY,
    name TEXT,
    category TEXT
);

CREATE TABLE orders (
    id VARCHAR PRIMARY KEY,
    product_id VARCHAR,
    customer_id VARCHAR,
    region TEXT,
    date_of_sale DATE,
    quantity_sold INT,
    unit_price NUMERIC,
    discount NUMERIC,
    shipping_cost NUMERIC,
    payment_method TEXT,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);
