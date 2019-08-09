```sql
CREATE TABLE sodas (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    label character varying(255)
);

ALTER TABLE sodas ADD CONSTRAINT sodas_pkey PRIMARY KEY (id);
```