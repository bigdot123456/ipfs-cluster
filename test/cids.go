package test

import (
	"encoding/hex"

	peer "github.com/libp2p/go-libp2p-peer"
)

// Common variables used all around tests.
var (
	TestCid1     = "QmP63DkAFEnDYNjDYBpyNDfttu1fvUw99x1brscPzpqmmq"
	TestCid2     = "QmP63DkAFEnDYNjDYBpyNDfttu1fvUw99x1brscPzpqmma"
	TestCid3     = "QmP63DkAFEnDYNjDYBpyNDfttu1fvUw99x1brscPzpqmmb"
	TestCid4     = "zb2rhiKhUepkTMw7oFfBUnChAN7ABAvg2hXUwmTBtZ6yxuc57"
	TestCid4Data = "Cid4Data" // Cid resulting from block put NOT ipfs add
	// ErrorCid is meant to be used as a Cid which causes errors. i.e. the
	// ipfs mock fails when pinning this CID.
	ErrorCid = "QmP63DkAFEnDYNjDYBpyNDfttu1fvUw99x1brscPzpqmmc"
	// Shard and Cdag Cids
	TestShardCid    = "zdpuAoiNm1ntWx6jpgcReTiCWFHJSTpvTw4bAAn9p6yDnznqh"
	TestCdagCid     = "zdpuApF6HZBu8rscHSVJ7ra3VSYWc5dJnnxt42bGKyZ1a4qPo"
	TestMetaRootCid = "QmYCLpFCj9Av8NFjkQogvtXspnTDFWaizLpVFEijHTH4eV"

	TestShardData, _ = hex.DecodeString("a16130d82a58230012209273fd63ec94bed5abb219b2d9cb010cabe4af7b0177292d4335eff50464060a")
	TestCdagData, _  = hex.DecodeString("a16130d82a5825000171122030e9b9b4f1bc4b5a3759a93b4e77983cd053f84174e1b0cd628dc6c32fb0da14")

	TestPeerID1, _ = peer.IDB58Decode("QmXZrtE5jQwXNqCJMfHUTQkvhQ4ZAnqMnmzFMJfLewuabc")
	TestPeerID2, _ = peer.IDB58Decode("QmUZ13osndQ5uL4tPWHXe3iBgBgq9gfewcBMSCAuMBsDJ6")
	TestPeerID3, _ = peer.IDB58Decode("QmPGDFvBkgWhvzEK9qaTWrWurSwqXNmhnK3hgELPdZZNPa")
	TestPeerID4, _ = peer.IDB58Decode("QmZ8naDy5mEz4GLuQwjWt9MPYqHTBbsm8tQBrNSjiq6zBc")
	TestPeerID5, _ = peer.IDB58Decode("QmZVAo3wd8s5eTTy2kPYs34J9PvfxpKPuYsePPYGjgRRjg")
	TestPeerID6, _ = peer.IDB58Decode("QmR8Vu6kZk7JvAN2rWVWgiduHatgBq2bb15Yyq8RRhYSbx")
)

// Hashes of test files generated in this module through calls to GetTestingDir*
// and ipfs add with default chunking and layout
var TestDirCids = [29]string{
	"QmdWLMAyMF23KyBfAXbNzy7sDu3zGGu27eGQRkQTPMqfoE",
	"QmbiPudqP8264Ccia1iyTebFrtGmG3JCW85YmT5Gds1Wt9",
	"QmesCMDCHRETdDYfuGjUaGVZGSE2nGxGETPoBUgwitnsCT",
	"Qmbiz4Y6ByNTrzEwE2veqy7S8gUBJNNvNqxAy6bBShpvr4",
	"QmNtq6PD9Ef6V1UtikhemHBcupjsvr2sDu7bu2DrBSoHWw",
	"QmVorRQhT4DbND8JyhAW7HkNPd7bUzqof8XJKcfGcGmvHF",
	"QmanFw3Fn96DkMc9XSuhZdvXWDk37cpLrKA6L54MniGL9Z",
	"QmRVFNBFwtUKpE7u3Bbd6Nj1QsnyHgryZSTM86bBuphPAn",
	"QmUNLLsPACCz1vLxQVkXqqLX5R1X345qqfHbsf67hvA3Nn",
	"QmPZLJ3CZYgxH4K1w5jdbAdxJynXn5TCB4kHy7u8uHC3fy",
	"QmQntQGk13CkPgr1b3RAseJFaTpVMqQu8zgh21ox1RnwBf",
	"QmbR4x5HwcQLiipfyHazhKYA1Z2yN9burWWdAKJBhoZnK3",
	"QmYdHmrwb9Wd8kkjLg4yKr7EPqKNGx5vHuREU5HNc7sxnk",
	"QmVMmDqWhdH8d4QWyhkkVHYvQYara6ijgCg3PNpvRhbmHo",
	"QmX2Erb4SBNfcv8X5MFa4z1jGPfaaYA1snMUhQyYVdDqTA",
	// Hashes that are not printed out (chunks of files)
	"QmavW3cdGuSfYMEQiBDfobwVtPEjUnML2Ry1q8w8X3Q8Wj",
	"QmfPHRbeerRWgbu5BzxwK7UhmJGqGvZNxuFoMCUFTuhG3H",
	"QmaYNfhw7L7KWX7LYpwWt1bh6Gq2p7z1tic35PnDRnqyBf",
	"QmWWwH1GKMh6GmFQunjq7CHjr4g4z6Q4xHyDVfuZGX7MyU",
	"QmVpHQGMF5PLsvfgj8bGo9q2YyLRPMvfu1uTb3DgREFtUc",
	"QmUrdAn4Mx4kNioX9juLgwQotwFfxeo5doUNnLJrQynBEN",
	"QmdJ86B7J8mfGq6SjQy8Jz7r5x1cLcXc9M2a7T7NmSMVZx",
	"QmS77cTMdyx8P7rP2Gij6azgYPpjp2J34EVYuhB6mfjrQh",
	"QmbsBsDspFcqi7xJ4xPxcNYnduzQ5UQDw9y6trQWZGoEHq",
	"QmakAXHMeyE6fHHaeqicSKVMM2QyuGbS2g8dgUA7ns8gSY",
	"QmTC6vGbH9ABkpXfrMmYkXbxEqH12jEVGpvGzibGZEDVHK",
	"QmRLF116yZdLLw7bLvjnHxXVHrjB2snNoJ1itpQxi8NkZP",
	"QmZ2iUT3W7jh8QNnpWSiMZ1QYgpommCSQFZiPY5VdoCHyv",
	"QmR5mq8smc6zCvo3eRhS47ScgEwKpPw7b1joExysyqgbee",
}
