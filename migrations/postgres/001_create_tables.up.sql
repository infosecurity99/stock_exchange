CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



CREATE TABLE stocks (
    stock_id SERIAL PRIMARY KEY,
    symbol VARCHAR(10) UNIQUE NOT NULL,
    company_name VARCHAR(100) NOT NULL,
    current_price NUMERIC(15, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE market_updates (
    update_id SERIAL PRIMARY KEY,
    stock_id INT REFERENCES stocks(stock_id),
    price NUMERIC(15, 2) NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE orders (
    order_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(user_id),
    stock_id INT REFERENCES stocks(stock_id),
    order_type VARCHAR(10) CHECK (order_type IN ('buy', 'sell')),
    quantity INT NOT NULL,
    price NUMERIC(15, 2) NOT NULL,
    status VARCHAR(10) CHECK (status IN ('pending', 'completed', 'cancelled')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE order_book (
    order_id INT REFERENCES orders(order_id),
    order_type VARCHAR(10) CHECK (order_type IN ('buy', 'sell')),
    quantity INT NOT NULL,
    price NUMERIC(15, 2) NOT NULL
);



CREATE TABLE portfolios (
    portfolio_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(user_id),
    stock_id INT REFERENCES stocks(stock_id),
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE trade_confirmations (
    confirmation_id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(order_id),
    user_id INT REFERENCES users(user_id),
    stock_id INT REFERENCES stocks(stock_id),
    quantity INT NOT NULL,
    price NUMERIC(15, 2) NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE market_news (
    news_id SERIAL PRIMARY KEY,
    headline TEXT NOT NULL,
    content TEXT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE account_information (
    account_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(user_id),
    balance NUMERIC(15, 2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE fund_transfers (
    transfer_id SERIAL PRIMARY KEY,
    from_user_id INT REFERENCES users(user_id),
    to_user_id INT REFERENCES users(user_id),
    amount NUMERIC(15, 2) NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
