import socket


class TCPClient:
    def __init__(self, host: str, port: int):
        self.host = host
        self.port = port

    def send_message(self, message: str) -> str:
        with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as client:
            client.connect((self.host, self.port))

            client.sendall((message + "\n").encode())

            response = client.recv(1024).decode()
            return response


def main():
    client = TCPClient("localhost", 8081)

    response = client.send_message("Hello from Python")
    print("Ответ сервера:", response)


if __name__ == "__main__":
    main()