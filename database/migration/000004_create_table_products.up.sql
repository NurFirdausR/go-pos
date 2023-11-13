CREATE TABLE products (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    price_net FLOAT NOT NULL,
    price_gross FLOAT NOT NULL,
    stock_qty INT NOT NULL,
    description TEXT NOT NULL,
    image TEXT NOT NULL,
    exp_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    company_id INT,
    category_id INT,
    PRIMARY KEY (id),
    FOREIGN KEY (company_id) REFERENCES companies(id),
    FOREIGN KEY (category_id) REFERENCES categories(id)
) ENGINE=InnoDB;
