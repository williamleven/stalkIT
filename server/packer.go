package main

func packer(input chan *User, action string, output chan *Message) {
	for {
		output <-&Message{
			action,
			<-input,
		}
	}
}
