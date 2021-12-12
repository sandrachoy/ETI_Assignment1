CREATE database assignment1;
CREATE database assignment1_driver;
CREATE database assignment1_trip;

USE assignment1;

CREATE TABLE Passengers (PassengerID INT auto_increment NOT NULL PRIMARY KEY, FirstName VARCHAR(30) NOT NULL, LastName VARCHAR(30) NOT NULL, MobileNo INT NOT NULL, Email VARCHAR(30) NOT NULL); 

USE assignment1_driver;

CREATE TABLE Drivers (DriverID INT auto_increment NOT NULL PRIMARY KEY, FirstName VARCHAR(30) NOT NULL, LastName VARCHAR(30) NOT NULL, MobileNo INT NOT NULL, Email VARCHAR(30) NOT NULL, IdentificationNo VARCHAR(30) NOT NULL, CarLicenseNo VARCHAR(30) NOT NULL); 

use assignment1_trip;

CREATE TABLE Trips (TripID INT auto_increment NOT NULL PRIMARY KEY, Pickup VARCHAR(6) NOT NULL, Dropoff VARCHAR(6) NOT NULL, Assigned VARCHAR(30) NOT NULL); 

