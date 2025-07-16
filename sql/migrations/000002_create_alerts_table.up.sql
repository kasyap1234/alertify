CREATE TABLE IF NOT EXISTS alerts (
                        id UUID PRIMARY KEY,
                        product_id INTEGER REFERENCES products(id) ,
    alert_message TEXT NOT NULL,
                        created_at TIMESTAMPTZ DEFAULT NOW()
);
