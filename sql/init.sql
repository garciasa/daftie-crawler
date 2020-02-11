CREATE TABLE public.house(
   id UUID PRIMARY KEY,
   url TEXT NOT NULL,
   price VARCHAR (50) NOT NULL,
   title TEXT NOT NULL,
   beds SMALLINT ,
   baths SMALLINT ,
   provider VARCHAR(100),
   eircode VARCHAR(7),
   date_renewed DATE,
   first_listed DATE,
   propertyid VARCHAR(100)
);

CREATE TABLE public.stat(
  id serial unique not null,
  name varchar(100),
  start_date TIMESTAMP,
  end_date TIMESTAMP 
);