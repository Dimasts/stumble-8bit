## StumbleExploitGo
a simple exploit that can give you unlimited trophies.

reference (https://github.com/revan-ar/bot-stumble)

![image](https://cdn.discordapp.com/attachments/961293295886147675/986579408800321556/Screenshot_1.png)

Tutorial Video( i Skiped When i Play Stumble )(https://www.mediafire.com/file/isjwgxpttbs3g0r/vidma_recorder_16062022_211853.mp4/file)

How To Use? 

Mobile Phone :

1. Run HTTPCANARY
2. Open stumble guys.
3. Play the game until you reach at least round 2/3
4. If You Eleminated Or Win Dont Claim Prize
5. Goto HTTPCANARY 
6. And Click At Link https://hkitkabackend.eastus.cloudapp.azure.com:5010/round/finishv2
7. Goto Request 
8. Copy Value
9. And Goto Ur Platform For Running A Exploit

Termux:

1. apt update && apt upgrade
2. apt install git
3. pkg install golang
4. git clone https://github.com/Dimasts/stumble-burik
5. cd StumbleExploitGo
6. go run main.go

PC : 
1. git clone https://github.com/Dimasts/stumble-burik
2. Open stumble guys.
3. Play the game until you reach at least 1 win.
4. When you playing make sure to sniff by using **HTTP Debuger/Wireshark**.
5. View the packets (**DO NOT CLAIM THE PRIZE or the exploit won't work**)
6. Copy the packets
7. Take the json part and paste it in Tools
8. Run the program (go run main.go)
