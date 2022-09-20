CREATE TABLE `users` (
    id INT auto_increment,
    name VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT NOW(),

    PRIMARY KEY (`id`)
);
