CREATE TABLE character (
  idcharacter int NOT NULL,
  nombre varchar(45) DEFAULT NULL,
  estado varchar(45) DEFAULT NULL,
  especie varchar(45) DEFAULT NULL,
  PRIMARY KEY (idcharacter)
);

CREATE TABLE location (
  idlocation int NOT NULL,
  nombre varchar(45) DEFAULT NULL,
  tipo varchar(45) DEFAULT NULL,
  dimension varchar(45) DEFAULT NULL,
  PRIMARY KEY (idlocation)
);