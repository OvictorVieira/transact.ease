CREATE TABLE IF NOT EXISTS transact_ease.operations_types (
    operation_type_id SERIAL PRIMARY KEY,
    description VARCHAR(255) NOT NULL UNIQUE
);
