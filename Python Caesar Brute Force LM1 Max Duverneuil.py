ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
alphabet = "abcdefghijklmnopqrstuvwxyz"

def cipher(n, plaintextview):
    #The following line will encrypt the string input by a user and return the ciphertext.
    result = ''
    for l in plaintextview:
        try:
            if l.isupper():
                index = ALPHABET.index(l)
                i = (index + n) % 26
                result += ALPHABET[i]
            else:
                index = alphabet.index(l)
                i = (index + n) % 26
                result += alphabet[i]
        except ValueError:
            result += l
    return result

def decipher(n, cipheredtext):
    #The following line will decrypt the ciphered text from above and return a plain text legible response.
    result = ''
    for l in cipheredtext:
        try:
            if l.isupper():
                index = ALPHABET.index(l)
                i = (index - n) % 26
                result += ALPHABET[i]
            else:
                index = alphabet.index(l)
                i = (index - n) % 26
                result += alphabet[i]
        except ValueError:
            result += l
    return result


message = "IST Four Oh Two With Joe Oakes Is Alright."
key = 22
enc = cipher(key, message)
print("%d . %s" % (key, enc))

dec = decipher(key, enc)
print("%d . %s" % (key, dec))
