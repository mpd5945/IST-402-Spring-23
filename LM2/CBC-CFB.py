# CBC and CFB Block Cipher Modes

# Define a codebook array with 4 rows and 2 columns
codebook = [
    [0b00, 0b01],
    [0b01, 0b10],
    [0b10, 0b11],
    [0b11, 0b00]
]

# Define a message array with 4 elements
message = [0b00, 0b01, 0b10, 0b11]

# Define an empty message1 array
message1 = []

# Define an initialization vector (IV)
iv = 0b10

# Define a function to perform a codebook lookup
def codebookLookup(xor):
    for i in range(4):
        if codebook[i][0] == xor:
            lookupValue = codebook[i][1]
            return lookupValue

# Convert the message array to binary and append it to the message1 array
for i in range(len(message)):
    a = f"{message[i]:b}"
    message1.append(a)

# CBC encryption
print("\nCBC encryption details:")
print(f"Plaintext: {message1}" )

stream = iv
ciphertext = []
for i in range(len(message)):
    # XOR the current message element with the IV or the previous ciphertext value
    xor = message[i] ^ stream
    # Perform a codebook lookup on the XOR value
    ciphertext.append(codebookLookup(xor))
    # Set the current ciphertext value as the next IV
    stream = ciphertext[i]
    # Print the ciphered value to the console
    print(f"The ciphered value of {message[i]:b} is {ciphertext[i]:b}")

# Reverse the order of the ciphertext and message arrays for decryption
ciphertext.reverse()
message.reverse()

# CBC decryption
print("\nCBC decryption details:")
print(f"Ciphertext: {ciphertext}")

stream = iv
plaintext = []
for i in range(len(ciphertext)):
    # Perform a codebook lookup on the current ciphertext value
    xor = codebookLookup(ciphertext[i])
    # XOR the result with the IV or the previous ciphertext value
    plaintext.append(xor ^ stream)
    # Set the current ciphertext value as the next IV
    stream = ciphertext[i]
    # Print the deciphered value to the console
    print(f"The deciphered value of {ciphertext[i]:b} is {plaintext[i]:b}")

# Reverse the order of the plaintext array to get the original message
plaintext.reverse()
print(f"\nOriginal message: {plaintext}")

# CFB encryption
print("\nCFB encryption details:")
print(f"Plaintext: {message1}")

stream = iv
ciphertext = []
for i in range(len(message)):
    # XOR the current message element with the IV or the previous ciphertext value
    xor = message[i] ^ stream
    # Perform a codebook lookup on the XOR value
    ciphertext.append(codebookLookup(xor))
    # XOR the current ciphertext value with the IV to get the next IV
    stream = ciphertext[i] ^ iv
    # Print the ciphered value to the console
    print(f"The ciphered value of {message[i]:b} is {ciphertext[i]:b}")

# Reverse the order of the ciphertext and message arrays for decryption
ciphertext.reverse()
message.reverse()

# CFB decryption
print("\nCFB decryption details:")
print(f"Ciphertext: {ciphertext}")

stream = iv
plaintext = []
for i in range(len(ciphertext)):
    # Perform a codebook lookup on the current ciphertext value
    xor = codebookLookup(ciphertext[i])
    # XOR the result with the IV or the previous ciphertext value
    plaintext.append(xor ^ stream)
    # XOR the current ciphertext value with the IV to get the next IV
    stream = ciphertext[i] ^ plaintext[i]
    # Print the deciphered value to the console
    print(f"The deciphered value of {ciphertext[i]:b} is {plaintext[i]:b}")

# Reverse the order of the plaintext array to get the original message
plaintext.reverse()
print(f"\nOriginal message: {plaintext}")