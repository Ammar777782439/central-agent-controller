package main

import (
	"context"
	"log"
	"net"
	"sync"

	// Make sure this path matches the folder name in your project
	"central-agent-controller/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


var (
	agentStore = make(map[string]*proto.AgentConfigRequest)
	storeMutex = &sync.RWMutex{}
)


type server struct {
	proto.UnimplementedAgentServiceServer
}


func (s *server) SendConfig(ctx context.Context, req *proto.AgentConfigRequest) (*proto.AgentCommandResponse, error) {
	log.Printf("Received configuration from Agent ID: %s | Capabilities: %v", req.AgentId, req.Capabilities)

	// -- Task Policy Map --
	// This is the "brain" of the server, mapping a capability to a command.
	var taskPolicy = map[string]string{
		"filesystem-scanner":  "START_FILESYSTEM_SCAN",
		"network-monitor":     "MONITOR_INCOMING_TRAFFIC",
		"database-auditor":    "AUDIT_DATABASE_LOGS",
		"process-analyzer":    "ANALYZE_RUNNING_PROCESSES",
	}

	var commandsToSend []string

	
	for _, capability := range req.Capabilities {
		if command, ok := taskPolicy[capability]; ok {
			log.Printf("Policy Matched: Agent has capability '%s'. Adding command '%s' to the queue.", capability, command)
			commandsToSend = append(commandsToSend, command)
		}
	}
	
	
	if req.AgentVersion != "2.1.0" {
		log.Printf("Policy Matched: Agent version is outdated. Adding command 'UPDATE_AGENT' to the queue.")
		commandsToSend = append(commandsToSend, "UPDATE_AGENT")
	}

	
	if len(commandsToSend) > 0 {
		log.Printf("Sending command list to Agent %s: %v", req.AgentId, commandsToSend)
		return &proto.AgentCommandResponse{
			Commands: commandsToSend,
			Reason:   "Found one or more tasks based on your current state and capabilities.",
		}, nil
	}
	
	log.Printf("No tasks for Agent %s. Issuing 'WAIT' command.", req.AgentId)
	return &proto.AgentCommandResponse{
		Commands: []string{"WAIT"},
		Reason:   "No suitable tasks available for you at this time.",
	}, nil
}




// func (s *server) SendResult(ctx context.Context, req *proto.AgentCommandResultRequest) (*proto.AcknowledgementResponse, error) {
// 	log.Printf("Received result from Agent ID: %s for Command: '%s' | Status: %s", req.AgentId, req.ExecutedCommand, req.Status)

// 	storeMutex.Lock()
// 	agentStore[req.AgentId] = req.CurrentConfig
// 	storeMutex.Unlock()

	
// 	var nextCommand = ""
// 	var ackMessage = "Result received and processed successfully."

// 	if req.ExecutedCommand == "START_FILESYSTEM_SCAN" && req.Status == "SUCCESS" {
// 		log.Printf("Follow-up Policy: Scan succeeded. Issuing next command: 'ARCHIVE_SCAN_RESULTS'")
// 		nextCommand = "ARCHIVE_SCAN_RESULTS"
// 		ackMessage = "Scan result received successfully. Next task is ready."
// 	}

// 	if req.ExecutedCommand == "UPDATE_AGENT" && req.Status == "SUCCESS" {
// 		log.Printf("Follow-up Policy: Update succeeded. Issuing next command: 'RESTART_AGENT_SERVICE'")
// 		nextCommand = "RESTART_AGENT_SERVICE"
// 		ackMessage = "Update was successful. Please restart the service."
// 	}

// 	return &proto.AcknowledgementResponse{
// 		Message:     ackMessage,
// 		NextCommand: nextCommand,
// 	}, nil
// }

// main function starts the server
func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen on port 9090: %v", err)
	}

	s := grpc.NewServer()

	proto.RegisterAgentServiceServer(s, &server{})
    reflection.Register(s)
	log.Println("Server is running on port 9090, focusing on capability-based task distribution...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}