# create a request to the server to test the server

import requests

url = 'http://localhost:8080'


def voronoi_test():
    testUrl = url + '/api/voronoi'
    data = {
        "matrix": [
            [0, -1, 0, 0, 0, 0, -1, 0, 12, 0],
            [47, -1, 0, -1, 0, -1, 0, 0, 0, 0],
            [0, -1, 0, -1, 0, -1, 0, 0, 0, 0],
            [0, -1, 0, -1, 23, -1, 0, 0, 0, 0],
            [0, -1, 0, -1, 0, -1, 16, 0, 0, 0],
            [0, -1, 0, -1, 0, -1, 0, 0, 0, 0],
            [0, -1, 0, -1, 0, -1, 0, 0, 0, 0],
            [0, 0, 0, -1, 0, -1, 0, 0, 45, 0],
            [0, -1, 0, -1, 0, -1, 0, 0, 0, 0],
            [0, 0, 0, -1, 0, 0, 0, 0, 0, 0],
        ],
        "size": 10
    }
    # load the data as json
    response = requests.post(testUrl, json=data)
    print(response.json())


def bathroom_write_test():
    testUrl = url + '/api/bathroom/write'
    data = { 
        "name": "Stony", 
        "coordinates": [{"lat": 40.909119, "lng": -73.1194032}, {"lat": 40.907119, "lng": -73.1214032}],
        "grid": [
            [0, -1, 0, 0, 0, 0, -1, 0, 12, 0],
            [47, -1, 0, -1, 0, -1, 0, 0, 0, 0],
            [0, -1, 0, -1, 0, -1, 0, 0, 0, 0],
            [0, -1, 0, -1, 23, -1, 0, 0, 0, 0],
            [0, -1, 0, -1, 0, -1, 16, 0, 0, 0],
            [0, -1, 0, -1, 0, -1, 0, 0, 0, 0],
            [0, -1, 0, -1, 0, -1, 0, 0, 0, 0],
            [0, 0, 0, -1, 0, -1, 0, 0, 45, 0],
            [0, -1, 0, -1, 0, -1, 0, 0, 0, 0],
            [0, 0, 0, -1, 0, 0, 0, 0, 0, 0],
        ],
        "bathrooms": [
            {"id": 1, "name": "Copium", "gender": "M",
                "accessible": True, "menstrualProducts": False},
            {"id": 2, "name": "Cope", "gender": "F",
                "accessible": False, "menstrualProducts": True}
        ]
    }
    # load the data as json
    response = requests.post(testUrl, json=data)
    print(response.text)

def bathroom_get_id_test():
    testUrl = url + '/api/bathroom/get/id'
    data = {
        "ID": 922558862
    }
    # load the data as json
    response = requests.get(testUrl, json=data)
    print(response.text)
    print(response.json())

def bathroom_object_write_test():
    testUrl = url + '/api/bathroom/object/write'
    data = {
      "name": "Copium", 
      "gender": "M",
      "accessible": True, 
      "menstrualProducts": False,
    }
    # load the data as json
    response = requests.post(testUrl, json=data)
    print(response.text)
    print(response.json())
        

# voronoi_test()
# bathroom_write_test()
bathroom_object_write_test()
# bathroom_get_id_test()

# [[2, -1, 3, 3, 3, 3, -1, 1, 12, 1],
#  [47, -1, 3, -1, 3, -1, 1, 1, 1, 1],
#  [2, -1, 3, -1, 3, -1, 1, 1, 1, 1],
#  [2, -1, 3, -1, 23, -1, 4, 1, 1, 4],
#  [2, -1, 3, -1, 3, -1, 16, 4, 5, 4],
#  [2, -1, 3, -1, 3, -1, 4, 4, 5, 4],
#  [2, -1, 2, -1, 3, -1, 4, 5, 5, 5],
#  [2, 2, 3, -1, 3, -1, 5, 5, 45, 5],
#  [2, -1, 3, -1, 3, -1, 5, 5, 5, 5],
#  [3, 3, 3, -1, 3, 5, 5, 5, 5, 5]]
