/*
 * Copyright (C) 2019 The "MysteriumNetwork/terms" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package publish

import (
	cienv "github.com/mysteriumnetwork/go-ci/env"
	"github.com/mysteriumnetwork/go-ci/github"
	"github.com/mysteriumnetwork/terms/ci/env"
	"github.com/pkg/errors"
)

// PublishGithub publishes next Github release
func PublishGithub() error {
	nextVersion := cienv.Str(env.NextVersion)

	releaser, err := github.NewReleaser(cienv.Str(cienv.GithubOwner), cienv.Str(cienv.GithubRepository), cienv.Str(cienv.GithubAPIToken))
	if err != nil {
		return err
	}

	if _, err := releaser.Create(nextVersion); err != nil {
		return errors.Wrap(err, "could not create next tag")
	}

	return nil
}
