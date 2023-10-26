CREATE INDEX IF NOT EXISTS idx_transactions_account_operation ON transact_ease.transactions (account_id, operation_type_id);
