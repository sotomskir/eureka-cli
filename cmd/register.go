// Copyright © 2019 Robert Sotomski <sotomski@gmail.com>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"fmt"
	"github.com/sotomskir/go-eureka-client/eureka"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"
)

var registerCmd = &cobra.Command{
	Use:     "register",
	Aliases: []string{"r"},
	Short:   "Register to Eureka server",
	Run: func(cmd *cobra.Command, args []string) {
		server := viper.GetString("EUREKA_SERVER_URL")
		appID := viper.GetString("EUREKA_APP_ID")
		appIP := viper.GetString("EUREKA_APP_IP")
		appPort, err := strconv.Atoi(viper.GetString("EUREKA_APP_PORT"))
		if err != nil {
			fmt.Printf("Error parsing port: %s", viper.GetString("EUREKA_APP_PORT"))
		}
		config := eureka.GetDefaultEurekaClientConfig()
		config.UseDnsForFetchingServiceUrls = false
		config.Region = "region-cn-hd-1"
		config.AvailabilityZones = map[string]string{
			"region-cn-hd-1": "zone-cn-hz-1",
		}
		config.ServiceUrl = map[string]string{
			"zone-cn-hz-1": server,
		}

		eureka.SetLogger(func(level int, format string, a ...interface{}) {
			if level == eureka.LevelError {
				fmt.Println("[custom logger error] "+format, a)
			} else {
				fmt.Println("[custom logger debug] "+format, a)
			}
		})

		client := eureka.DefaultClient.Config(config).
			Register(appID, appPort)
		if appIP != "" {
			client.GetInstance().IppAddr = appIP
			client.GetInstance().Hostname = appIP
		}
		client.Run()
		select {}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("server", "s", "", "Eureka server url. Also read from EUREKA_SERVER_URL env variable")
	viper.BindPFlag("EUREKA_SERVER_URL", registerCmd.Flags().Lookup("server"))
}
