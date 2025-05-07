CREATE TABLE IF NOT EXISTS public.products (
    id SERIAL PRIMARY KEY,
    product_name VARCHAR(50) NOT NULL,
    price NUMERIC(10,2) NOT NULL
);

CREATE TABLE IF NOT EXISTS public.users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL
);
