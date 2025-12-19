insert into books (title, author, shelf_id, create_at, update_at)
values ('The Stand', 'Stephen King', 2, datetime('now'), datetime('now')),
       ('c++', 'Stephen King', 2, datetime('now'), datetime('now')),
       ('java', 'Stephen King', 2, datetime('now'), datetime('now')),
       ('golang', 'Stephen King', 2, datetime('now'), datetime('now'));



insert into books (title, author, shelf_id, create_at, update_at)
values ('The Stand', 'Stephen King', 1, datetime('now'), datetime('now')),
       ('rust', 'Stephen King', 2, datetime('now'), datetime('now'));


select * from books;
select * from shelves;