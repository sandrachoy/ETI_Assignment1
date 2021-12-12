# ETI_Assignment1

<h1>5.1.3.1	Design consideration</h1>
The microservices have a loosely coupled architecture design. The microservices created include Passengers, Drivers, Trips services. Due to them being microservices, each service needs to be started and maintained independently, and have their own initialised go modules. The microservices all use their own created database to hold their data separately as well. Each function has the ability to not only run independently, but also effectively communicate with each other when it is required. There is a single console application to simulate a real ride-sharing platform frontend. Through the console application, the user is able to access all the independent functions from each microservice. 

<h1>5.1.3.2	Architecture diagram</h1>

![Architecture Diagram](https://user-images.githubusercontent.com/64128624/145708437-42c149ee-dd3f-41d7-847f-61ceef9dd8f4.jpg)

<h3>Passengers Service</h3>
The Passengers API allows the user to create a new passenger account in the application, and edit their information after account creation. The user is unable to delete their passengers account.

<h3>Drivers Service</h3>
The Drivers API allows the user to create a new driver account in the application, and edit their information after account creation. The user is unable to delete their drivers account.

<h3>Trips Service</h3>
The Trips API allows the passenger to request a trip and view all past trips, and the driver to initiate start trips and end trips. The driver is able to view all requested trips and select one to start.

<h3>Console</h3>
The Console is the frontend of the entire application which allows the user to input and call the services and functions implemented inside.


<h1>5.1.3.3	Instructions for setting up and running the microservices</h1>

Clone the Repository

Create a new Connection Assignment1 in MySQL Workbench (Hostname:localhost, Port: 3306)

Under the assignment folder, there is a file called CreateAssignment1.sql

To create the databases required to run the microservices, run CreateAssignment1.sql in MYSQL Workbench

Navigate to each of the microservice folders (Passenger, Driver, Trip) and enter ```go run main.go``` in new terminal windows to run them

Ensure that "Listening at port 5000" is shown in each terminal

Navigate to src/Assignment/Console in a new terminal and enter ```go run main.go``` to run the console



