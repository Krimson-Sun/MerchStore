-- +goose NO TRANSACTION
-- +goose Up
CREATE TABLE items (
                       id VARCHAR(36) PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       description TEXT,
                       image_url TEXT,
                       price INTEGER NOT NULL CHECK (price > 0),
                       in_stock INTEGER NOT NULL CHECK (in_stock >= 0),
                       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE cart_items (
                            id VARCHAR(36) PRIMARY KEY,
                            user_id VARCHAR(36) NOT NULL,
                            item_id VARCHAR(36) NOT NULL,
                            quantity INTEGER NOT NULL CHECK (quantity > 0),
                            created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
                            FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE,
                            UNIQUE(user_id, item_id)
);

create index concurrently idx_cart_items_user_id on cart_items (user_id);

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS idx_cart_items_user_id;