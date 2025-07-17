CREATE EXTENSION IF NOT EXISTS "pgcrypto";
DO $$ BEGIN
    CREATE TYPE alert_type_enum AS ENUM ('low_stock','out_of_stock','overstock');
    EXCEPTION
    WHEN duplicate_object THEN NULL;
end $$;
DO $$ BEGIN
CREATE TYPE status_type_enum AS ENUM ('sent','pending','acknowledged');
EXCEPTION
WHEN duplicate_object  THEN NULL;
end $$;
CREATE TABLE IF NOT EXISTS alerts (
                                      id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                      product_id INTEGER REFERENCES products(id),
                                      alert_message TEXT NOT NULL,
                                      alert_type alert_type_enum NOT  NULL,
    status status_type_enum NOT NULL,
                                      created_at TIMESTAMPTZ DEFAULT NOW()
);
