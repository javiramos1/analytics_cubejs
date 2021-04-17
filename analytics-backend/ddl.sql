create schema if not exists analytics;

create table if not exists analytics.events (
	event_id VARCHAR ( 99 ) PRIMARY KEY,
	event_type VARCHAR ( 20 ),
	page_path VARCHAR ( 199 ),
	user_ip VARCHAR ( 50 ),
	user_city VARCHAR ( 120 ),
	user_country_code VARCHAR ( 2 ),
	user_country_name VARCHAR ( 30 ),
	user_country_region VARCHAR ( 50 ),
	user_continent VARCHAR ( 40 ),
	user_timezone VARCHAR ( 50 ),
	user_languages VARCHAR ( 90 ),
	user_in_eu BOOLEAN,
	user_agent VARCHAR ( 120 ),
	user_os VARCHAR ( 60 ),
	is_mobile BOOLEAN,
	click_text VARCHAR ( 599 ),
	click_id VARCHAR ( 90 ),
	click_class VARCHAR ( 299 ),
	event_time TIMESTAMP 
);