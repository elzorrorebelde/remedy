/*
 * Copyright 2018-2019 Florent Biville (@fbiville)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"flag"
	"github.com/elzorrorebelde/remedy/internal/pkg/core"
	"github.com/elzorrorebelde/remedy/internal/pkg/fs"
	"log"
)

func main() {
	log.Print("Starting...")

	// dependency graph - begin
	environment := core.DefaultEnvironment()
	fileSystem := environment.FileSystem
	configLoader := &core.ConfigurationFileLoader{
		Reader:         fileSystem.FileReader,
		SchemaLocation: environment.SchemaLocation,
		SchemaLoader:   &core.JsonSchemaFileLoader{},
	}
	executionTracker := &core.ExecutionVcsTracker{
		Versioning:   environment.VersioningClient.GetClient(),
		FileSystem:   fileSystem,
		Clock:        environment.Clock,
		ConfigLoader: configLoader,
	}
	configurationResolver := &core.ConfigurationResolver{
		Environment:      environment,
		ExecutionTracker: executionTracker,
		PathMatcher:      &fs.ZglobPathMatcher{},
	}
	remedy := &core.Remedy{Fs: fileSystem}
	// dependency graph - end

	configFile, configuration := loadConfiguration(configLoader, configurationResolver)
	if len(configuration.Files) > 0 {
		remedy.Run(configuration)
		if err := executionTracker.TrackExecution(configFile); err != nil {
			log.Printf("remedy warning, could not save current execution, see below for details\n\t%v\n", err)
		}
	} else {
		log.Print("No files to process")
	}

	log.Print("Done!")
}

func loadConfiguration(configLoader *core.ConfigurationFileLoader, configResolver *core.ConfigurationResolver) (*string, *core.ChangeSet) {
	configFile := flag.String("configuration", "remedy.json", "Path to configuration file")
	flag.Parse()

	userConfiguration, err := configLoader.ValidateAndLoad(*configFile)
	if err != nil {
		log.Fatalf("remedy configuration error, cannot load\n\t%v\n", err)
	}
	configuration, err := configResolver.ResolveEagerly(userConfiguration)
	if err != nil {
		log.Fatalf("remedy configuration error, cannot parse\n\t%v\n", err)
	}
	return configFile, configuration
}
