# golang-mess-handler
Golang Mess Handler is just a go file which automatically "creates" the projects,src,pkg and the bin directories.

# Seriously?
Yes guys the Golang Mess Handler is a silly project but it was made for lazy people(like me) in mind.
# Well how do i install it?
There are a steps 
2.If your go workspace isn't C:\GoCode then you have to edit it 

a.Open the go file

b.Replace the strings containing C:\GoCode with the path to yor Workspace

For Example :=

If my Workspace is C:\GoWork

I can change it by

Replacing the 24th Line with os.Chdir("C:\\GoWork)
(don't forget the double backslash)

Do it for every line that contains C:\\GoCode

c.Save it 

2.Create a bin directory (or any diretory) for the go file.
 
3.cd into the directory and build the go file with go build

4.Add the bin directory to the PATH environment variable.

5.open up the cmd and type gonew

6.It should open up

# Limitations
Only works with windows(although it can be edited to support unix)

# Questions or Bugs

Email me at planetoid128@gmail.com

# License
This project is under the MIT Licene.
