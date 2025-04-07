CREATE TABLE IF NOT EXISTS public.users (
	user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name VARCHAR(100) NULL,
	password VARCHAR(6) NULL,
	dummy_col_1 varchar(255) DEFAULT NULL,
	created_by VARCHAR(100) NULL,
	created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
	updated_by VARCHAR(100) NULL,
	updated_date TIMESTAMPTZ NULL
);
create index users_name_idx on public.users using  btree (name);



CREATE TABLE IF NOT EXISTS public.user_greetings (
	user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	greeting VARCHAR(255) NULL,
	dummy_col_2 varchar(255) DEFAULT NULL,
	created_by VARCHAR(100) NULL,
	created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
	updated_by VARCHAR(100) NULL,
	updated_date TIMESTAMPTZ NULL
);

CREATE TABLE IF NOT EXISTS public.transactions (
	transaction_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID DEFAULT gen_random_uuid(),
	name VARCHAR(100) NULL,
	image VARCHAR(255) NULL,
	isBank bool DEFAULT false NOT NULL,
	dummy_col_6 varchar(255) DEFAULT NULL,
	created_by VARCHAR(100) NULL,
	created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);
create index transactions_user_id_idx on public.transactions using  btree (user_id);


CREATE TABLE IF NOT EXISTS public.debit_card_status (
	card_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID DEFAULT gen_random_uuid(),
	status VARCHAR(20) NULL,
	dummy_col_8 varchar(255) DEFAULT NULL,
	created_by VARCHAR(100) NULL,
	created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
	updated_by VARCHAR(100) NULL,
	updated_date TIMESTAMPTZ NULL
);
create index debit_card_status_user_id_idx on public.debit_card_status using  btree (user_id);

CREATE TABLE IF NOT EXISTS public.debit_card_details (
	card_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID DEFAULT gen_random_uuid(),
	issuer VARCHAR(100) NULL,
	number VARCHAR(25) NULL,
	dummy_col_10 varchar(255) DEFAULT NULL,
	created_by VARCHAR(100) NULL,
	created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
	updated_by VARCHAR(100) NULL,
	updated_date TIMESTAMPTZ NULL
);
create index debit_card_details_user_id_idx on public.debit_card_details using  btree (user_id);	

CREATE TABLE IF NOT EXISTS public.debit_card_design (
	card_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID DEFAULT gen_random_uuid(),
	color VARCHAR(10) NULL,
	border_color VARCHAR(10) NULL,
	dummy_col_9 varchar(255) DEFAULT NULL,
	created_by VARCHAR(100) NULL,
	created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
	updated_by VARCHAR(100) NULL,
	updated_date TIMESTAMPTZ NULL
);
create index debit_card_design_user_id_idx on public.debit_card_design using  btree (user_id);	


CREATE TABLE IF NOT EXISTS public.debit_cards (
	card_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID DEFAULT gen_random_uuid(),
	name VARCHAR(100) NULL, 
	dummy_col_7 varchar(255) DEFAULT NULL,
	created_by VARCHAR(100) NULL,
	created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
	updated_by VARCHAR(100) NULL,
	updated_date TIMESTAMPTZ NULL
);
create index debit_cards_user_id_idx on public.debit_cards using  btree (user_id);


CREATE OR REPLACE VIEW vw_debit_cards AS
SELECT 
	dc.card_id,
	dc.user_id,
	dc."name",
	dcs.status,
	dcd."number",
	dcd.issuer,
	dcd2.color,
	dcd2.border_color,
	dc.created_by,
	dc.created_date,
	dc.updated_by,
	dc.updated_date
FROM public.debit_cards dc  
LEFT JOIN public.debit_card_status dcs
	ON dc.card_id = dcs.card_id
LEFT JOIN public.debit_card_details dcd 
	ON dc.card_id = dcd.card_id 
LEFT JOIN public.debit_card_design dcd2 
	ON dc.card_id = dcd2.card_id;	


CREATE TABLE IF NOT EXISTS public.banners (
	banner_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID DEFAULT gen_random_uuid(),
	title VARCHAR(255) NULL,
	description VARCHAR(255) NULL,
	image VARCHAR(255) NULL,
	dummy_col_11 varchar(255) DEFAULT NULL,
	created_by VARCHAR(100) NULL,
	created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
	updated_by VARCHAR(100) NULL,
	updated_date TIMESTAMPTZ NULL
);
create index banners_user_id_idx on public.banners using  btree (user_id);			


CREATE TABLE IF NOT EXISTS public.account_flags (
	flag_id SERIAL PRIMARY KEY, 
	account_id UUID DEFAULT gen_random_uuid(),
	user_id UUID DEFAULT gen_random_uuid(),
	flag_type VARCHAR(50) NOT NULL,
	flag_value VARCHAR(30) NOT NULL,
	created_by VARCHAR(100) NULL,
	created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
	updated_by VARCHAR(100) NULL,
	updated_date TIMESTAMPTZ NULL
);
create index account_flags_account_id_idx on public.account_flags using  btree (account_id);
create index account_flags_user_id_idx on public.account_flags using  btree (user_id);		


CREATE TABLE IF NOT EXISTS public.account_details (
	account_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID DEFAULT gen_random_uuid(),
	name varchar(100) NULL,
	color VARCHAR(10) NULL,
	is_main_account bool DEFAULT false NOT NULL,
	progress BIGINT NULL,
	dummy_col_5 varchar(255) DEFAULT NULL,
	created_by VARCHAR(100) NULL,
	created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
	updated_by VARCHAR(100) NULL,
	updated_date TIMESTAMPTZ NULL
);
create index account_details_user_id_idx on public.account_details using  btree (user_id);	




CREATE TABLE IF NOT EXISTS public.account_balances (
	account_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID DEFAULT gen_random_uuid(),
	amount DECIMAL(15,2) NULL,
	dummy_col_4 varchar(255) DEFAULT NULL,
	created_by VARCHAR(100) NULL,
	created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
	updated_by VARCHAR(100) NULL,
	updated_date TIMESTAMPTZ NULL
);
create index account_balances_user_id_idx on public.account_balances using  btree (user_id);


CREATE TABLE IF NOT EXISTS public.accounts (
	account_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID DEFAULT gen_random_uuid(),
	type VARCHAR(50) NULL,
	currency VARCHAR(10) NULL,
	account_number VARCHAR(20) NULL,
	issuer VARCHAR(100) NULL,
	dummy_col_3 varchar(255) DEFAULT NULL,
	created_by VARCHAR(100) NULL,
	created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
	updated_by VARCHAR(100) NULL,
	updated_date TIMESTAMPTZ NULL
);
create index accounts_user_id_idx on public.accounts using  btree (user_id);	



CREATE OR REPLACE VIEW vw_account AS
SELECT 
	a.account_id,
	a.user_id,
	ad.is_main_account,
	ad."name",
	a."type",
	a.account_number,
	a.issuer,
	ab.amount,
	a.currency,
	ad.color,
	ad.progress,
	a.created_by,
	a.created_date,
	a.updated_by,
	a.updated_date
FROM public.accounts a
LEFT JOIN public.account_balances ab 
	ON a.account_id = ab.account_id 
LEFT JOIN public.account_details ad 
	ON a.account_id = ad.account_id;