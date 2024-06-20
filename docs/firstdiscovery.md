
this is software for handling the orders 

customer send request with POST HTTP method to the order service and order service depend on the 
request body handle the request

order service connect to the user and product services with grpc to retrieve required data

save the order in the database

publish the message to the rabbitmq broker
