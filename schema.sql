-- Create the Cages table
CREATE TABLE IF NOT EXISTS Cages (
    ID INTEGER PRIMARY KEY,
    Status TEXT,
    BreedingBird1ID INTEGER,
    BreedingBird2ID INTEGER,
    Timestamp DATETIME,
    CleaningRequired BOOLEAN,
    EggsCurr INTEGER,
    ChicksCurr INTEGER,
    Notes TEXT
);

-- Create the Birds table
CREATE TABLE IF NOT EXISTS Birds (
    ID INTEGER PRIMARY KEY,
    CageID INTEGER,
    Specie TEXT,
    Gender TEXT,
    YearAcquired INTEGER,
    FatherBirdID INTEGER,
    MotherBirdID INTEGER
);

-- Create the Broods table
CREATE TABLE IF NOT EXISTS Broods (
    ID INTEGER PRIMARY KEY,
    CageID INTEGER,
    NumEggs INTEGER,
    NumFertilized INTEGER,
    Timestamp DATETIME
);

-- Create the Chicks table
CREATE TABLE IF NOT EXISTS Chicks (
    ID INTEGER PRIMARY KEY,
    CageID INTEGER,
    Ringed BOOLEAN,
    Timestamp DATETIME,
    Status TEXT
);

