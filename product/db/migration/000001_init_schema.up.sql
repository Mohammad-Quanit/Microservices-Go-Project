CREATE TABLE products (
    "id" bigserial PRIMARY KEY,
    "name" VARCHAR (255) NOT NULL,
    "description" TEXT,
    "price" DECIMAL(10, 2) NOT NULL,
    "category" VARCHAR (255) NOT NULL,
    "stock_quantity" INT NOT NULL
);

CREATE INDEX ON "products" ("name");