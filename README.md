# Distributed-Reservation-System
Distributed-Reservation-System

### High level features
1. User browses through Rooms available for a given date range
2. User reserves a type of Room in a particular Hotel
3. On check-in, hotel manager assigns a Room of that type to the user

### Performance Considerations
#### When do WRITE operations happen ?
1. User reserves a Room
2. User cancels a resevation
3. New Hotel or room is added
#### When do READ operations happen ?
1. Browsing through hotel catalog
2. Browsing through hotel features
- Our system is read-heavy -- Significantly higher amount of READ than WRITE.

### API Requirement
#### Reservation
- GET /api/v1/reservations
- GET /api/v1/reservations/123
- POST /api/v1/reservations
- DELETE /api/v1/reservations/123

### Data Model
- Let's go with relational database i.e PostgreSQL
- Why:
1. ACID properties, transactional guarantees
2. Easier Locking mechanisms
3. Data can be easily sharded for scalability
4. Easier to model hotel and resevation data
5. More READs than WRITEs
6. Mostly CRUD Operations

