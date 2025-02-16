# IntelliAgric Api in Golang

Basically, IntelliAgric is a capstone project I worked on during Uni. It is a tech-driven solution, that leverages IoT sensors, machine learning, with large language models (GPT-3.5) to optimize small-scale farming operations. The platform collects real-time data on soil conditions, and weather patterns, which is processed by AI algorithms to generate actionable insights for farmers. These insights include personalized crop management strategies, disease detection, and fertilizer recommendations, enhancing productivity and sustainability. The system also incorporates an e-commerce feature, facilitating direct transactions between farmers and buyers. The farmer interacts with the system through a chat in a conversational way. 

This repo is for converting the [project](https://github.com/tinotenda-alfaneti/intelliagric-capstone) to Go. 

## Run Locally

1. Run/Start your database - `docker run --name <container name> -e POSTGRES_USER=<username> -e POSTGRES_PASSWORD=<password> -e POSTGRES_DB=<dbname> -p 5432:5432 -d postgres`

2. Start the server -  `make start`



## Technologies/Concepts

1. Docker
2. Golang
3. Gin framework
4. Repository design pattern
5. IoC using Dependency Injection
6. Testing
7. GORM
8. Api documentation in Go
9. JWT
