import requests
import random
import time

trivia_api = 'https://the-trivia-api.com/api/questions?categories=science&limit=20&difficulty=medium'


def Input_data(trivia_api):
    # Make a request to the API
    response = requests.get(trivia_api)

    url = "http://localhost:4500/MCQ"

    headers = {
        "Content-Type": "application/json"
    }

# Check the status code to make sure the request was successful
    if response.status_code == 200:
        # Convert the response from JSON to a Python dictionary
        data = response.json()["results"]

        list1 = []

        for i in range(len(data)):
            dict1 = {}
            my_values = [data[i]["correct_answer"], data[i]["incorrect_answers"]
                         [0], data[i]["incorrect_answers"][1], data[i]["incorrect_answers"][2]]
            random.shuffle(my_values)
            dict1["Sub"] = data[i]["category"]
            dict1["Ans"] = data[i]["correct_answer"]
            dict1["O1"] = my_values[0]
            dict1["O2"] = my_values[1]
            dict1['O3'] = my_values[2]
            dict1['O4'] = my_values[3]
            dict1["Question"] = data[i]["question"]
            list1.append(dict1)
            response = requests.post(url, json=dict1, headers=headers)
            print("element included")
            time.sleep(5)

    # Do something with the data
        print(list1)
    else:
        # If the request was unsuccessful, print an error message
        print(f"Error: {response.status_code}")


# Input_data("https://opentdb.com/api.php?amount=25&category=19&type=multiple")

def Put_Random_where_sting_is_empty():
    response = requests.get("http://localhost:4500/MCQs")
    data = response.json()
    for i in range(len(data["mcqs"])):
        if data["mcqs"][i]["Sub"] == "":
            id = data["mcqs"][i]["ID"]
            url = f"http://localhost:4500/MCQ/{id}"
            res = requests.put(url=url, data={"Sub" : "general"})
            print(res.json())
        

#Put_Random_where_sting_is_empty()

print(56 % 10)
print(56 / 10)