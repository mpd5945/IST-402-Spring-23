class Rotor:
    """
    Models a 'rotor' in an Enigma machine
    Rotor("BCDA", 1) means that A->B, B->C, C->D, D->A and the rotor has been
    rotated once from ABCD (the clear text character 'B' is facing the user)
    Args:
        mappings (string) encipherings for the machine's alphabet.
        offset (int) the starting position of the rotor
    """
    def __init__(self, mappings, offset=0):
        self.initial_offset = offset
        self.reset()
        self.forward_mappings = dict(zip(self.alphabet, mappings))
        self.reverse_mappings = dict(zip(mappings, self.alphabet))

    def reset(self):
        """
        Helper to re-initialize the rotor to its initial configuration
        Returns: void
        """
        self.alphabet = Machine.ALPHABET
        self.rotate(self.initial_offset)
        self.rotations = 1

    def rotate(self, offset=1):
        """
        Rotates the rotor the given number of characters
        Args: offset (int) how many turns to make
        Returns: void
        """

        for _ in range(offset):
            self.alphabet = self.alphabet[1:] + self.alphabet[0]
        self.rotations = offset

    def encipher(self, character):
        """
        Gets the cipher text mapping of a plain text character

        Args: character (char)

        Returns: char
        """
        return self.forward_mappings[character]

    def decipher(self, character):
        """
        Gets the plain text mapping of a cipher text character
        Args: character (char)
        Returns: char
        """
        return self.reverse_mappings[character]

class Reflector:
    """
    Models a 'reflector' in the Enigma machine. Reflector("CDAB")
    means that A->C, C->A, D->B, B->D
    Args: mappings (string) bijective map representing the reflection
          of a character
    """
    def __init__(self, mappings):
        self.mappings = dict(zip(Machine.ALPHABET, mappings))

        for x in self.mappings:
            y = self.mappings[x]
            if x != self.mappings[y]:
                raise ValueError("Mapping for {0} and {1} is invalid".format(x, y))

    def reflect(self, character):
        """
        Returns the reflection of the input character
        Args: character (char)
        Returns: char
        """
        return self.mappings[character]

class Machine:
    """
    Models an Enigma machine (https://en.wikipedia.org/wiki/Enigma_machine)
    Args:
        rotors (list[Rotor]) the configured rotors
        reflector (Reflector) to use
    """

    ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

    def __init__(self, rotors, reflector):
        self.rotors = rotors
        self.reflector = reflector

    def encipher(self, text):
        """
        Encipher the given input
        Args: text (string) plain text to encode
        Returns: string
        """
        return "".join((self.encipher_character(x) for x in text.upper()))

    def decipher(self, text):
        """
        Deccipher the given input
        Args: text (string) cipher text to decode
        Returns: string
        """

        for rotor in self.rotors:
            rotor.reset()

        return self.encipher(text)

    def encipher_character(self, x):
        """
        Runs a character through the machine's cipher algorithm

        1. If x is not in the known character set, don't encipher it
        2. For each of the rotors, determine the character in contact with x.
           Determine the enciphering for that character, and use it as the next
           letter to pass through to the next rotor in the machine's sequence
        3. Once we get to the reflector, get the reflection and repeat the above
           in reverse
        4. Rotate the first rotor, and check if any other rotor should be rotated
        5. Return the character at the terminating contact position as the input
           character's enciphering

        Args: x (char) the character to encode

        Returns: char
        """

        if x not in Machine.ALPHABET:
            return x

        # compute the contact position of the first rotor and machine's input
        contact_index = Machine.ALPHABET.index(x)

        # propagate contact right
        for rotor in self.rotors:
            contact_letter = rotor.alphabet[contact_index]
            x = rotor.encipher(contact_letter)
            contact_index = rotor.alphabet.index(x)

        # reflect and compute the starting contact position with the right rotor
        contact_letter = Machine.ALPHABET[contact_index]
        x = self.reflector.reflect(contact_letter)
        contact_index = Machine.ALPHABET.index(x)

        # propagate contact left
        for rotor in reversed(self.rotors):
            contact_letter = rotor.alphabet[contact_index]
            x = rotor.decipher(contact_letter)
            contact_index = rotor.alphabet.index(x)

        # rotate the first rotor and anything else that needs it
        self.rotors[0].rotate()
        for index in range(1, len(self.rotors)):
            rotor = self.rotors[index]
            turn_frequency = len(Machine.ALPHABET)*index
            if self.rotors[index-1].rotations % turn_frequency == 0:
                rotor.rotate()

        # finally 'light' the output bulb
        return Machine.ALPHABET[contact_index]


r1 = Rotor("VEADTQRWUFZNLHYPXOGKJIMCSB", 1)
r2 = Rotor("WNYPVJXTOAMQIZKSRFUHGCEDBL", 2)
r3 = Rotor("DJYPKQNOZLMGIHFETRVCBXSWAU", 3)
reflector = Reflector("EJMZALYXVBWFCRQUONTSPIKHGD")
machine = Machine([r1, r2, r3], reflector)
pt = "ATTACK AT DAWN"
print("Ciphertext: " + pt)
ct = machine.encipher(pt)
print("Ciphertext: " + ct)
print("Plaintext: " + machine.decipher(ct))
