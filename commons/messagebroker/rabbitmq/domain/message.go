/*
BMSFES. 
Copyright (C) 2022-2035 

This file is part of the BMSFES solution. 
BMSFES is free software: you can redistribute it and/or modify 
it under the terms of the GNU Affero General Public License as 
published by the Free Software Foundation, either version 3 of the 
License, or (at your option) any later version.
 
BMSFES solution is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the 
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License 
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package domain

/*todo create base model with builder*/

// Message used by producer to insert messages to queues
// Payload is Sms/Email object deserialized to byte slice
type Message struct {
	Payload []byte
	MessageParams
}

// MessageParams used to specify parameters for rabbitMQ Message
//contentType specifies message type (default: "application/json"),
//messageId is message identifier (optional),
//isTransient specifies if message should be persistent in queue (default: "false")
type MessageParams struct {
	ContentType string
	MessageId   string
	IsTransient bool
	Headers     map[string]interface{}
}
