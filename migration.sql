CREATE TABLE IF NOT EXISTS users(
    id_user INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    password TEXT NOT NULL,
    ocupation TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS pokemons(
    id_pokemon INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    level INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS user_pokemon(
    id_user INTEGER NOT NULL, 
    id_pokemon INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS attacks(
    id_attack INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    power INTEGER NOT NULL,
    defense INTEGER NOT NULL,
    speed INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS pokemon_attack(
    id_pokemon INTEGER NOT NULL, 
    id_attack INTEGER NOT NULL
);

INSERT INTO users (name, password, ocupation) 
    VALUES
        ('Sand Juarez', 'password', 'Pokemon master'),
        ('Ash Ketchum','ash123', 'Trainer'),
        ('Marco Diaz', 'marco123', 'Caretaker'),
        ('Dante Gomez', 'dante123', 'Traveler');

INSERT INTO pokemons (name, type, level) 
    VALUES
        ('Pikachu', 'Electric', 80),
        ('Mewtwo', 'Legendary', 1000),
        ('Charmander', 'Fire', 70),
        ('Eevee', 'Normal', 15),
        ('Vaporeon', 'Water', 150),
        ('Chikorita', 'Plant', 30),
        ('Noctowl', 'Air', 1);

INSERT INTO user_pokemon (id_user, id_pokemon) 
    VALUES
        (1,3),
        (3,7),
        (1,4),
        (2,1),
        (4,6),
        (2,2),
        (5,5);

INSERT INTO attacks (name, power, defense, speed) 
    VALUES
        ('Blades blow', 50, 80, 20),
        ('Slicer', 100, 80, 5),
        ('Fleeting', 70, 40, 10),
        ('Swift hit', 10, 10, 100);

INSERT INTO pokemon_attack (id_pokemon, id_attack) 
    VALUES
        (1,3),
        (3,7),
        (1,4),
        (2,1),
        (4,6),
        (2,2);
        