CREATE TABLE public."Country"
(
    "Id" integer NOT NULL DEFAULT nextval('"Counry_Id_seq"'::regclass),
    "Name_ru" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    "Name_en" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    "Name_de" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "Counry_pkey" PRIMARY KEY ("Id")
)

    TABLESPACE pg_default;

ALTER TABLE public."Country"
    OWNER to postgres;

CREATE TABLE public."Region"
(
    "Id_region" integer NOT NULL DEFAULT nextval('"Region_Id_region_seq"'::regclass),
    "Name_ru" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    "Name_en" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    "Name_de" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    id_country integer NOT NULL,
    CONSTRAINT "Region_pkey" PRIMARY KEY ("Id_region"),
    CONSTRAINT id FOREIGN KEY (id_country)
        REFERENCES public."Country" ("Id") MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
)

    TABLESPACE pg_default;

ALTER TABLE public."Region"
    OWNER to postgres;

CREATE TABLE public."Content"
(
    "IdContent" integer NOT NULL DEFAULT nextval('"Content_IdContent_seq"'::regclass),
    "Article" text COLLATE pg_catalog."default" NOT NULL,
    "DateCreation" text COLLATE pg_catalog."default" NOT NULL,
    "AuthorID" integer NOT NULL,
    "Text" text COLLATE pg_catalog."default" NOT NULL,
    "MiniContent" text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "Content_pkey" PRIMARY KEY ("IdContent"),
    CONSTRAINT "Content_AuthorID_fkey" FOREIGN KEY ("AuthorID")
        REFERENCES public."Users" ("Id_user") MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
)

    TABLESPACE pg_default;

ALTER TABLE public."Content"
    OWNER to postgres;

CREATE TABLE public."Advertisment"
(
    "IdAdvertisment" integer NOT NULL DEFAULT nextval('"Advertisment_IdAdvertisment_seq"'::regclass),
    "TypeAdvertisment" text COLLATE pg_catalog."default" NOT NULL,
    "Price" integer NOT NULL,
    "TotalArea" integer NOT NULL,
    "YearOfContribution" integer NOT NULL,
    "Address" text COLLATE pg_catalog."default" NOT NULL,
    "Description" text COLLATE pg_catalog."default" NOT NULL,
    "NumberOfRooms" integer NOT NULL,
    "IsCommercial" integer NOT NULL,
    CONSTRAINT "Advertisment_pkey" PRIMARY KEY ("IdAdvertisment")
)

    TABLESPACE pg_default;

ALTER TABLE public."Advertisment"
    OWNER to postgres;

CREATE TABLE public."Content"
(
    "IdContent" integer NOT NULL DEFAULT nextval('"Content_IdContent_seq"'::regclass),
    "Article" text COLLATE pg_catalog."default" NOT NULL,
    "DateCreation" text COLLATE pg_catalog."default" NOT NULL,
    "AuthorID" integer NOT NULL,
    "Text" text COLLATE pg_catalog."default" NOT NULL,
    "MiniContent" text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "Content_pkey" PRIMARY KEY ("IdContent"),
    CONSTRAINT "Content_AuthorID_fkey" FOREIGN KEY ("AuthorID")
        REFERENCES public."Users" ("Id_user") MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
)

    TABLESPACE pg_default;

ALTER TABLE public."Content"
    OWNER to postgres;