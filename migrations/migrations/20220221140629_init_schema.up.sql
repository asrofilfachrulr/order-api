-- Create table of item and its price
CREATE TABLE menu_list (
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    PRIMARY KEY (id)
);

-- Seed menu_list with dummy items and prices
INSERT INTO 
    menu_list (name, price) 
VALUES 
    ('ati ampela', 8000),
    ('tahu telor', 9000),
    ('nasi goreng', 12000),
    ('nasi campur', 8000),
    ('soto ayam', 10000);

-- Create table of list of regeistered order
CREATE TABLE order_list (
    id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    status varchar(10) NOT NULL,
    total INT NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT id_unique 
        UNIQUE (id)
);

-- Create table of list of user's ordered item
CREATE TABLE order_item (
    order_id varchar(255) NOT NULL,
    menu_id INT NOT NULL,
    qty INT NOT NULL,
    CONSTRAINT fk_order_id
        FOREIGN KEY(order_id)
            REFERENCES order_list(id)  
);