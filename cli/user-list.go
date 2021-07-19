/*
 * Copyright 2021 The Gort Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cli

import (
	"fmt"
	"sort"

	"github.com/getgort/gort/client"
	"github.com/spf13/cobra"
)

const (
	userListUse   = "list"
	userListShort = "List all existing users"
	userListLong  = "List all existing users."
	userListUsage = `Usage:
  gort user list [flags]

Flags:
  -h, --help   Show this message and exit

Global Flags:
  -P, --profile string   The Gort profile within the config file to use
`
)

// GetUserListCmd is a command
func GetUserListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   userListUse,
		Short: userListShort,
		Long:  userListLong,
		RunE:  userListCmd,
	}

	cmd.SetUsageTemplate(userListUsage)

	return cmd
}

func userListCmd(cmd *cobra.Command, args []string) error {

	gortClient, err := client.Connect(FlagGortProfile)
	if err != nil {
		return err
	}

	users, err := gortClient.UserList()
	if err != nil {
		return err
	}

	// Sort by username, for presentation purposes.
	sort.Slice(users, func(i, j int) bool { return users[i].Username < users[j].Username })

	// We want a dynamic format that looks something like: %-10s%-10s%s
	const metaFormat = "%%-%ds%%-%ds%%s\n"

	lenUsername, lenFullName := len("USERNAME"), len("FULL NAME")

	for _, u := range users {
		if len(u.Username) > lenUsername {
			lenUsername = len(u.Username)
		}

		if len(u.FullName) > lenFullName {
			lenFullName = len(u.FullName)
		}
	}

	format := fmt.Sprintf(metaFormat, lenUsername+3, lenFullName+3)

	fmt.Printf(format, "USERNAME", "FULL NAME", "EMAIL ADDRESS")
	for _, u := range users {
		fmt.Printf(format, u.Username, u.FullName, u.Email)
	}

	return nil
}
