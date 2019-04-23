import hashlib
import sqlite3

conn = sqlite3.connect("../data.sqlite3")
c = conn.cursor()

userName = "shinka"
rawPassword = "morisama"

salt = 'chu2byo'

pw = rawPassword.encode()+salt.encode()

encripted = hashlib.sha256(pw).hexdigest()

#dummyuser = (userName, encripted)
dummyuser = (userName, rawPassword)

c.execute('INSERT INTO users(screenname, password) VALUES(?, ?)', dummyuser)

conn.commit()
conn.close()
