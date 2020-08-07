# csv2text

Quick 'n' dirty app I wrote to take each line in a CSV
file and create a single text file for each one, splitting
fields into newlines and labeling them with the column name.
Filenames are created by concatenating the first two column
values with a space, then appending with ".txt". This is because
this was written to specifically deal with a csv where the first
two columns were "first name" and "last name".   It's easy enough 
to change, though.

It takes two arguments, one required, one optional:

* --in <filepath>: path to input CSV file.
* --outdir <dirpath>: (optional) path to directory for output files, defaults to "." 

Important: This is destructive. If you have a file in your 
working directory that's named the same as one that it will
generate, you will overwrite the contents of the existing 
file. It will not ask, it will just do it. Be careful around
files you don't want to destroy.