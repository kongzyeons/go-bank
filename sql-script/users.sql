CREATE TABLE IF NOT EXISTS public.users (
	user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name VARCHAR(100) NULL,
	password VARCHAR(6) NOT NULL,
	created_by VARCHAR(100) NOT NULL,
	created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
	updated_by VARCHAR(100) NULL,
	updated_date TIMESTAMPTZ NULL
);
create index users_name_idx on public.users using  btree (name);

INSERT INTO public.users (name, password, created_by, updated_by)  
VALUES ('admin', '123456', 'admin', 'admin')  
RETURNING user_id;