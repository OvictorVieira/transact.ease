INSERT INTO transact_ease.accounts (document_number) VALUES ('12345678900')
ON CONFLICT (document_number) DO UPDATE
    SET document_number = EXCLUDED.document_number;

INSERT INTO transact_ease.accounts (document_number) VALUES ('98765432100')
ON CONFLICT (document_number) DO UPDATE
    SET document_number = EXCLUDED.document_number;

INSERT INTO transact_ease.accounts (document_number) VALUES ('11223344556')
ON CONFLICT (document_number) DO UPDATE
    SET document_number = EXCLUDED.document_number;
