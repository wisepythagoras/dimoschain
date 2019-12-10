import sys

sys.path.append(sys.path[0] + r'/../')

from encryption import AESCipher

# Create a new object.
aes = AESCipher("This!s@test K3Y&")
message = "This is a test"

# Encrypt the string.
encrypted = aes.encrypt(message)

print("Encrypted: " + encrypted.encode("hex"))

# Now try to decrypt.
decrypted = aes.decrypt(encrypted)

if decrypted != message:
    print("The decrypted text didn't match the original message")
    sys.exit(1)

print("Decrypted message: " + decrypted)
