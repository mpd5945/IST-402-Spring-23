import string

def caesar_decrypt(cipher_text, key):
    #Decrypts the given cipher text using the given key
    characters = string.ascii_letters + string.digits + string.punctuation + " "
    plain_text = ""
    for char in cipher_text:
        index = characters.find(char)
        index = (index - key) % len(characters)
        plain_text += characters[index]
    return plain_text

def caesar_brute_force(cipher_text):
    #Brute forces all possible keys to decrypt the given cipher text
    for key in range(len(string.printable)):
        plain_text = caesar_decrypt(cipher_text, key)
        print("Key: {} | Decrypted text: {}".format(key, plain_text))

# Example usage:
encrypted = 'Khoor Zruog $'
caesar_brute_force(encrypted)
