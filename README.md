# mysql-bug

## Issue
When SetConnMaxLifetime is set on a MySQL connection with TLS, connections are closed when their lifetime is up, then another attempt is made to reuse the connection.

    # tcpdump -i eth0 -nn port 3306 |grep Flags
    tcpdump: listening on eth0, link-type EN10MB (Ethernet), capture size 1500 bytes

    # Client sends FIN packet after signaling TLS connection close (omitted for brevity)
    CLIENT.48830 > SERVER.3306: Flags [F.], seq 1077, ack 1874, win 280, options [nop,nop,TS val 64393779 ecr 71514784], length 0

    # Server sends FIN packet and RST to finalize connection
    SERVER.3306 > CLIENT.48830: Flags [F.], seq 1874, ack 1078, win 243, options [nop,nop,TS val 71515744 ecr 64393779], length 0
    SERVER.3306 > CLIENT.48830: Flags [R.], seq 1875, ack 1078, win 243, options [nop,nop,TS val 0 ecr 64393779], length 0

    # Client attempts to reuse the connection
    CLIENT.48830 > SERVER.3306: Flags [.], seq 1078, ack 1875, win 280, options [nop,nop,TS val 64393779 ecr 71515744], length 0

    # Server sends another RST packet
    SERVER.3306 > CLIENT.48830: Flags [R], seq 2333675659, win 0, length 0

    # Client starts new connection
    CLIENT.48832 > SERVER.3306: Flags [S], seq 1951693141, win 29200, options [mss 1460,sackOK,TS val 64393780 ecr 0,nop,wscale 7], length 0

## Usage
Copy dsn.txt.eample to dsn.txt and edit the dsn with valid MySQL server details (TLS is required for this behavior)

## Version Details
Works as tested on Mariadb as distributed with CentOS 7, and a variety of recent versions of Go (1.6, 1.7, 1.8)
