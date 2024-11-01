CREATE TABLE IF NOT EXISTS categories (
  category TEXT PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS balances (
  date TEXT PRIMARY KEY,
  category TEXT NOT NULL,
  amount_cents INTEGER NOT NULL,
  FOREIGN KEY (category) REFERENCES categories(category)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS transactions (
  id INTEGER PRIMARY KEY,
  date TEXT NOT NULL,
  amount_cents INTEGER NOT NULL,
  description TEXT NOT NULL
)
