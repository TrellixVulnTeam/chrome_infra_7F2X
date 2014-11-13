// Copyright 2014 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// This file is generated by gen_pubkeys.py, do not modify.

// +build !release

package keys

// Used by KeysetName().
var keysetName = "testing"

// To make a new private-public key pair:
// > openssl genrsa -out private.pem 4096
// > openssl rsa -in private.pem -pubout -out public.pub

var publicKeys = []PublicKey{
	PublicKey{
		Valid:       true,
		Name:        "testing/testing.pub",
		Fingerprint: "5972fa12856804c5fda4416eed2bae3bd102ba34",
		PEM: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvtM8UlwQC7m8YsfFA1zC
oTY75RDAo42gMWHnJYaopfOetC3PHqW6iux+3e3C3b50d1yBIKYCaz3m0wrjNrYH
/Ry5cZF7gU7vXYmkjFY2F7O/lXXqrGAVZA82GIz6dnBJrUC/kGPh4Um4AqfxabWA
RFJomKn2B6c/t3L4w1x62p6IFhVnsD1ACebILaroFXJbXu750V+9abTaQgDO1n/8
3BD4fupTmEoGCtVGvbclZ6D4MEgnooF6LJvqyoHvPtUCiLgXhdr5gHTvkJ6Ice03
7fTs3oy3MR+hZdV0aelwVxA7S1eR+Rn29Nua50hSbJvP7SKgvHy+tb83QH1fjAbn
pQIDAQAB
-----END PUBLIC KEY-----
`,
	},
}
