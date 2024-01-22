

Sleeping Barber Problem in Golang
=================================

Overview
--------

This is a Golang implementation of the Sleeping Barber Problem. The Sleeping Barber Problem is a classic synchronization and concurrency problem, where there is a barbershop with a certain number of chairs for waiting customers and a group of barbers who cut hair. If the waiting room is full, clients leave; otherwise, they take a seat. If the barber is free, a waiting client is served immediately; otherwise, the client waits in the waiting room. If a barber is free and there are currently no waiting clients, the barber can take a nap.
 The problem is to design a system that ensures the correct and efficient handling of customers and barbers while respecting certain constraints.#The barbers serve clients by cutting their hair, and clients may arrive at any time.

 This project is part of a series of short CLI-based projects inspired by classic computer science problems. I have been reading up on computer science theories and thought there is no better way to test my understanding than with mini golang CLI projects.

Problem Description
-------------------

The Sleeping Barber Problem involves the following entities:

-   **Barber Shop:** The main entity representing the barbershop, which has a specified capacity, a duration for each haircut, and a list of barbers and clients.

-   **Barber:** Each barber is represented as a goroutine in this implementation. Barbers check for clients in the waiting room and cut their hair.

-   **Client:** Clients are represented as goroutines that arrive at random intervals. They join the waiting room and get their hair cut by an available barber.

Techniques Used
---------------

-   **Goroutines:** The solution utilizes Goroutines to represent barbers and clients, allowing concurrent execution of tasks.

-   **Channels:** Channels are used for communication between different goroutines. The main communication channels include `ClientsChan` for adding clients to the waiting room and `BarbersDoneChan` for signaling when a barber has finished for the day.

-   **Randomized Arrival:** Clients arrive at random intervals, adding a level of unpredictability to the simulation.

-   **Time-Based Shop Closure:** The barbershop closes after a specified duration (`timeOpen`), and all barbers finish their current tasks before going home.

Code Structure
---------------

### `main.go`

-   The `main` function initializes the barbershop, sets parameters, and starts the simulation.
-   Barbers are added to the shop, and the shop is opened for a specified duration.
-   A Goroutine continuously adds clients to the waiting room at random intervals until the shop closes.

### `barbershop.go`

-   The `BarberShop` struct represents the state of the barbershop and contains methods for managing barbers, clients, and the overall shop.
-   `addBarber` adds a new barber as a Goroutine, handling their actions, such as sleeping, waking up, and cutting hair.
-   `cutHair` simulates the process of a barber cutting a client's hair with a specified duration.
-   `sendBarberHome` signals that a barber is going home, closing their Goroutine.
-   `closeShopForToday` closes the shop, signaling all barbers to go home and waiting for their Goroutines to finish.
-   `addClient` adds a new client to the waiting room if the shop is open and there is space.

How to Run
----------

1.  Ensure you have Golang installed on your machine.

2.  Clone this repository.

3.  Navigate to the project directory in the terminal.

4.  Run the following command:
```
    bash

 go run main.go barbershop.go
```
Example Output
-------------

Below is an example of terminal output after running the program:

```The Sleeping Barber Problem
---------------------------
The shop is open for today
Frank goes to the waiting room to check for clients.
Frank does not have to cut anyone's hair right now, so he has a nap.
Gerard goes to the waiting room to check for clients.
Gerard does not have to cut anyone's hair right now, so he has a nap.
Milton goes to the waiting room to check for clients.
Milton does not have to cut anyone's hair right now, so he has a nap.
Patrick goes to the waiting room to check for clients.
Patrick does not have to cut anyone's hair right now, so he has a nap.
Jim goes to the waiting room to check for clients.
Jim does not have to cut anyone's hair right now, so he has a nap.
*** Client #1 has arrived.
Client #1 takes a seat in the waiting room.
Client #1 wakes up Frank.
Frank is cutting Client #1's hair.
*** Client #2 has arrived.
Client #2 wakes up Gerard.
Gerard is cutting Client #2's hair.
Client #2 takes a seat in the waiting room.
*** Client #3 has arrived.
Client #3 takes a seat in the waiting room.
Client #3 wakes up Milton.
Milton is cutting Client #3's hair.
*** Client #4 has arrived.
Client #4 takes a seat in the waiting room.
Client #4 wakes up Patrick.
Patrick is cutting Client #4's hair.
*** Client #5 has arrived.
Client #5 takes a seat in the waiting room.
Client #5 wakes up Jim.
Jim is cutting Client #5's hair.
*** Client #6 has arrived.
Client #6 takes a seat in the waiting room.
*** Client #7 has arrived.
Client #7 takes a seat in the waiting room.
*** Client #8 has arrived.
Client #8 takes a seat in the waiting room.
*** Client #9 has arrived.
Client #9 takes a seat in the waiting room.
Frank has finished cutting Client #1's hair
Frank is cutting Client #6's hair.
Gerard has finished cutting Client #2's hair
Gerard is cutting Client #7's hair.
*** Client #10 has arrived.
Client #10 takes a seat in the waiting room.
Milton has finished cutting Client #3's hair
Milton is cutting Client #8's hair.
*** Client #11 has arrived.
Patrick has finished cutting Client #4's hair
Client #11 takes a seat in the waiting room.
Patrick is cutting Client #9's hair.
*** Client #12 has arrived.
Client #12 takes a seat in the waiting room.
*** Client #13 has arrived.
Client #13 takes a seat in the waiting room.
Jim has finished cutting Client #5's hair
Jim is cutting Client #10's hair.
*** Client #14 has arrived.
Client #14 takes a seat in the waiting room.
*** Client #15 has arrived.
Client #15 takes a seat in the waiting room.
*** Client #16 has arrived.
Client #16 takes a seat in the waiting room.
*** Client #17 has arrived.
Client #17 takes a seat in the waiting room.
*** Client #18 has arrived.
Client #18 takes a seat in the waiting room.
Frank has finished cutting Client #6's hair
Frank is cutting Client #11's hair.
*** Client #19 has arrived.
Client #19 takes a seat in the waiting room.
*** Client #20 has arrived.
Client #20 takes a seat in the waiting room.
Gerard has finished cutting Client #7's hair
Gerard is cutting Client #12's hair.
Milton has finished cutting Client #8's hair
Milton is cutting Client #13's hair.
*** Client #21 has arrived.
Client #21 takes a seat in the waiting room.
Patrick has finished cutting Client #9's hair
Patrick is cutting Client #14's hair.
*** Client #22 has arrived.
Client #22 takes a seat in the waiting room.
Jim has finished cutting Client #10's hair
*** Client #23 has arrived.
Jim is cutting Client #15's hair.
Client #23 takes a seat in the waiting room.
*** Client #24 has arrived.
Client #24 takes a seat in the waiting room.
*** Client #25 has arrived.
Client #25 takes a seat in the waiting room.
*** Client #26 has arrived.
The waiting room is full, so Client #26 leaves.
*** Client #27 has arrived.
The waiting room is full, so Client #27 leaves.
*** Client #28 has arrived.
The waiting room is full, so Client #28 leaves.
Frank has finished cutting Client #11's hair
Frank is cutting Client #16's hair.
Gerard has finished cutting Client #12's hair
Gerard is cutting Client #17's hair.
*** Client #29 has arrived.
Client #29 takes a seat in the waiting room.
Milton has finished cutting Client #13's hair
Milton is cutting Client #18's hair.
Patrick has finished cutting Client #14's hair
Patrick is cutting Client #19's hair.
*** Client #30 has arrived.
Client #30 takes a seat in the waiting room.
*** Client #31 has arrived.
Client #31 takes a seat in the waiting room.
Jim has finished cutting Client #15's hair
Jim is cutting Client #20's hair.
*** Client #32 has arrived.
Client #32 takes a seat in the waiting room.
*** Client #33 has arrived.
Client #33 takes a seat in the waiting room.
*** Client #34 has arrived.
The waiting room is full, so Client #34 leaves.
*** Client #35 has arrived.
The waiting room is full, so Client #35 leaves.
*** Client #36 has arrived.
The waiting room is full, so Client #36 leaves.
*** Client #37 has arrived.
The waiting room is full, so Client #37 leaves.
Frank has finished cutting Client #16's hair
Frank is cutting Client #21's hair.
Gerard has finished cutting Client #17's hair
Gerard is cutting Client #22's hair.
Milton has finished cutting Client #18's hair
*** Client #38 has arrived.
Milton is cutting Client #23's hair.
Client #38 takes a seat in the waiting room.
*** Client #39 has arrived.
Client #39 takes a seat in the waiting room.
Patrick has finished cutting Client #19's hair
Patrick is cutting Client #24's hair.
*** Client #40 has arrived.
Client #40 takes a seat in the waiting room.
*** Client #41 has arrived.
Client #41 takes a seat in the waiting room.
Jim has finished cutting Client #20's hair
Jim is cutting Client #25's hair.
*** Client #42 has arrived.
Client #42 takes a seat in the waiting room.
*** Client #43 has arrived.
The waiting room is full, so Client #43 leaves.
*** Client #44 has arrived.
The waiting room is full, so Client #44 leaves.
*** Client #45 has arrived.
The waiting room is full, so Client #45 leaves.
*** Client #46 has arrived.
The waiting room is full, so Client #46 leaves.
*** Client #47 has arrived.
The waiting room is full, so Client #47 leaves.
*** Client #48 has arrived.
The waiting room is full, so Client #48 leaves.
Frank has finished cutting Client #21's hair
Frank is cutting Client #29's hair.
Gerard has finished cutting Client #22's hair
Gerard is cutting Client #30's hair.
*** Client #49 has arrived.
Client #49 takes a seat in the waiting room.
Milton has finished cutting Client #23's hair
Milton is cutting Client #31's hair.
*** Client #50 has arrived.
Client #50 takes a seat in the waiting room.
Patrick has finished cutting Client #24's hair
Patrick is cutting Client #32's hair.
*** Client #51 has arrived.
Client #51 takes a seat in the waiting room.
*** Client #52 has arrived.
Client #52 takes a seat in the waiting room.
Jim has finished cutting Client #25's hair
Jim is cutting Client #33's hair.
*** Client #53 has arrived.
Client #53 takes a seat in the waiting room.
*** Client #54 has arrived.
The waiting room is full, so Client #54 leaves.
*** Client #55 has arrived.
The waiting room is full, so Client #55 leaves.
*** Client #56 has arrived.
The waiting room is full, so Client #56 leaves.
*** Client #57 has arrived.
Frank has finished cutting Client #29's hair
Frank is cutting Client #38's hair.
The waiting room is full, so Client #57 leaves.
Gerard has finished cutting Client #30's hair
Gerard is cutting Client #39's hair.
Milton has finished cutting Client #31's hair
Milton is cutting Client #40's hair.
*** Client #58 has arrived.
Client #58 takes a seat in the waiting room.
Patrick has finished cutting Client #32's hair
Patrick is cutting Client #41's hair.
*** Client #59 has arrived.
Client #59 takes a seat in the waiting room.
*** Client #60 has arrived.
Client #60 takes a seat in the waiting room.
*** Client #61 has arrived.
Client #61 takes a seat in the waiting room.
*** Client #62 has arrived.
The waiting room is full, so Client #62 leaves.
Jim has finished cutting Client #33's hair
Jim is cutting Client #42's hair.
*** Client #63 has arrived.
Client #63 takes a seat in the waiting room.
*** Client #64 has arrived.
The waiting room is full, so Client #64 leaves.
*** Client #65 has arrived.
The waiting room is full, so Client #65 leaves.
*** Client #66 has arrived.
The waiting room is full, so Client #66 leaves.
*** Client #67 has arrived.
Frank has finished cutting Client #38's hair
The waiting room is full, so Client #67 leaves.
Frank is cutting Client #49's hair.
Gerard has finished cutting Client #39's hair
Gerard is cutting Client #50's hair.
Milton has finished cutting Client #40's hair
Milton is cutting Client #51's hair.
*** Client #68 has arrived.
Client #68 takes a seat in the waiting room.
*** Client #69 has arrived.
Client #69 takes a seat in the waiting room.
*** Client #70 has arrived.
Client #70 takes a seat in the waiting room.
Patrick has finished cutting Client #41's hair
Patrick is cutting Client #52's hair.
*** Client #71 has arrived.
Client #71 takes a seat in the waiting room.
Jim has finished cutting Client #42's hair
Jim is cutting Client #53's hair.
*** Client #72 has arrived.
Client #72 takes a seat in the waiting room.
*** Client #73 has arrived.
The waiting room is full, so Client #73 leaves.
*** Client #74 has arrived.
The waiting room is full, so Client #74 leaves.
*** Client #75 has arrived.
The waiting room is full, so Client #75 leaves.
*** Client #76 has arrived.
The waiting room is full, so Client #76 leaves.
*** Client #77 has arrived.
Frank has finished cutting Client #49's hair
The waiting room is full, so Client #77 leaves.
Frank is cutting Client #58's hair.
Gerard has finished cutting Client #50's hair
Gerard is cutting Client #59's hair.
*** Client #78 has arrived.
Client #78 takes a seat in the waiting room.
Milton has finished cutting Client #51's hair
Milton is cutting Client #60's hair.
*** Client #79 has arrived.
Client #79 takes a seat in the waiting room.
*** Client #80 has arrived.
Client #80 takes a seat in the waiting room.
*** Client #81 has arrived.
The waiting room is full, so Client #81 leaves.
Patrick has finished cutting Client #52's hair
Patrick is cutting Client #61's hair.
*** Client #82 has arrived.
Client #82 takes a seat in the waiting room.
*** Client #83 has arrived.
The waiting room is full, so Client #83 leaves.
Jim has finished cutting Client #53's hair
*** Client #84 has arrived.
Jim is cutting Client #63's hair.
Client #84 takes a seat in the waiting room.
*** Client #85 has arrived.
The waiting room is full, so Client #85 leaves.
*** Client #86 has arrived.
The waiting room is full, so Client #86 leaves.
*** Client #87 has arrived.
The waiting room is full, so Client #87 leaves.
*** Client #88 has arrived.
The waiting room is full, so Client #88 leaves.
Frank has finished cutting Client #58's hair
Frank is cutting Client #68's hair.
*** Client #89 has arrived.
Gerard has finished cutting Client #59's hair
Client #89 takes a seat in the waiting room.
Gerard is cutting Client #69's hair.
Milton has finished cutting Client #60's hair
Milton is cutting Client #70's hair.
*** Client #90 has arrived.
Client #90 takes a seat in the waiting room.
Patrick has finished cutting Client #61's hair
Patrick is cutting Client #71's hair.
*** Client #91 has arrived.
Client #91 takes a seat in the waiting room.
*** Client #92 has arrived.
Client #92 takes a seat in the waiting room.
Jim has finished cutting Client #63's hair
Jim is cutting Client #72's hair.
*** Client #93 has arrived.
Client #93 takes a seat in the waiting room.
*** Client #94 has arrived.
The waiting room is full, so Client #94 leaves.
Closing shop for today
Frank has finished cutting Client #68's hair
Frank is cutting Client #78's hair.
Gerard has finished cutting Client #69's hair
Gerard is cutting Client #79's hair.
Milton has finished cutting Client #70's hair
Milton is cutting Client #80's hair.
Patrick has finished cutting Client #71's hair
Patrick is cutting Client #82's hair.
Jim has finished cutting Client #72's hair
Jim is cutting Client #84's hair.
Frank has finished cutting Client #78's hair
Frank is cutting Client #89's hair.
Gerard has finished cutting Client #79's hair
Gerard is cutting Client #90's hair.
Milton has finished cutting Client #80's hair
Milton is cutting Client #91's hair.
Patrick has finished cutting Client #82's hair
Patrick is cutting Client #92's hair.
Jim has finished cutting Client #84's hair
Jim is cutting Client #93's hair.
Frank has finished cutting Client #89's hair
Frank does not have to cut anyone's hair right now, so he has a nap.
Frank is going home.
Gerard has finished cutting Client #90's hair
Gerard does not have to cut anyone's hair right now, so he has a nap.
Gerard is going home.
Milton has finished cutting Client #91's hair
Milton does not have to cut anyone's hair right now, so he has a nap.
Milton is going home.
Patrick has finished cutting Client #92's hair
Patrick does not have to cut anyone's hair right now, so he has a nap.
Patrick is going home.
Jim has finished cutting Client #93's hair
Jim does not have to cut anyone's hair right now, so he has a nap.
Jim is going home.
--------------------------------------------------------------
The barbershop has closed for today and everyone has gone home
```


Configuration
-------------

You can customize the behavior of the simulation by adjusting the following variables in `main.go`:

-   `seatCapacity`: The capacity of the waiting room.
-   `arrivalRate`: The average arrival rate of clients.
-   `cutDuration`: The duration it takes to cut one client's hair.
-   `timeOpen`: The duration for which the barbershop remains open.

Feel free to experiment with different configurations to observe how the system behaves under various conditions.
