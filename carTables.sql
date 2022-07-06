CREATE TABLE engineLocation(
engineLocationID INT NOT NULL PRIMARY KEY,
engineLocationName VARCHAR(100) NOT NULL
);

CREATE TABLE driveWheel(
driveWheelID INT NOT NULL PRIMARY KEY,
driveWheelName VARCHAR(100) NOT NULL
);

CREATE TABLE fuelType(
fuelTypeID INT NOT NULL PRIMARY KEY,
fuelTypeName VARCHAR(100) NOT NULL
);

CREATE TABLE fuelSystem(
fuelSystemID INT NOT NULL PRIMARY KEY,
fuelSystemName VARCHAR(100) NOT NULL
);

CREATE TABLE engineType(
engineTypeID INT NOT NULL PRIMARY KEY,
engineTypeName VARCHAR(100) NOT NULL
);

CREATE TABLE aspiration(
aspirationID INT NOT NULL PRIMARY KEY,
aspirationType VARCHAR(100) NOT NULL
);

CREATE TABLE carBrand(
carBrandID INT NOT NULL PRIMARY KEY,
carBrandName VARCHAR(100) NOT NULL
);

CREATE TABLE carBody(
carBodyID INT NOT NULL PRIMARY KEY,
carBodyName VARCHAR(100) NOT NULL
);

CREATE TABLE engine(
engineID INT NOT NULL PRIMARY KEY,
aspirationID INT,
fuelTypeID INT,
fuelSystemID INT ,
engineTypeID INT,
cylinderNum INT NOT NULL,
engineSize DECIMAL(4, 2) NOT NULL, 
boreRatio DECIMAL(4, 2) NOT NULL,
stroke DECIMAL(4, 2) NOT NULL,
compressionRatio DECIMAL(4, 2) NOT NULL,
horsePower INT NOT NULL,
peakRPM INT NOT NULL,
FOREIGN KEY(aspirationID) references aspiration(aspirationID),
FOREIGN KEY(fueltypeID) REFERENCES fuelType(fuelTypeID),
FOREIGN KEY (fuelSystemID) REFERENCES fuelSystem(fuelSystemID),
FOREIGN KEY (engineTypeID) REFERENCES engineType(engineTypeID)
);

CREATE TABLE car(
carID INT NOT NULL  PRIMARY KEY,
carBrandID INT,
carBodyID INT,
driveWheelID INT,
engineID INT,
engineLocationID INT,
carName VARCHAR(100) NOT NULL,
price MEDIUMBLOB NOT NULL,
wheelBase DECIMAL(8, 2) NOT NULL,
carLength DECIMAL(8, 2) NOT NULL,
carWidth DECIMAL(8, 2) NOT NULL,
carHeight DECIMAL(8, 2) NOT NULL,
curbWeight DECIMAL(10, 2) NOT NULL,
doorNumber INT NOT NULL,
cityMPG INT NOT NULL,
highwayMPG INT NOT NULL,
FOREIGN KEY (carBrandID) REFERENCES carBrand(carBrandID),
FOREIGN KEY (carBodyID) REFERENCES carBody(carBodyID),
FOREIGN KEY (driveWheelID) REFERENCES driveWheel(driveWheelID),
FOREIGN KEY (engineID) REFERENCES engine(engineID),
FOREIGN KEY (engineLocationID) REFERENCES engineLocation(engineLocationID)
);

ALTER TABLE car MODIFY COLUMN carID INT auto_increment
