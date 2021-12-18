import std/[terminal, os, strutils, osproc, httpclient]


const urls = @["https://github.com/", "https://gitlab.com/", "https://bitbucket.org/"]
const git = slurp("Git-2.34.1-64-bit.exe")
let client = httpclient.newHttpClient()
if findExe("git") != "":
    if os.paramCount() > 0:
        if os.paramStr(1) == "find":
            if os.paramCount() == 3 or os.paramCount() == 4: #?Possible fix
                if os.paramStr(3).startsWith("-b=") or os.paramStr(4).startsWith("-b="): #*has to check paramCount first
                    stdout.styledWriteLine(fgBlue, "Downloading branch ", fgGreen, paramStr(3).split("=")[1], fgBlue, " in repository ", fgGreen, paramStr(2))
                    if client.get(urls[0] & paramStr(2)).status != "404 Not Found":
                        discard execCmd("git clone " & urls[0] & paramStr(2) & " -b " & paramStr(3).split("=")[1])
                        stdout.styledWriteLine(fgGreen, "Successfully downloaded repository!")
                    elif client.get(urls[1] & paramStr(2)).status != "404 Not Found":
                        discard execCmd("git clone " & urls[1] & paramStr(2) & " -b " & paramStr(3).split("=")[1])
                        stdout.styledWriteLine(fgGreen, "Successfully downloaded repository!")
                    elif client.get(urls[2] & paramStr(2)).status != "404 Not Found":
                        discard execCmd("git clone " & urls[2] & paramStr(2) & " -b " & paramStr(3).split("=")[1])
                        stdout.styledWriteLine(fgGreen, "Successfully downloaded repository!")
                    else:
                        stdout.styledWriteLine(fgRed, "Error: ", resetStyle, "Failed to download repository, either unsupported git website, or repository/branch does not exist.")
                else: 
                    stdout.styledWriteLine(fgBlue, "Downloading repository ", fgGreen, paramStr(2))
                    if client.get(urls[0] & paramStr(2)).status != "404 Not Found":
                    # echo client.get(urls[0] & paramStr(2)).status (debug)
                        discard execCmd("git clone " & urls[0] & paramStr(2))
                        stdout.styledWriteLine(fgGreen, "Successfully downloaded repository!")
                    elif client.get(urls[1] & paramStr(2)).status != "404 Not Found":
                        discard execCmd("git clone " & urls[1] & paramStr(2))
                        stdout.styledWriteLine(fgGreen, "Successfully downloaded repository!")
                    elif client.get(urls[2] & paramStr(2)).status != "404 Not Found":
                        discard execCmd("git clone " & urls[2] & paramStr(2))
                        stdout.styledWriteLine(fgGreen, "Successfully downloaded repository!")
                    else:
                        stdout.styledWriteLine(fgRed, "Error: ", resetStyle, "Failed to download repository, either unsupported git website, or repository does not exist.")
                if os.paramStr(3) == "-D" or os.paramStr(4) == "-D":
                    setCurrentDir(getAppDir() & DirSep & paramStr(2).split("/")[1])
                    removeDir(".git") #TODO: Remove other git-related files like .gitignore if found
        elif os.paramStr(1) == "-h":
            echo "gofind - A small, simple wrapper around git clone.\n\nUsage: gof find [git url] [options]\n\n-h            Display this help text\n-b=<branch name>            Clone a specific branch of a repository\n-D            Delete .git folder after clone\n\n\n" #help text
        else:
            stdout.styledWriteLine(fgRed, "Unknown option given, use ", fgYellow, "gof -h")
    else:
        stdout.styledWriteLine(fgWhite, "gofind: ", fgRed, "No options given. Use gof -h for a list of options to use.")                
else:
    stdout.styledWriteLine(fgCyan, "No git executable found in your path, would you like to download it? [Y/N]")
    if stdin.readLine().toLower() == "y":
        writeFile("Git-2.34.1-64-bit.exe", git)
        setCurrentDir(getAppDir())
        discard execCmd(".\\Git-2.34.1-64-bit.exe")
    else:
        stdout.styledWriteLine(fgRed, styleItalic, "Exiting program, you can download git from https://git-scm.com/")