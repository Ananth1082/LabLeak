# Initialize the simulator
set ns [new Simulator]
# Open the trace files
set nf [open out.nam w]
set tf [open out.tr w]
$ns namtrace-all $nf
$ns trace-all $tf

# Define the 'finish' procedure
proc finish {} {
    global ns nf tf
    $ns flush-trace
    close $nf
    close $tf
    exec nam out.nam &
    exit 0
}

# Create the nodes
set n0 [$ns node]
set n1 [$ns node]
set n2 [$ns node]
set n3 [$ns node]

# Create the duplex links
$ns duplex-link $n0 $n2 0.1Mb 10ms DropTail
$ns duplex-link $n1 $n2 2Mb 10ms DropTail
$ns duplex-link $n2 $n3 0.5Mb 20ms DropTail

# Set the queue limits
$ns queue-limit $n0 $n2 50
$ns queue-limit $n1 $n2 50
$ns queue-limit $n2 $n3 50

# Create and attach TCP agents
set tcp0 [new AgentTCP]
set sink0 [new AgentTCPSink]
$ns attach-agent $n0 $tcp0
$ns attach-agent $n3 $sink0
$ns connect $tcp0 $sink0

# Create and configure TCP traffic
set ftp0 [new ApplicationFTP]
$ftp0 attach-agent $tcp0
$ns at 0.5 $ftp0 start
$ns at 4.5 $ftp0 stop

# Create and attach UDP agents
set udp1 [new AgentUDP]
set null1 [new AgentNull]
$ns attach-agent $n1 $udp1
$ns attach-agent $n3 $null1
$ns connect $udp1 $null1

# Create and configure CBR traffic for UDP
set cbr1 [new ApplicationTrafficCBR]
$cbr1 set packet-size_ 500
$cbr1 set interval_ 0.01
$cbr1 attach-agent $udp1
$ns at 0.5 $cbr1 start
$ns at 4.5 $cbr1 stop

# Schedule the finish time
$ns at 5.0 finish

# Run the simulation
$ns run