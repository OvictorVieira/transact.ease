CREATE TABLE IF NOT EXISTS transact_ease.accounts (
    account_id SERIAL PRIMARY KEY,
    document_number VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
