import requests
import json

def main():
	# getClassWithTime("月2","test_id")
	postClass()

def getClassWithTime(time,user_id):
	uri = "http://itolab@hydra.itolab.nitech.ac.jp:8000/classes/" + time + "/" + user_id
	r = requests.get(uri)
	print(r.json())

def postClass():
	# uri = "http://itolab@hydra.itolab.nitech.ac.jp:8000/classes"
	uri = "http://localhost:8000/classes"
	# data = {'Name': 'yoshino'}
	data = {'Name': '知能情報処理モデル特論', 'Time': '月1'}
	headers = {'content-type': 'application/json'}
	r = requests.post(uri, data=json.dumps(data), headers=headers)
	# print(r.body)

if __name__ == '__main__':
	main()
