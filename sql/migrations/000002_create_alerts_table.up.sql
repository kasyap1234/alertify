CREATE TABLE alerts (
                        id UUID PRIMARY KEY,
                        product_id INTEGER REFERENCES products(id) ON DELETE CASCADE,
                        created_at TIMESTAMPTZ DEFAULT NOW()
);
