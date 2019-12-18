@echo off
swagger generate server -t go -m restapi/models
java -jar swagger-codegen-cli.jar generate -i swagger.yml -l javascript -o webapp/src/Swagger


