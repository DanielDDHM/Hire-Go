CREATE TABLE Models (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES Users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    age INT,
    height DECIMAL(4,2),
    weight DECIMAL(4,2),
    bio TEXT,
    photos TEXT[],
    address: VARCHAR(50),
    city: VARCHAR(25),
    country: VARCHAR(25),
    created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Bookers (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES Users(id) ON DELETE CASCADE,
    name VARCHAR(255),
    phone VARCHAR(20),
    address: VARCHAR(50),
    city: VARCHAR(25),
    country: VARCHAR(25),
    created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Reviews (
    id SERIAL PRIMARY KEY,
    reviewer_id INT REFERENCES Users(id) ON DELETE SET NULL,
    reviewed_id INT REFERENCES Users(id) ON DELETE SET NULL,
    rating DECIMAL(2,1) CHECK (rating >= 1 AND rating <= 5),
    comment TEXT,
    created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
);