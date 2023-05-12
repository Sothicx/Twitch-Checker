import requests, threading
from random import choice,choices
def checker():
    while True:
        proxy = choice(open('proxies.txt', 'r').read().splitlines())
        proxies = {'http': f'http://{proxy}', 'http': f'http://{proxy}'}
        username = ''.join(choices('abcdefghijklmnopqrstuvwxyz', k=4))
        headers = {
            'Accept': '*/*',
            'Accept-Language': 'nl-NL',
            'Client-Id': 'kimne78kx3ncx6brgo4mv6wki5h1ko',
            'Client-Integrity': 'v4.public.eyJjbGllbnRfaWQiOiJraW1uZTc4a3gzbmN4NmJyZ280bXY2d2tpNWgxa28iLCJjbGllbnRfaXAiOiIyMTMuMTE4LjIxNi4yMTYiLCJkZXZpY2VfaWQiOiI2bXE1ZFRUZWllZUI5dXdNNll5UnJxN3VjYVRTenlTdyIsImV4cCI6IjIwMjMtMDUtMTJUMTI6NDY6NTNaIiwiaWF0IjoiMjAyMy0wNS0xMVQyMDo0Njo1M1oiLCJpc19iYWRfYm90IjoiZmFsc2UiLCJpc3MiOiJUd2l0Y2ggQ2xpZW50IEludGVncml0eSIsIm5iZiI6IjIwMjMtMDUtMTFUMjA6NDY6NTNaIiwidXNlcl9pZCI6IiJ9cqJ2B1vbqkNbdOkMdI_vrPyaJw2o3SEL_g4YeytVxUX_nGyq6XLjZtrc-H-Q5_8KKtn-7_VtdyDtAtxfA7t3Bg',
            'Client-Session-Id': '4c71d60dbc263026',
            'Client-Version': 'd64a8cf8-6cad-4b23-8301-265afcd5ee9b',
            'Connection': 'keep-alive',
            'Content-Type': 'text/plain;charset=UTF-8',
            'Origin': 'https://www.twitch.tv',
            'Referer': 'https://www.twitch.tv/',
            'Sec-Fetch-Dest': 'empty',
            'Sec-Fetch-Mode': 'cors',
            'Sec-Fetch-Site': 'same-site',
            'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36',
            'X-Device-Id': '6mq5dTTeieeB9uwM6YyRrq7ucaTSzySw',
            'sec-ch-ua': '"Google Chrome";v="113", "Chromium";v="113", "Not-A.Brand";v="24"',
            'sec-ch-ua-mobile': '?0',
            'sec-ch-ua-platform': '"Windows"',
        }

        data = '[{"operationName":"UsernameValidator_User","variables":{"username":"'+username+'"},"extensions":{"persistedQuery":{"version":1,"sha256Hash":"fd1085cf8350e309b725cf8ca91cd90cac03909a3edeeedbd0872ac912f3d660"}}}]'

        response = requests.post('https://gql.twitch.tv/gql', headers=headers, data=data,proxies=proxies).json()[0]['data']['isUsernameAvailable']
        if response == True:
            print(f"{username} Available")
            with open("valid.txt", 'a') as f:
                f.write(username + '\n')
        elif response == False:
            print("Taken")
        else:
            print("ratelimit")
if __name__ == '__main__':
    for _ in range(10):
        threading.Thread(target=checker).start()