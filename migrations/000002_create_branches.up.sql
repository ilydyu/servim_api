create table if not exists branches (
    id serial primary key,
    title varchar(255) not null,
    address varchar(500) not null,
    company_id integer references companies(id) on delete cascade,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON branches
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();