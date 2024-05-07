CREATE TABLE users (
    _id text PRIMARY KEY,
    name text NOT NULL,
    status text NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL
);

CREATE TABLE user_bank_cards (
    _id text PRIMARY KEY,
    number text NOT NULL,
    user_id text NOT NULL,
    is_default BOOLEAN DEFAULT false,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    CONSTRAINT fk_user_bank_cards_users FOREIGN KEY (user_id) REFERENCES users(_id)
);

CREATE TABLE shops (
    _id text PRIMARY KEY,
    name text NOT NULL,
    status text NOT NULL,
    user_id text NOT NULL,
    cash decimal NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    CONSTRAINT fk_shops_users FOREIGN KEY (user_id) REFERENCES users(_id)
);

CREATE TABLE shop_bank_cards (
    _id text PRIMARY KEY,
    number text NOT NULL,
    shop_id text NOT NULL,
    is_default BOOLEAN DEFAULT false,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    CONSTRAINT fk_shop_bank_cards_shops FOREIGN KEY (shop_id) REFERENCES shops(_id)
);

CREATE TABLE carts (
    _id text PRIMARY KEY,
    user_id text NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    CONSTRAINT fk_carts_users FOREIGN KEY (user_id) REFERENCES users(_id)
);

CREATE TABLE locations (
    _id text PRIMARY KEY,
    name text NOT NULL,
    address text NOT NULL,
    status text NOT NULL,
    is_default BOOLEAN DEFAULT false,
    user_id text NOT NULL,
    latitude decimal NOT NULL,
    longitude decimal NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    CONSTRAINT fk_locations_users FOREIGN KEY (user_id) REFERENCES users(_id)
);

CREATE TABLE warehouses (
    _id text PRIMARY KEY,
    name text NOT NULL,
    address text NOT NULL,
    status text NOT NULL,
    user_id text NOT NULL,
    cash decimal NOT NULL,
    latitude decimal NOT NULL,
    longitude decimal NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    CONSTRAINT fk_shops_users FOREIGN KEY (user_id) REFERENCES users(_id)
);

CREATE TABLE warehouse_bank_cards (
    _id text PRIMARY KEY,
    number text NOT NULL,
    warehouse_id text NOT NULL,
    is_default BOOLEAN DEFAULT false,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    CONSTRAINT fk_warehouse_bank_cards_warehouses FOREIGN KEY (warehouse_id) REFERENCES warehouses(_id)
);

CREATE TABLE products (
    _id text PRIMARY KEY,
    name text NOT NULL,
    quantity INT NOT NULL,
    price decimal NOT NULL,
    status text NOT NULL,
    shop_id text NOT NULL,
    warehouse_id text NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    CONSTRAINT fk_products_shops FOREIGN KEY (shop_id) REFERENCES shops(_id),
    CONSTRAINT fk_products_warehouses FOREIGN KEY (warehouse_id) REFERENCES warehouses(_id)
);

CREATE TABLE stock_keeping_units (
    _id text PRIMARY KEY,
    name text NOT NULL,
    sku text NOT NULL,
    quantity INT NOT NULL,
    price decimal NOT NULL,
    status text NOT NULL,
    product_id text NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    CONSTRAINT fk_stock_keeping_units_products FOREIGN KEY (product_id) REFERENCES products(_id)
);

CREATE TABLE cart_items (
    shop_id text NOT NULL,
    cart_id text NOT NULL,
    product_id text NOT NULL,
    sku_id text NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    PRIMARY KEY (shop_id, cart_id, product_id, sku_id),
    CONSTRAINT fk_cart_items_products FOREIGN KEY (product_id) REFERENCES products(_id),
    CONSTRAINT fk_cart_items_carts FOREIGN KEY (cart_id) REFERENCES carts(_id),
    CONSTRAINT fk_cart_items_shops FOREIGN KEY (shop_id) REFERENCES shops(_id),
    CONSTRAINT fk_cart_items_stock_keeping_units FOREIGN KEY (sku_id) REFERENCES stock_keeping_units(_id)
);

CREATE TABLE payments (
    _id text PRIMARY KEY,
    code text NOT NULL,
    method text NOT NULL,
    amount decimal NOT NULL,
    status text NOT NULL,
    user_id text NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    CONSTRAINT fk_payments_users FOREIGN KEY (user_id) REFERENCES users(_id)
);

CREATE TABLE deliveries (
    _id text PRIMARY KEY,
    code text NOT NULL,
    title text NOT NULL,
    ship_fee decimal NOT NULL,
    status text NOT NULL,
    user_id text NOT NULL,
    estimate_date timestamp with time zone NOT NULL,
    user_info jsonb,
    warehouse_id text NOT NULL,
    warehouse_info jsonb,
    location_id text NOT NULL,
    location_info jsonb,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    CONSTRAINT fk_deliveries_users FOREIGN KEY (user_id) REFERENCES users(_id),
    CONSTRAINT fk_deliveries_warehouses FOREIGN KEY (warehouse_id) REFERENCES warehouses(_id),
    CONSTRAINT fk_deliveries_locations FOREIGN KEY (location_id) REFERENCES locations(_id)
);

CREATE TABLE orders (
    _id text PRIMARY KEY,
    code text NOT NULL,
    amount decimal NOT NULL,
    note text NOT NULL,
    priority text NOT NULL,
    status text NOT NULL,
    user_id text NOT NULL,
    shop_id text NOT NULL,
    warehouse_id text NOT NULL,
    destination_warehouse_id text NOT NULL,
    delivery_id text NOT NULL,
    payment_id text NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    CONSTRAINT fk_orders_users FOREIGN KEY (user_id) REFERENCES users(_id),
    CONSTRAINT fk_orders_warehouses FOREIGN KEY (warehouse_id) REFERENCES warehouses(_id),
    CONSTRAINT fk_orders_destination_warehouses FOREIGN KEY (destination_warehouse_id) REFERENCES warehouses(_id),
    CONSTRAINT fk_orders_shops FOREIGN KEY (shop_id) REFERENCES shops(_id),
    CONSTRAINT fk_orders_deliveries FOREIGN KEY (delivery_id) REFERENCES deliveries(_id),
    CONSTRAINT fk_orders_payments FOREIGN KEY (payment_id) REFERENCES payments(_id)
);

CREATE TABLE order_items (
    order_id text NOT NULL,
    product_id text NOT NULL,
    sku_id text NOT NULL,
    name text NOT NULL,
    price decimal NOT NULL,
    quantity int NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    PRIMARY KEY (order_id, product_id, sku_id),
    CONSTRAINT fk_order_items_orders FOREIGN KEY (order_id) REFERENCES orders(_id),
    CONSTRAINT fk_order_items_products FOREIGN KEY (product_id) REFERENCES products(_id),
    CONSTRAINT fk_order_items_stock_keeping_units FOREIGN KEY (sku_id) REFERENCES stock_keeping_units(_id)
);

CREATE TABLE order_sessions (
    _id text PRIMARY KEY,
    order_items jsonb,
    user_id text,
    delivery_session_id text,
    amount decimal NOT NULL,
    created_at timestamp with time zone NOT NULL
);

CREATE TABLE delivery_sessions (
    _id text PRIMARY KEY,
    code text NOT NULL,
    title text NOT NULL,
    ship_fee decimal NOT NULL,
    status text NOT NULL,
    user_id text NOT NULL,
    estimate_date timestamp with time zone NOT NULL,
    warehouse_id text NOT NULL,
    location_id text NOT NULL,
    created_at timestamp with time zone NOT NULL
);
