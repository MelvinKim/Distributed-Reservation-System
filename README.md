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

