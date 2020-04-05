INSERT INTO public.users(id, username, password, email, mood)
VALUES
(1, 'Trung', 'zxc', 'trung@mail.com', 'sad'),
(2, 'Diego', 'zxc', 'diego@mail.com', 'happy'),
(3, 'Anas', 'zxc', 'anas@mail.com', 'neutral')  


-- do $$
-- begin
-- for i in 1..1000 loop
-- INSERT INTO public.holidays(name, date, country, org_id, created_at, updated_at) VALUES (i, '2020-03-06', CONCAT('Sản phẩm ', i), 321001, '2020-03-13 04:14:27.555083', '2020-03-13 04:14:27.555083');
-- end loop;
-- end;
-- $$;
