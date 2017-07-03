package main

import(
"fmt"
"os"
"log"
"time"
)

func main(){
//Start
fmt.Print("Yipee New Project!!\n");
fmt.Print("Start [Y/N] := ")
var va string;
fmt.Scanln(&va);
if len(va) == 0 || va != "Y" && va != "N"{
// its like if there is no answer or if it ain't "Y" or "N" 
log.Fatal("Not An Affirmative Question")
}
if va == "Y"{
fmt.Print("Project Name := ")
var ar string
fmt.Scanln(&ar)
os.Chdir("C:\\GoCode")
os.Mkdir(ar, os.ModeDir | 0755)
//the permission bits
op := "C:\\GoCode\\" + ar 
// we do it like this or else os.Chsir won't accept two arguments
os.Chdir(op)
os.Mkdir("src", os.ModeDir | 0755)// the permission bits
os.Mkdir("bin", os.ModeDir | 0755)//same
os.Mkdir("pkg", os.ModeDir | 0755)
oe := "C:\\GoCode\\" + ar + "\\src"
os.Chdir(oe)
os.Mkdir("main" , os.ModeDir | 0755)
fmt.Print("Generating......")
// just for fun
time.Sleep(2 * time.Second)
fmt.Print("All Done")
}
if va == "N"{
log.Fatal("Bummer Dude")
}
}