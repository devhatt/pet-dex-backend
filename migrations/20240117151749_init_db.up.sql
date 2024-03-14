CREATE TABLE IF NOT EXISTS users
(
    id         uuid primary key default UUID(),
    name       varchar(120),
    type       varchar(20) check(type IN ('fisica', 'juridica')),
    documentId varchar(14),
    avatarUrl  varchar(255),
    email      varchar(128) not null,
    phone      varchar(12)  not null
);

CREATE TABLE IF NOT EXISTS legal_persons
(
    id             uuid primary key default UUID(),
    userId         uuid REFERENCES users (id),
    phone          varchar(12) not null,
    links          varchar(255),
    openingHours   varchar(120)        not null,
    adoptionPolicy longtext    not null
);

CREATE TABLE IF NOT EXISTS persons
(
    id        uuid primary key default UUID(),
    userId    uuid REFERENCES users (id),
    birthdate date not null
);

CREATE TABLE IF NOT EXISTS addresses
(
    id        uuid primary key default UUID(),
    userId    uuid REFERENCES users (id),
    address   varchar(255),
    city      varchar(50),
    state     varchar(20),
    latitude  float,
    longitute float
);

CREATE TABLE IF NOT EXISTS breeds
(
    id              uuid primary key default UUID(),
    name            varchar(255)                      not null,
    specie          varchar(255)                      not null,
    size            varchar(20) check(size IN ('small', 'medium', 'large', 'giant')),
    description     varchar(255),
    height          varchar(10),
    weight          varchar(10),
    physicalChar    varchar(255),
    disposition     varchar(255),
    idealFor        varchar(255),
    fur             varchar(50),
    imgUrl          varchar(255),
    weather         varchar(255),
    dressage        varchar(255),
    orgId           varchar(25),
    lifeExpectancy  varchar(30)
);

CREATE TABLE IF NOT EXISTS pets
(
    id                  uuid primary key default UUID(),
    name                varchar(128)                      not null,
    breedId             uuid REFERENCES breeds(id),
    size                varchar(20) check(size IN ('small', 'medium', 'large', 'giant')),
    weight              decimal(3, 2)                     not null,
    weightMeasure       varchar(2) check(weightMeasure IN ('kg', 'lb')),
    adoptionDate        date                              not null,
    birthdate           date                              not null,
    comorbidity         varchar(255),
    tags                varchar(255),
    castrated           bool,
    availableToAdoption bool             default true,
    userId              uuid REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS pets_image
(
    id    uuid primary key default UUID(),
    url   varchar(255),
    petId uuid REFERENCES pets (id)
);

CREATE TABLE IF NOT EXISTS vaccines
(
    id        uuid primary key default UUID(),
    petId     uuid REFERENCES pets (id),
    name      varchar(128) not null,
    date      date         not null,
    doctorCRM varchar(15)  not null
);