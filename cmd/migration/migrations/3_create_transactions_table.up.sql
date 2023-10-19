CREATE TABLE IF NOT EXISTS transact_ease.transactions (
    transaction_id SERIAL PRIMARY KEY,
    account_id INTEGER REFERENCES transact_ease.accounts(account_id),
    operation_type_id INTEGER REFERENCES transact_ease.operations_types(operation_type_id),
    amount DECIMAL(15, 2) NOT NULL,
    event_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
