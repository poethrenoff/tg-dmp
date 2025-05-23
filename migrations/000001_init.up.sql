CREATE TABLE dmp_profile
(
    id         varchar not null primary key,
    interests  varchar
);

INSERT INTO dmp_profile (id, interests) VALUES ('1', '1,2,3'), ('2', '2,3,4'), ('3', '3,4,5'), ('4', '4,5,6'), ('5', '5,6,7')