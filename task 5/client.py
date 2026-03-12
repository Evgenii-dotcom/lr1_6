import socket

client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

client.connect(("localhost", 8081))

client.send("Hello from Python\n".encode())

response = client.recv(1024).decode()

print("Ответ сервера:", response)

client.close()