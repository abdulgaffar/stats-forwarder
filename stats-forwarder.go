func checkError(err error) {
        if err != nil {
                fmt.Println("Error: ", err)
                os.Exit(0)
        }
}

func forward(addr *net.UDPAddr, stat string, blacklist []string) {

        for _, item := range blacklist {
                if(len(item) > 0) {
                        matched, err := regexp.MatchString(fmt.Sprintf("^%s", item), stat)
                        checkError(err)
                        if !matched {
                                targetConn, _ := net.DialUDP("udp", nil, addr)
                                targetConn.Write([]byte(stat))
                        }
                }
        }


}

func getBlacklist() []string {
        file, err := ioutil.ReadFile("blacklist.txt")
        checkError(err)

        return strings.Split(string(file), "\n")
}

func main() {
        blacklist := getBlacklist()

        listenerAddr, _ := net.ResolveUDPAddr("udp", ":8126")
        targetAddr, _ := net.ResolveUDPAddr("udp", ":8125")

        conn, err := net.ListenUDP("udp", listenerAddr)
        checkError(err)

        buf := make([]byte, 5120)

        for {
                bufLen, _, _ := conn.ReadFromUDP(buf)
                go forward(targetAddr, string(buf[0:bufLen-1]), blacklist)
        }
}
