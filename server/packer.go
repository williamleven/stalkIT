package main

// Pack users with an action into a message
func packer(input chan *User, action string, output chan *Message) {
	for {
		output <-&Message{
			action,
			<-input,
		}
	}
}
