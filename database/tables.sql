CREATE TABLE IF NOT EXISTS "customers" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "customer_id" TEXT NOT NULL,
    "first_name" TEXT NOT NULL,
    "last_name" TEXT NOT NULL,
    "email" TEXT NOT NULL
);