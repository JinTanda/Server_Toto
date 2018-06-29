import requests
import json

def main():
	uri = "http://localhost:8000/users"
	data = {'Name': 'yoshino'}
	headers = {'content-type': 'application/json'}
	r = requests.post(uri, data=json.dumps(data), headers=headers)

if __name__ == '__main__':
	main()