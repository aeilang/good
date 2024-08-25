alter table jobs
    add constraint
        keyword_length_check check ( array_length(keyword, 1) = 3);