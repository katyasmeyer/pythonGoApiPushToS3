import os.path
import boto3
import requests
import json
import os


def myFunctionCreate(r):
    f=open("KatyaTest.txt", "w+")
    f.write(r.text)
    f.close()


os.getenv("URL")
os.getenv("bucketname")
data={"Input":"Love"}
data_json=json.dumps(data)
urlPath = "products"
r=requests.get(url=os.getenv("URL"), data=data_json)
print("Response ="+r.text)
myFunctionCreate(r)
print(r.url)
s3 = boto3.client('s3')
filename="KatyaTest.txt"
s3.upload_file(filename, os.getenv("bucketname"), filename)




# if not (os.path.exists('./KatyaTest.txt')):
    # myFunction()





