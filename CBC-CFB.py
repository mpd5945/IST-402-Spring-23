codebook = [[0b00, 0b01], [0b01, 0b10], [0b10, 0b11], [0b11, 0b00]]
message = [0b00, 0b01, 0b10, 0b11]
message1 = []
iv = 0b10

def codebookLookup(xor):
    for i in range(4):
        if codebook[i][0] == xor:
            lookupValue = codebook[i][1]
            return lookupValue


for i in range(len(message)):
    a = f"{message[i]:b}"
    message1.append(a)
  
# CBC encryption
print("\nCBC encryption details:")  
print(f"Plaintext: {message1}" )

stream = iv
ciphertext = []
for i in range(len(message)):
    xor = message[i] ^ stream
    ciphertext.append(codebookLookup(xor))
    stream = ciphertext[i]
    print(f"The ciphered value of {message1[i]} is {ciphertext[i]:b}")

# CFB encryption
print("\nCFB encryption details:")
print(f"Plaintext: {message1}")
stream = iv
ciphertext = []
for i in range(len(message)):
    xor = message[i] ^ stream
    ciphertext.append(codebookLookup(xor))
    stream = ciphertext[i] ^ iv
    print(f"The ciphered value of {message1[i]} is {ciphertext[i]:b}")
