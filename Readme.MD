# csv2text

Quick 'n' dirty app I wrote to take each line in a CSV
file and create a single text file for each one, splitting
fields into newlines and labeling them with the column name.

Important: This is destructive. If you have a file in your 
working directory that's named the same as one that it will
generate, you will overwrite the contents of the existing 
file. It will not ask, it will just do it. Be careful around
files you don't want to destroy.