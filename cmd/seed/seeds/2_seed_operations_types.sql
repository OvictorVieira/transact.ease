INSERT INTO transact_ease.operations_types (description) VALUES ('CASH PURCHASE')
ON CONFLICT (description) DO UPDATE
    SET description = EXCLUDED.description;

INSERT INTO transact_ease.operations_types (description) VALUES ('INSTALLMENT PURCHASE')
ON CONFLICT (description) DO UPDATE
    SET description = EXCLUDED.description;

INSERT INTO transact_ease.operations_types (description) VALUES ('WITHDRAWAL')
ON CONFLICT (description) DO UPDATE
    SET description = EXCLUDED.description;

INSERT INTO transact_ease.operations_types (description) VALUES ('PAYMENT')
ON CONFLICT (description) DO UPDATE
    SET description = EXCLUDED.description;
