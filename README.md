# gomart-product-service
This repository contains the code and Dockerfile for the product microservice of the **goMart** application, along with the Jenkinsfile describing the CI/CD pipeline for the microservice.

To run the code, you need to have Golang package installed:

1- Download the package from [the official website](https://go.dev/doc/install)
##### For Linux users:
2- Remove any previous Go installation  then extract the archive you just downloaded into /usr/local:
```
 $ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
```
3- Add /usr/local/go/bin to the PATH environment variable:
```
$ export PATH=$PATH:/usr/local/go/bin
```
4- Verify that you've installed Go:
```
$ go version
```
##### For Windows users:
Follow the prompt after opening the MSI file you downloaded from [the official website](https://go.dev/doc/install).

### To test the microservice locally:
Start by making sure all the dependencies are installed and run the code, it will tell you that it's listenning on the application port configured:
```
$ go mod tidy
$ make proto
$ make server
```
To test the microservice, the [API Gateway](https://github.com/RaniaMidaoui/gomart-gateway) must be running in order to redirect the request to the product microservice, you must already have registered and logged in a user with the [authentication microservice](https://github.com/RaniaMidaoui/gomart-authentication-service) and got his authorization token (\$TOKEN):
```
#Create a product
PRODUCT=$(curl --request POST \
  --url http://localhost:3000/product \
  --header "Authorization: Bearer $TOKEN" \
  --header 'Content-Type: application/json' \
  --data '{
 "name": "Product A",
 "stock": 5,
 "price": 15
}')

# Get product ID, you must have jq installed
ID=`jq '.id' <<< $PRODUCT`
echo $ID
echo ""

#Get product by IDID..."
curl --request GET \
  --url "http://localhost:3000/product/$ID" \
  --header "Authorization: Bearer $TOKEN"
echo ""
```