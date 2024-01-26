# **Agencies and Offer Management:**

These 2 services were written in GOLANG and used a PostgreSQL database, the Agencies micro-service is responsible for managing agencies. It provides functionalities such as creating, updating, retrieving, and deleting agencies.

The Offers micro-service is responsible for managing offers. It provides functionalities such as creating, updating, retrieving, and deleting offers.

[https://en.wikipedia.org/wiki/Go_(programming_language)](https://en.wikipedia.org/wiki/Go_(programming_language))

It uses a PostgreSQL database ( Relational database ) as  a primary database to store its data, and it provides a GRPC interface to deal with it, it’s dumps and all that it knows is Agency/ offers only, the methods it provides are
## Agency
	rpc GetAgencies(GetAgenciesRequest) returns (GetAgenciesResponse);
    rpc GetAgency(GetAgencyRequest) returns (GetAgencyResponse);
    rpc CreateAgency(CreateAgencyRequest) returns (CreateAgencyResponse);
    rpc UpdateAgency(UpdateAgencyRequest) returns (UpdateAgencyResponse);
    rpc DeleteAgency(DeleteAgencyRequest) returns (DeleteAgencyResponse);

## Offers
	rpc GetOffers(GetOffersRequest) returns (GetOffersResponse);
    rpc GetOffer(GetOfferRequest) returns (GetOfferResponse);
    rpc CreateOffer(CreateOfferRequest) returns (CreateOfferResponse);
    rpc UpdateOffer(UpdateOfferRequest) returns (UpdateOfferResponse);
    rpc DeleteOffer(DeleteOfferRequest) returns (DeleteOfferResponse);
