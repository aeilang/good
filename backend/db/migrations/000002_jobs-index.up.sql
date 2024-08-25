create index if not exists jobs_title_idx on jobs using gin (to_tsvector('simple', title));
create index if not exists jobs_keyword_idx on jobs using gin (keyword);

create index if not exists jobs_company_name_idx on jobs using gin (to_tsvector('simple', company_name))
