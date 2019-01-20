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

package cmd

import (
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate a strong password and add it to the ring",
	Run: func(cmd *cobra.Command, args []string) {
		label, err := cmd.Flags().GetString("label")
		if err != nil {
			panic(err)
		}

		if label == "" {
			println("no label")
			return
		}
		pwd := password()
		pass := genPwd()
		var kr KeyRing
		if _, err := os.Stat("data"); os.IsNotExist(err) {
			kr.Data = make(map[string]string)
			kr.Data[label] = pass
			encryptFile("data", kr.Bytes(), pwd)
		} else {
			d := decryptFile("data", pwd)
			kr = KeyRingFromBytes(d)
			kr.Data[label] = pass
			encryptFile("data", kr.Bytes(), pwd)
		}
		print("\nPassword for '", label, "' : ", pass, "\n")
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().String("label", "", "label for the password")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func genPwd() string {
	return RandStringBytes(15)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`!@#$%^&*(){}[]+=`!@#$%^&*(){}[]+=`!@#$%^&*(){}[]+="

func RandStringBytes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
