package main

import (
	"context"
	"log"
	"time"

	// Make sure this path matches the folder name in your project
	"central-agent-controller/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// --- Agent Details ---
	// Change these values to test the server's different policies
	agentID := "agent-007"
	hostname := "WORKSTATION-BETA"
	ipAddress := "192.168.1.11"

	// Let's make multiple server policies true to test the command list
	osVersion := "Windows 10 Pro 22H2"
	agentVersion := "2.0.0" // Outdated version
	capabilities := []string{"filesystem-scanner", "useless-skill"} // A required capability

	// Establish a connection to the server
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()
	client := proto.NewAgentServiceClient(conn)
	
	// The agent's main loop for periodically checking in with the server
	
		// Create the configuration request to be sent to the server
		config := &proto.AgentConfigRequest{
			AgentId:      agentID,
			Hostname:     hostname,
			IpAddress:    ipAddress,
			OsVersion:    osVersion,
			AgentVersion: agentVersion,
			Capabilities: capabilities,
			Timestamp:    time.Now().Unix(),
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

		log.Println("Sending configuration to server...")
		cmdResponse, err := client.SendConfig(ctx, config)
		if err != nil {
			log.Printf("Could not get command from server: %v", err)
			cancel()
			time.Sleep(30 * time.Second) // Wait before retrying
			
		}

		// --- New Logic: Handle the list of commands ---
		log.Printf("Received command list from server: %v (Reason: %s)", cmdResponse.Commands, cmdResponse.Reason)

		// 1. Loop through each command in the list sent by the server
		for _, command := range cmdResponse.Commands {

			// If the command is "WAIT", do nothing and skip to the next main loop cycle
			if command == "WAIT" {
				log.Println("Server requested to WAIT. No actions will be performed.")
				continue
			}

			// 2. Simulate executing the command
			log.Printf("--> Executing command: '%s'", command)
			time.Sleep(2 * time.Second) // Simulate that the work takes time
		

			
			
		}

		cancel()

		// Wait for 60 seconds before the next check-in with the server
		log.Println("--------------------------------------------------")
		log.Println("All tasks completed. Waiting for 60 seconds before next check-in...")
		
	
}