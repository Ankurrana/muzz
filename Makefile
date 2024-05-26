$(V).SILENT:

run:
	go run . 

add_user1:
	curl --location 'localhost:8000/user/create' --header 'Content-Type: application/json' --data-raw '{"name": "User1","email": "user1@gmail.com","password": "user123","age": 15,"gender": "male","lat": 32.4124,"lon": 32.742}'

add_user2:
	curl --location 'localhost:8000/user/create' --header 'Content-Type: application/json' --data-raw '{"name": "User2","email": "user2@gmail.com","password": "user123","age": 20,"gender": "male","lat": 31.24,"lon": 32.52}'

add_user3:
	curl --location 'localhost:8000/user/create' --header 'Content-Type: application/json' --data-raw '{"name": "User3","email": "user3@gmail.com","password": "user123","age": 25,"gender": "female","lat": 32.124,"lon": 32.14}'

add_user4:
	curl --location 'localhost:8000/user/create' --header 'Content-Type: application/json' --data-raw '{"name": "User4","email": "user4@gmail.com","password": "user123","age": 30,"gender": "female","lat": 32.214,"lon": 32.457}'

login_user1:
	curl --location 'localhost:8000/login' --header 'Content-Type: application/json' --data-raw '{"email" : "user1@gmail.com","password" : "user123"}'

login_user2:
	curl --location 'localhost:8000/login' --header 'Content-Type: application/json' --data-raw '{"email" : "user2@gmail.com","password" : "user123"}'

login_user3:
	curl --location 'localhost:8000/login' --header 'Content-Type: application/json' --data-raw '{"email" : "user3@gmail.com","password" : "user123"}'

login_user4:
	curl --location 'localhost:8000/login' --header 'Content-Type: application/json' --data-raw '{"email" : "user4@gmail.com","password" : "user123"}'


token_user_1 = eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxQGdtYWlsLmNvbSIsInVzZXJfaWQiOjEsImV4cCI6MTcxNjczMjg2OX0.Q2icOvu_qk2y8H_jwxe1M6-BoRWOl-jVwJVGMcI1yxo
token_user_2 = eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIyQGdtYWlsLmNvbSIsInVzZXJfaWQiOjIsImV4cCI6MTcxNjczMjg3N30.1Y9Dn_dIaIsyenSUx_PQd4xr0CzPGrAbn1T1E8745MQ
token_user_3 = eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIzQGdtYWlsLmNvbSIsInVzZXJfaWQiOjMsImV4cCI6MTcxNjczMjg4NH0.T891ClIzuWkdjRYeEmGjCJZYk8wWnOYx_GYFqwrDoFs
token_user_4 = eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXI0QGdtYWlsLmNvbSIsInVzZXJfaWQiOjQsImV4cCI6MTcxNjczMjg5M30.T3RCva80qlKVgZEzUMYG4FJjew3gsIkebWdqkGJtCGE

dis_user1:
	curl --location --request GET 'localhost:8000/discover' --header 'Authorization: Bearer ${token_user_1}' --header 'Content-Type: application/json'

dis_user2:
	curl --location --request GET 'localhost:8000/discover?min_age=20' --header 'Authorization: Bearer ${token_user_2}' --header 'Content-Type: application/json'

dis_user3:
	curl --location --request GET 'localhost:8000/discover?min_age=20&max_age=30' --header 'Authorization: Bearer ${token_user_3}' --header 'Content-Type: application/json'

dis_user4:
	curl --location --request GET 'localhost:8000/discover?min_age=20&max_age=30&gender=male' --header 'Authorization: Bearer ${token_user_4}' --header 'Content-Type: application/json'


swipe_1_3:
	curl --location 'localhost:8000/swipe' --header 'Authorization: Bearer ${token_user_1}' --header 'Content-Type: application/json' --data '{"user_id" : 3,"preference" : "yes"}'

swipe_3_1:
	curl --location 'localhost:8000/swipe' --header 'Authorization: Bearer ${token_user_3}' --header 'Content-Type: application/json' --data '{"user_id" : 1,"preference" : "yes"}'

clean_db:
	export PGPASSWORD=postgres; psql -Upostgres -hlocalhost -c 'drop table users, swipes, matches, user_summaries'