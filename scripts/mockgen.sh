#!/bin/bash

mockgen_cmd="mockgen"
$mockgen_cmd -source=x/bech32ibc/types/expected_keepers.go -package testutil -destination x/bech32ibc/testutil/expected_keepers_mocks.go