// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import "github.com/felts94/mypass/cmd"

func main() {
	cmd.Execute()
}

// func main() {
// 	password := password()
// 	fmt.Printf("Password: %s\n", password)

// }

// // copy/pasterd https://stackoverflow.com/questions/2137357/getpasswd-functionality-in-go
// func password() string {

// 	fmt.Print("Enter Password: ")
// 	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
// 	if err != nil {
// 		panic(err)
// 	}
// 	password := string(bytePassword)

// 	return strings.TrimSpace(password)
// }
