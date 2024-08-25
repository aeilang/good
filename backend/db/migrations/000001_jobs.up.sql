-- CREATE EXTENSION "uuid-ossp";

CREATE TABLE IF NOT EXISTS jobs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    href text unique not null,
    company_name text NOT NULL,
    company_image text not null,
    title text NOT NULL,
    keyword text[] NOT NULL,
    city text NOT NULL,
    fulltime boolean NOT NULL,
    job_type text NOT NULL DEFAULT '社招',
    description text NOT NULL,
    requirement text NOT NULL,
    price_down integer NOT NULL DEFAULT 0,
    price_up integer NOT NULL DEFAULT 0,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    is_deleted bool not null default false,
    version integer not null default 1
);
